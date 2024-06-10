package main

import (
	"os"

	"github.com/urfave/cli/v2"
	log "unknwon.dev/clog/v2"

	"github.com/3tel/icp-web/internal/cmd"
)

func main() {
	app := cli.NewApp()
	app.Name = "icpweb"
	app.Commands = []*cli.Command{
		cmd.Web,
	}
	if err := app.Run(os.Args); err != nil {
		log.Fatal("Failed to start application: %v", err)
	}
}
