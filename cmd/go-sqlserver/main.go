// Copyright (C) 2024 The go-sqlserver Authors. All rights reserved.
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

package main

import (
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/cybergarage/go-logger/log"
	"github.com/cybergarage/go-sqlserver/sql"
	"github.com/urfave/cli/v2"
)

func main() {
	log.SetSharedLogger(log.NewStdoutLogger(log.LevelError))

	server := sql.NewServer()

	var configFile string

	app := &cli.App{
		Name:     sql.ProductName,
		Usage:    "SQL server",
		Version:  sql.Version,
		Compiled: time.Now(),
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:        "config-file",
				Value:       "",
				Usage:       "config file",
				Destination: &configFile,
			},
		},
		Action: func(cCtx *cli.Context) error {
			if cCtx.Bool("help") || cCtx.Bool("version") {
				return nil
			}
			if configFile != "" {
				conf, err := sql.NewConfigWithPath(configFile)
				if err != nil {
					log.Errorf("Couldn't load config file (%s)", err.Error())
					return err
				}
				server.SetConfig(conf)
			}
			err := server.Start()
			if err != nil {
				log.Errorf("%s couldn't be started (%s)", sql.ProductName, err.Error())
				return err
			}
			sigCh := make(chan os.Signal, 1)

			signal.Notify(sigCh,
				os.Interrupt,
				syscall.SIGHUP,
				syscall.SIGINT,
				syscall.SIGTERM)

			exitCh := make(chan int)

			go func() {
				for {
					s := <-sigCh
					switch s {
					case syscall.SIGHUP:
						log.Infof("Caught SIGHUP, restarting...")
						err = server.Restart()
						if err != nil {
							log.Infof("%s couldn't be restarted (%s)", sql.ProductName, err.Error())
							os.Exit(1)
						}
					case syscall.SIGINT, syscall.SIGTERM:
						log.Infof("Caught %s, stopping...", s.String())
						err = server.Stop()
						if err != nil {
							log.Infof("%s couldn't be stopped (%s)", sql.ProductName, err.Error())
							os.Exit(1)
						}
						exitCh <- 0
					}
				}
			}()

			code := <-exitCh

			os.Exit(code)

			return nil
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		os.Exit(1)
	}
}
