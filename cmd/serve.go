package cmd

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/hellohq/hqservice/internal/db"
	"github.com/hellohq/hqservice/internal/http/session"
	"github.com/hellohq/hqservice/internal/sharedconfig"
	"github.com/hellohq/hqservice/ms/auth"
	"github.com/hellohq/hqservice/ms/networth"
	"github.com/hellohq/hqservice/pkg/concurrent"
	"github.com/hellohq/hqservice/pkg/def"
	"github.com/powerman/appcfg"
	"github.com/powerman/structlog"
	"github.com/spf13/cobra"
)

type (
	Ctx             = context.Context
	embeddedService interface {
		Name() string
		Init(cfg *sharedconfig.Shared, serveCmd *cobra.Command) error
		RunServe(ctxStartup, ctxShutdown Ctx, shutdown func(), dbClient *db.Db, sessionManage *session.Manager) error
	}
)

var (
	embeddedServices = []embeddedService{
		&auth.Service{},
		&networth.Service{},
	}
	serveStartupTimeout  = appcfg.MustDuration("3s") // must be less than swarm's deploy.update_config.monitor
	serveShutdownTimeout = appcfg.MustDuration("9s") // `docker stop` use 10s between SIGTERM and SIGKILL
)

func NewServeCmd(cfg *sharedconfig.Shared) *cobra.Command {
	serveCmd := &cobra.Command{
		Use:   "serve",
		Short: "Starts ALL embedded microservices",
		Args:  cobra.NoArgs,
		RunE: func(_ *cobra.Command, _ []string) error {
			Log.Info("started", "version", def.Version())
			defer Log.Info("finished", "version", def.Version())
			ctxStartupCmdServe, cancel := context.WithTimeout(context.Background(), serveStartupTimeout.Value(nil))
			defer cancel()

			dbClient := db.InitDbClient(ctxStartupCmdServe, cfg)
			defer dbClient.PgClient.Close()

			jwkRepo := session.NewJwkRepo(dbClient)
			sessionManager := session.InitSessionManager(ctxStartupCmdServe, cfg, jwkRepo)

			ctxShutdownCmdServe, shutdown := context.WithCancel(context.Background())
			ctxShutdownCmdServe, _ = signal.NotifyContext(ctxShutdownCmdServe, syscall.SIGHUP, syscall.SIGINT, syscall.SIGQUIT, syscall.SIGABRT, syscall.SIGTERM)
			go func() {
				<-ctxShutdownCmdServe.Done()
				time.Sleep(serveShutdownTimeout.Value(nil))
				Log.PrintErr("failed to graceful shutdown", "version", def.Version())
				os.Exit(1)
			}()

			services := make([]func(Ctx) error, len(embeddedServices))
			for i := range embeddedServices {
				name := embeddedServices[i].Name()
				runServe := embeddedServices[i].RunServe
				log := structlog.New(structlog.KeyApp, name)
				ctxStartupCmdMs := structlog.NewContext(ctxStartupCmdServe, log)
				services[i] = func(ctxShutdown Ctx) error {
					ctxShutdowCmdMs := structlog.NewContext(ctxShutdownCmdServe, log)
					return runServe(ctxStartupCmdMs, ctxShutdowCmdMs, shutdown, dbClient, sessionManager)
				}
			}

			// Handle termination signals.
			return concurrent.Serve(ctxShutdownCmdServe, shutdown, services...)
		},
	}

	serveCmd.Flags().Var(&serveStartupTimeout, "timeout.startup", "must be less than swarm's deploy.update_config.monitor")
	serveCmd.Flags().Var(&serveShutdownTimeout, "timeout.shutdown", "must be less than 10s used by 'docker stop' between SIGTERM and SIGKILL")

	return serveCmd
}

func NewMsCmd(cfg *sharedconfig.Shared, serveCmd *cobra.Command) *cobra.Command {
	msCmd := &cobra.Command{
		Use:   "ms",
		Short: "Run given embedded microservice's command",
		RunE:  RequireFlagOrCommand,
	}

	seen := make(map[string]bool)
	for _, service := range embeddedServices {
		name := service.Name()
		if seen[name] {
			panic(fmt.Sprintf("duplicate service: %s", name))
		}
		seen[name] = true

		cmd := &cobra.Command{
			Use:   name,
			Short: fmt.Sprintf("Run %s microservice's command", name),
			RunE:  RequireFlagOrCommand, // TODO: Need to write execute function to run a single service
		}
		err := service.Init(cfg, serveCmd)
		if err != nil {
			Log.Fatalf("failed to init service %s: %s", name, err)
		}
		msCmd.AddCommand(cmd)
	}
	return msCmd
}
