package main

import (
	"github.com/hellohq/hqservice/cmd"
	"github.com/hellohq/hqservice/internal/sharedConfig"
	"github.com/hellohq/hqservice/pkg/def"
)

func main() {
	err := def.Init()
	if err != nil {
		cmd.Log.Fatalf("failed to get defaults: %s", err)
	}

	cfg, err := sharedConfig.Get()
	if err != nil {
		cmd.Log.Fatalf("failed to init config: %s", err)
	}

	rootCmd := cmd.NewRootCmd()
	serveCmd := cmd.NewServeCmd(cfg)
	msCmd := cmd.NewMsCmd(cfg, serveCmd)
	rootCmd.AddCommand(serveCmd, msCmd)

	err = rootCmd.Execute()
	if err != nil {
		cmd.Log.Fatal(err)
	}
}
