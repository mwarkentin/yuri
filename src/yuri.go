package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/url"
	"os"

	"gopkg.in/urfave/cli.v1"
)

func main() {
	app := cli.NewApp()
	app.Name = "yuri"
	app.Usage = "parse the urlz!"
	app.Action = func(c *cli.Context) error {
		if len(c.Args()) < 1 {
			log.Fatal("No arguments. Usage: yuri <URI>")
		}

		if len(c.Args()) > 1 {
			log.Fatal("More than 1 argument. Usage: yuri <URI>")
		}

		uri := c.Args().First()
		parsedUri, err := url.Parse(uri)
		if err != nil {
			log.Fatal(err)
		}

		b, err := json.Marshal(parsedUri)
		if err != nil {
			fmt.Println("error:", err)
		}
		os.Stdout.Write(b)

		return nil
	}

	app.Run(os.Args)
}
