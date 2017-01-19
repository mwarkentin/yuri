package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/url"
	"os"
	"strings"

	"github.com/urfave/cli"
	"gopkg.in/yaml.v2"
)

// Following code taken from https://github.com/golang/go/commit/1ff19201fd898c3e1a0ed5d3458c81c1f062570b
// TODO: Replace Hostname(), Port(), stripPort(), and portOnly() with native methods once Go 1.8 is available

// Hostname returns u.Host, without any port number.
//
// If Host is an IPv6 literal with a port number, Hostname returns the
// IPv6 literal without the square brackets. IPv6 literals may include
// a zone identifier.
func Hostname(hostport string) string {
	return stripPort(hostport)
}

// Port returns the port part of u.Host, without the leading colon.
// If u.Host doesn't contain a port, Port returns an empty string.
func Port(hostport string) string {
	return portOnly(hostport)
}

func stripPort(hostport string) string {
	colon := strings.IndexByte(hostport, ':')
	if colon == -1 {
		return hostport
	}
	if i := strings.IndexByte(hostport, ']'); i != -1 {
		return strings.TrimPrefix(hostport[:i], "[")
	}
	return hostport[:colon]
}

func portOnly(hostport string) string {
	colon := strings.IndexByte(hostport, ':')
	if colon == -1 {
		return ""
	}
	if i := strings.Index(hostport, "]:"); i != -1 {
		return hostport[i+len("]:"):]
	}
	if strings.Contains(hostport, "]") {
		return ""
	}
	return hostport[colon+len(":"):]
}

// CreateURIMap builds a map out of a URL struct for JSON encoding.
func CreateURIMap(u *url.URL) map[string]string {
	var m map[string]string
	m = make(map[string]string)

	m["scheme"] = u.Scheme
	m["opaque"] = u.Opaque
	m["host"] = u.Host
	m["hostname"] = Hostname(u.Host)
	m["port"] = Port(u.Host)
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
	app.Version = "0.2.0"
	app.Flags = []cli.Flag {
    cli.BoolFlag{
      Name: "yaml",
      Usage: "Enable YAML output",
    },
  }
	app.Action = func(c *cli.Context) error {
		if len(c.Args()) < 1 {
			log.Fatal("No arguments. Usage: yuri [-yaml] <URI>")
		}

		if len(c.Args()) > 2 {
			log.Fatal("More than 2 arguments. Usage: yuri [-yaml] <URI>")
		}

		uri := c.Args().First()
		parsedURI, err := url.Parse(uri)
		if err != nil {
			log.Fatal(err)
		}

		m := CreateURIMap(parsedURI)

		var b []byte
		if c.Bool("yaml") {
			b, err = yaml.Marshal(m)
		} else {
			b, err = json.Marshal(m)
		}
		if err != nil {
			fmt.Println("error:", err)
		}
		os.Stdout.Write(b)

		return nil
	}

	app.Run(os.Args)
}
