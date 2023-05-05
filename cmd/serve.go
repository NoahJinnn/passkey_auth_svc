package cmd

import (
	"context"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/hellohq/hqservice/internal/sharedConfig"
	"github.com/hellohq/hqservice/ms/auth"
	"github.com/hellohq/hqservice/ms/networth"
	"github.com/hellohq/hqservice/pkg/concurrent"
	"github.com/hellohq/hqservice/pkg/def"
	"github.com/powerman/appcfg"
	"github.com/powerman/structlog"
	"github.com/spf13/cobra"
)

type Ctx = context.Context
type embeddedService interface {
	Name() string
	Init(cfg *sharedConfig.Shared, cmd, serveCmd *cobra.Command) error
	RunServe(ctxStartup, ctxShutdown Ctx, shutdown func()) error
}

var (
	embeddedServices = []embeddedService{
		&auth.Service{},
		&networth.Service{},
	}
	serveStartupTimeout  = appcfg.MustDuration("3s") // must be less than swarm's deploy.update_config.monitor
	serveShutdownTimeout = appcfg.MustDuration("9s") // `docker stop` use 10s between SIGTERM and SIGKILL
)

func NewServeCmd(cfg *sharedConfig.Shared) *cobra.Command {
	serveCmd := &cobra.Command{
		Use:   "serve",
		Short: "Starts ALL embedded microservices",
		Args:  cobra.NoArgs,
		RunE:  runServeWithGracefulShutdown,
	}

	serveCmd.Flags().Var(&serveStartupTimeout, "timeout.startup", "must be less than swarm's deploy.update_config.monitor")
	serveCmd.Flags().Var(&serveShutdownTimeout, "timeout.shutdown", "must be less than 10s used by 'docker stop' between SIGTERM and SIGKILL")
	return serveCmd
}

func runServeWithGracefulShutdown(_ *cobra.Command, _ []string) error {
	Log.Info("started", "version", def.Version())
	defer Log.Info("finished", "version", def.Version())

	ctxStartup, cancel := context.WithTimeout(context.Background(), serveStartupTimeout.Value(nil))
	defer cancel()

	ctxShutdown, shutdown := context.WithCancel(context.Background())
	ctxShutdown, _ = signal.NotifyContext(ctxShutdown, syscall.SIGHUP, syscall.SIGINT, syscall.SIGQUIT, syscall.SIGABRT, syscall.SIGTERM)
	go func() {
		<-ctxShutdown.Done()
		time.Sleep(serveShutdownTimeout.Value(nil))
		Log.PrintErr("failed to graceful shutdown", "version", def.Version())
		os.Exit(1)
	}()

	services := make([]func(Ctx) error, len(embeddedServices))
	for i := range embeddedServices {
		name := embeddedServices[i].Name()
		runServe := embeddedServices[i].RunServe
		log := structlog.New(structlog.KeyApp, name)
		ctxStartup := structlog.NewContext(ctxStartup, log) //nolint:govet // Shadow.
		services[i] = func(ctxShutdown Ctx) error {
			ctxShutdown = structlog.NewContext(ctxShutdown, log)
			return runServe(ctxStartup, ctxShutdown, shutdown)
		}
	}
	return concurrent.Serve(ctxShutdown, shutdown, services...)
}
