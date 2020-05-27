package main

import (
	"fmt"
	watch "github.com/andiogenes/tinyviz/watcher"
	"log"
	"os"
	"time"

	"github.com/urfave/cli"
)

func main() {
	var format string
	var quality int
	var arrangement string
	var hotReload bool

	app := cli.NewApp()

	app.Name = "tinyviz"
	app.Version = "0.3.0"
	app.Description = "Graph visualization tool for educational purposes"
	app.Authors = []cli.Author{
		{Name: "Anton", Email: "megadeathlightsaber@gmail.com"},
	}

	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:        "format, f",
			Value:       "png",
			Usage:       "image output format (jpg/png)",
			Destination: &format,
		},
		cli.IntFlag{
			Name:        "quality, q",
			Value:       80,
			Usage:       "jpeg image ouput quality (from 0 to 100)",
			Destination: &quality,
		},
		cli.StringFlag{
			Name:        "arrangement, a",
			Value:       "random",
			Usage:       "image vertex arrangement (random/coord)",
			Destination: &arrangement,
		},
		cli.BoolFlag{
			Name:        "hotreload, hr",
			Usage:       "revisualize graph when config file changes (works only with single file)",
			Destination: &hotReload,
		},
	}

	app.Action = func(c *cli.Context) error {
		if format != "png" && format != "jpg" && format != "jpeg" {
			fmt.Println("Unknown format \"", format, "\", reset to \"png\"")
			format = "png"
		}

		if quality <= 0 || quality > 100 {
			fmt.Println("Unbounded quality value ", quality, ", reset to 80")
			quality = 80
		}

		if arrangement != "random" && arrangement != "coord" && arrangement != "rand" && arrangement != "r" && arrangement != "c" {
			fmt.Println("Unknown arrangement type \"", arrangement, "\", reset to \"random\"")
			arrangement = "random"
		}

		imgFormat, err := pickFormat(format)
		if err != nil {
			return err
		}

		arrangementFn, err := pickArrangementFn(arrangement)
		if err != nil {
			return err
		}

		dataLoaderFn, err := initDataLoader(arrangement)
		if err != nil {
			return err
		}

		if c.NArg() > 0 {
			if hotReload {
				watcher := watch.NewModTimeWatcher(func(f *os.File) error {
					// Ignore errors
					_ = visualize(f.Name(), imgFormat, quality, arrangementFn, dataLoaderFn)
					return nil
				})
				defer watcher.Close()

				_ = watcher.AddFile(c.Args()[0])
				watcher.Run(time.Millisecond * 100)

				err := <-watcher.Errors
				return err
			}

			err := visualize(c.Args()[0], imgFormat, quality, arrangementFn, dataLoaderFn)
			return err
		}

		visualizeFolder(imgFormat, quality, arrangementFn, dataLoaderFn)

		return nil
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
