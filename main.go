package main

import (
	"github.com/hellohq/hqservice/cmd"
	"github.com/hellohq/hqservice/internal/db/sqlite"
	"github.com/hellohq/hqservice/internal/sharedconfig"
	"github.com/hellohq/hqservice/pkg/def"
)

func main() {
	sqlite.NewSqliteClient("file:" + "userId" + ".db?cache=shared&_fk=1")

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
