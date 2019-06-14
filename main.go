package main

import (
	"fmt"
	"graph-labs/tinyviz/random"
	"log"
	"os"

	"github.com/urfave/cli"
)

func main() {
	var format string
	var quality int

	app := cli.NewApp()

	app.Name = "tinyviz"
	app.Version = "0.2.0"
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

		random.ShuffleSeed()

		imgFormat, err := pickFormat(format)
		if err != nil {
			return err
		}

		if c.NArg() > 0 {
			err := visualize(c.Args()[0], imgFormat, quality)
			return err
		}

		visualizeFolder(imgFormat, quality)

		return nil
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
