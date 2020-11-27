package main

import (
	"fmt"
	"os"

	"github.com/urfave/cli"
)


func main() {
	app := cli.NewApp()
	app.Name="jinlong"
	app.Usage="jinlong -c config/config.local.json"
	printVersion:=false
	app.Flags=[]cli.Flag{
		cli.StringFlag{
			Name:"conf, c",
			Value:"config/config.local.json",
			Usage:"config/config.{local|dev|test|pre|prod}.json",
		},
		cli.BoolFlag{
			Name:"version,v",
			Required:false,
			Usage: "-v",
			Destination: &printVersion,
		},
	}

	app.Action=func(c *cli.Context) error{
		if printVersion{
			fmt.Printf("{#{version,Get()}}")
			return nil
		}

		conf:=c.String("conf")
		config.Init(conf)
		return nil
	}

	app.Run(os.Args)
}