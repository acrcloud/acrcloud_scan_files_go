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
	//var user = os.Getenv("USER")

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

	//startingTime := time.Now().UTC()
	//fmt.Println(startingTime)
	//
	//time.Sleep(10 * time.Second)
	//endingTime := time.Now().UTC()
	//fmt.Println(endingTime)
	//util.DoRecognize("yellow.mp4", 270, 10)

	//if len(os.Args) > 1 {
	//	res := util.RecognizeFile(util.InitFile(os.Args[1], ""))
	//	//res := util.RecognizeFile(util.InitFile(os.Args[1],"http://192.168.10.122:1024/tiktok.mp3"))
	//	util.ExportToCsv("", res)
	//	//for _, v := range res {
	//	//	spew.Dump(v)
	//	//
	//	//}
	//} else {
	//	fmt.Println("Usage: ./main filename")
	//}

}
