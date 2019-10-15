package commands

import "github.com/urfave/cli"

// Global CLI flags
var GlobalFlags = []cli.Flag{
	cli.BoolFlag{
		Name:   "debug",
		Usage:  "run in debug mode",
		EnvVar: "IP_LOCATION_FINDER_DEBUG",
	},
	cli.IntFlag{
		Name:   "http-port, p",
		Usage:  "HTTP server port",
		Value:  8081,
		EnvVar: "CLEAN_PROJECT_FRAME_PORT",
	},
	cli.StringFlag{
		Name:   "http-host, i",
		Usage:  "HTTP server host",
		Value:  "",
		EnvVar: "CLEAN_PROJECT_FRAME_HOST",
	},
	cli.StringFlag{
		Name:   "http-mode, m",
		Usage:  "debug, release or test",
		Value:  "",
		EnvVar: "CLEAN_PROJECT_FRAME_MODE",
	},
}
