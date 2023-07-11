package cmd

import (
	"errors"
	"fmt"
	"runtime"

	"github.com/hellohq/hqservice/pkg/def"
	"github.com/powerman/appcfg"
	"github.com/powerman/structlog"
	"github.com/spf13/cobra"
)

var (
	Log                     = structlog.New(structlog.KeyUnit, "main")
	logLevel                = appcfg.MustOneOfString("debug", []string{"debug", "info", "warn", "err"})
	ErrRequireFlagOrCommand = errors.New("require flag or command")
	// RequireFlagOrCommand should be used as cobra.Command.RunE for "empty"
	// commands which are just a containers for subcommands.
	RequireFlagOrCommand = func(_ *cobra.Command, _ []string) error {
		return ErrRequireFlagOrCommand
	}
)

func NewRootCmd() *cobra.Command {
	rootCmd := &cobra.Command{
		Use:           def.ProgName,
		Short:         "Monolith with embedded microservices",
		Version:       fmt.Sprintf("%s", runtime.Version()),
		SilenceErrors: true,
		SilenceUsage:  true,
		RunE:          RequireFlagOrCommand,
	}

	rootCmd.PersistentFlags().Var(&logLevel, "log.level", "log level [debug|info|warn|err]")
	cobra.OnInitialize(func() {
		structlog.DefaultLogger.SetLogLevel(structlog.ParseLevel(logLevel.String()))
	})
	return rootCmd
}
