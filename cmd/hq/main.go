// Example monolith with embedded microservices.
package main

import (
	"context"
	"errors"
	"fmt"
	"os"
	"os/signal"
	"runtime"
	"syscall"
	"time"

	"github.com/hellohq/hqservice/internal/sharedConfig"
	auth "github.com/hellohq/hqservice/ms/auth"
	"github.com/hellohq/hqservice/pkg/concurrent"
	"github.com/hellohq/hqservice/pkg/def"
	"github.com/powerman/appcfg"
	"github.com/powerman/structlog"
	"github.com/spf13/cobra"
)

// Ctx is a synonym for convenience.
type Ctx = context.Context

type embeddedService interface {
	Name() string
	Init(cfg *sharedConfig.Shared, cmd, serveCmd *cobra.Command) error
	RunServe(ctxStartup, ctxShutdown Ctx, shutdown func()) error
}

//nolint:gochecknoglobals // Main.
var (
	embeddedServices = []embeddedService{
		&auth.Service{},
	}
	log                     = structlog.New(structlog.KeyUnit, "main")
	logLevel                = appcfg.MustOneOfString("debug", []string{"debug", "info", "warn", "err"})
	ErrRequireFlagOrCommand = errors.New("require flag or command")
	// RequireFlagOrCommand should be used as cobra.Command.RunE for "empty"
	// commands which are just a containers for subcommands.
	RequireFlagOrCommand = func(_ *cobra.Command, _ []string) error {
		return ErrRequireFlagOrCommand
	}
	serveStartupTimeout  = appcfg.MustDuration("3s") // must be less than swarm's deploy.update_config.monitor
	serveShutdownTimeout = appcfg.MustDuration("9s") // `docker stop` use 10s between SIGTERM and SIGKILL

	rootCmd = &cobra.Command{
		Use:           def.ProgName,
		Short:         "Example monolith with embedded microservices",
		Version:       fmt.Sprintf("%s %s", def.Version(), runtime.Version()),
		SilenceErrors: true,
		SilenceUsage:  true,
		RunE:          RequireFlagOrCommand,
	}
	serveCmd = &cobra.Command{
		Use:   "serve",
		Short: "Starts embedded microservices",
		Args:  cobra.NoArgs,
		RunE:  runServeWithGracefulShutdown,
	}
	msCmd = &cobra.Command{
		Use:   "ms",
		Short: "Run given embedded microservice's command",
		RunE:  RequireFlagOrCommand,
	}
)

func main() {
	err := def.Init()
	if err != nil {
		log.Fatalf("failed to get defaults: %s", err)
	}

	cfg, err := sharedConfig.Get()
	if err != nil {
		log.Fatalf("failed to init config: %s", err)
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
			RunE:  RequireFlagOrCommand,
		}
		err := service.Init(cfg, cmd, serveCmd)
		if err != nil {
			log.Fatalf("failed to init service %s: %s", name, err)
		}
		msCmd.AddCommand(cmd)
	}

	rootCmd.PersistentFlags().Var(&logLevel, "log.level", "log level [debug|info|warn|err]")
	serveCmd.Flags().Var(&serveStartupTimeout, "timeout.startup", "must be less than swarm's deploy.update_config.monitor")
	serveCmd.Flags().Var(&serveShutdownTimeout, "timeout.shutdown", "must be less than 10s used by 'docker stop' between SIGTERM and SIGKILL")
	rootCmd.AddCommand(serveCmd, msCmd)

	cobra.OnInitialize(func() {
		structlog.DefaultLogger.SetLogLevel(structlog.ParseLevel(logLevel.String()))
	})
	err = rootCmd.Execute()
	if err != nil {
		log.Fatal(err)
	}
}

func runServeWithGracefulShutdown(_ *cobra.Command, _ []string) error {
	log.Info("started", "version", def.Version())
	defer log.Info("finished", "version", def.Version())

	ctxStartup, cancel := context.WithTimeout(context.Background(), serveStartupTimeout.Value(nil))
	defer cancel()

	ctxShutdown, shutdown := context.WithCancel(context.Background())
	ctxShutdown, _ = signal.NotifyContext(ctxShutdown, syscall.SIGHUP, syscall.SIGINT, syscall.SIGQUIT, syscall.SIGABRT, syscall.SIGTERM)
	go func() {
		<-ctxShutdown.Done()
		time.Sleep(serveShutdownTimeout.Value(nil))
		log.PrintErr("failed to graceful shutdown", "version", def.Version())
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
