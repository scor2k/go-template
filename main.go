package main

import (
	"github.com/op/go-logging"
	"github.com/urfave/cli/v2"
	"os"
)

var appVersion = "0.1.0"
var log = logging.MustGetLogger("APP_NAME")
var format = logging.MustStringFormatter(
	`%{color}%{time:15:04:05.000} ▶ %{level:.8s} %{shortfile} %{shortfunc} %{color:reset}| %{message}`,
)

func main() {
	appLogs := logging.NewLogBackend(os.Stderr, "", 0)
	appLogsFormatter := logging.NewBackendFormatter(appLogs, format)
	appLogsLeveled := logging.AddModuleLevel(appLogsFormatter)
	appLogsLeveled.SetLevel(logging.INFO, "")
	logging.SetBackend(appLogsLeveled)

	app := &cli.App{
		Name:     "APP_NAME",
		HelpName: "APP_NAME",
		Usage:    "APP_DESCRIPTION",
		Flags:    []cli.Flag{},
		Commands: []*cli.Command{
			{
				Name:  "serve",
				Usage: "start API server and begin work",
				Flags: []cli.Flag{
					&cli.BoolFlag{
						Name:  "debug",
						Value: false,
						Usage: "Enable debug logging",
					},
				},
				Action: func(c *cli.Context) error {
					if c.Bool("debug") {
						appLogsLeveled.SetLevel(logging.DEBUG, "")
					}
					Serve()
					return nil
				},
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
