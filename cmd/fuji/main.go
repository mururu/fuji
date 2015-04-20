// Copyright 2015 Shiguredo Inc. <fuji@shiguredo.jp>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Fuji application command code
package main

import (
	"os"

	log "github.com/Sirupsen/logrus"
	"github.com/codegangsta/cli"

	"github.com/shiguredo/fuji"
)

var app *cli.App
var version string

func main() {
	app = cli.NewApp()
	app.Name = "fuji-gw"
	app.Usage = "fuji-gw -c config-file"
	app.Version = version

	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:   "conf, c",
			Value:  "/etc/fuji-gw/config.ini",
			Usage:  "config filepath",
			EnvVar: "FUJI_CONFIG_FILE",
		},
		cli.BoolFlag{
			Name:  "d",
			Usage: "run in verbose mode",
		},
	}

	app.Action = Action
	app.Run(os.Args)
}

func Action(c *cli.Context) {
	err := ValidateArgs(c)
	if err != nil {
		log.Error(err)
		cli.ShowAppHelp(c)
		return
	}

	if c.Bool("d") {
		log.SetLevel(log.DebugLevel)
	}

	log.Println("start fuji gateway")
	fuji.Start(c.String("conf"))
}

func ValidateArgs(c *cli.Context) error {
	return nil
}
