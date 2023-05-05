package cmd

import (
	"fmt"

	"github.com/hellohq/hqservice/internal/sharedConfig"
	"github.com/spf13/cobra"
)

func NewMsCmd(cfg *sharedConfig.Shared, serveCmd *cobra.Command) *cobra.Command {
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
		err := service.Init(cfg, cmd, serveCmd)
		if err != nil {
			Log.Fatalf("failed to init service %s: %s", name, err)
		}
		msCmd.AddCommand(cmd)
	}
	return msCmd
}
