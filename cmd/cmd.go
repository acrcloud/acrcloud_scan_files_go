package cmd

import (
	"acrcloud-scan-tool-golang/util"
	"github.com/urfave/cli"
)

var Scan = cli.Command{
	Name:        "scan",
	Usage:       "start to scan a file",
	Description: "start to scan a file",
	Action:      util.Scan,
	Flags: []cli.Flag{
		stringFlag("mode, m", "local", "`MODE`: local, network"),
		stringFlag("type, t", "file", "`TYPE`: folder, file"),
		stringFlag("filename, f", "", "`PATH`: the file need to scan"),
		stringFlag("url, u", "", "`URL`: the network file you want to scan (when using network mode)"),
		stringFlag("output, o", "", "`OUTPUT`: the directory to save the results"),
		stringFlag("filter, l", "", "`FILTER`: combine, fuzzy"),
	},
}

func stringFlag(name, value, usage string) cli.StringFlag {
	return cli.StringFlag{
		Name:  name,
		Value: value,
		Usage: usage,
	}
}

func intFlag(name string, value int, usage string) cli.IntFlag {
	return cli.IntFlag{
		Name:  name,
		Value: value,
		Usage: usage,
	}
}

func boolFlag(name, usage string) cli.BoolFlag {
	return cli.BoolFlag{
		Name:  name,
		Usage: usage,
	}
}
