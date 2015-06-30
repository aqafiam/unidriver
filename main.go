package main

import (
	"fmt"
	"os"
	"regexp"
	"unidriver/Godeps/_workspace/src/github.com/codegangsta/cli"
	"unidriver/commands"
	"unidriver/parsers"
)

func main() {
	app := cli.NewApp()
	app.Name = "unidriver"
	app.Version = "0.1.0"
	app.HideHelp = true
	app.Usage = "E2E Testing Application"
	app.Author = "aqafiam"
	app.Before = doBefore
	app.Action = doMain
	app.Flags = []cli.Flag{
		cli.HelpFlag,
		browserFlag,
		remoteFlag,
	}
	app.Run(os.Args)
}

func doBefore(c *cli.Context) error {

	cli.AppHelpTemplate = `NAME:
  {{.Name}} - {{.Usage}}

USAGE:
  {{.Name}} [options] [arguments...]

VERSION:
  {{.Version}}{{if or .Author .Email}}

 OPTIONS:
  {{range .Flags}}{{.}}
  {{end}}
  `
	// show help
	args := c.Args()
	if len(args) == 0 {
		cli.ShowAppHelp(c)
		os.Exit(1)
	}
	return nil
}

func doMain(c *cli.Context) {

	// check name of args
	args := c.Args()
	var yamlfiles []string
	for _, arg := range args {
		ok, _ := regexp.MatchString(".ya?ml$", arg)
		if !ok {
			fmt.Println(arg + "is not yaml file.")
			cli.ShowAppHelp(c)
			os.Exit(0)
		} else {
			yamlfiles = append(yamlfiles, arg)
		}
	}

	// open and read yamlfiles
	datas := parsers.ParseYaml(yamlfiles)

	// validate command names
	commands.Validate(datas)

	// driv'n it
	commands.Do(c.String("browser"), c.String("remote"), datas)

	os.Exit(999)
}

var browserFlag = cli.StringFlag{
	Name:  "browser, b",
	Value: "firefox",
	Usage: "browser name",
}

var remoteFlag = cli.StringFlag{
	Name:  "remote, r",
	Value: "http://localhost:4444/wd/hub",
	Usage: "RemoteWebDriver server URL",
}
