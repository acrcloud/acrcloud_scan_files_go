package main

import (
	"acrcloud-scan-tool-golang/cmd"
	"acrcloud-scan-tool-golang/logger"
	"fmt"
	"github.com/urfave/cli"
	"os"
)

func logo() {

	logo := `
    _   ___ ___  ___ _    ___  _   _ ___  
   /_\ / __| _ \/ __| |  / _ \| | | |   \ 
  / _ \ (__|   / (__| |_| (_) | |_| | |) |
 /_/ \_\___|_|_\\___|____\___/ \___/|___/ 
                                          
`
	fmt.Print(logo)

}

func main() {
	logo()
	app := cli.NewApp()
	app.Name = "ACRCloud Scan Tool"
	app.Usage = "Generate the reports for your media file"
	app.Authors = []cli.Author {
		cli.Author {
			Name:  "ACRCloud",
			Email: "support@acrcloud.com",
		},
	}
	app.Version = "0.0.1"

	app.Commands = []cli.Command{cmd.Scan}
	app.Flags = append(app.Flags, cmd.Scan.Flags...)

	err := app.Run(os.Args)
	if err != nil {
		logger.LogFatal("main", "start error", err)
	}

}
