package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/url"
	"os"

	"github.com/urfave/cli"
)

// CreateURIMap builds a map out of a URL struct for JSON encoding.
func CreateURIMap(u *url.URL) map[string]string {
	var m map[string]string
	m = make(map[string]string)

	m["scheme"] = u.Scheme
	m["opaque"] = u.Opaque
	m["host"] = u.Host
	m["path"] = u.Path
	m["rawpath"] = u.EscapedPath()
	m["rawquery"] = u.RawQuery
	m["fragment"] = u.Fragment

	if u.User == nil {
		m["username"] = ""
		m["password"] = ""
	} else {
		m["username"] = u.User.Username()
		m["password"], _ = u.User.Password()
	}

	return m
}

func main() {
	app := cli.NewApp()
	app.Name = "yuri"
	app.Usage = "parse the urlz!"
	app.Version = "0.1.0"
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

		m := CreateURIMap(parsedURI)

		b, err := json.Marshal(m)
		if err != nil {
			fmt.Println("error:", err)
		}
		os.Stdout.Write(b)

		return nil
	}

	app.Run(os.Args)
}
