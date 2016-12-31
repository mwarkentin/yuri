package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/url"
	"os"

	"github.com/urfave/cli"
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
		parsedURI, err := url.Parse(uri)
		if err != nil {
			log.Fatal(err)
		}

		var m map[string]string
		m = make(map[string]string)

		// TODO: Factor out into function for testing
		m["scheme"] = parsedURI.Scheme
		m["opaque"] = parsedURI.Opaque
		m["host"] = parsedURI.Host
		m["path"] = parsedURI.Path
		m["rawpath"] = parsedURI.RawPath
		m["rawquery"] = parsedURI.RawQuery

		if parsedURI.User == nil {
			m["username"] = ""
			m["password"] = ""
		} else {
			m["username"] = parsedURI.User.Username()
			m["password"], _ = parsedURI.User.Password()
		}

		b, err := json.Marshal(m)
		if err != nil {
			fmt.Println("error:", err)
		}
		os.Stdout.Write(b)

		return nil
	}

	app.Run(os.Args)
}
