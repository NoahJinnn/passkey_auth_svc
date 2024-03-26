package main

import (
	"github.com/NoahJinnn/passkey_auth_svc/cmd"
	"github.com/NoahJinnn/passkey_auth_svc/internal/sharedconfig"
	"github.com/NoahJinnn/passkey_auth_svc/pkg/def"
)

func main() {
	err := def.Init()
	if err != nil {
		cmd.Log.Fatalf("failed to get defaults: %s", err)
	}

	cfg, err := sharedconfig.Get()
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
