# yuri
CLI tool to parse URIs into their components

## Installation

### Homebrew

This application can be installed via [Homebrew](http://brew.sh/):

```
$ brew install eqnxio/michael_warkentin/yuri
```

### Command line

Command line installation instructions can be found on [equinox](https://dl.equinox.io/michael_warkentin/yuri/stable).

### Source

If you want to install from source:

```
$ git clone git@github.com:mwarkentin/yuri.git
$ cd yuri
$ make install
```

## Examples

### CLI

```
$ yuri https://username:password@stage.example.com:443/path | jq
{
  "Scheme": "https",
  "Opaque": "",
  "User": {},
  "Host": "stage.example.com:443",
  "Path": "/path",
  "RawPath": "",
  "ForceQuery": false,
  "RawQuery": "",
  "Fragment": ""
}
```

### Available fields

Yuri currently leverages the Go [`URL` type](https://golang.org/pkg/net/url/#URL) directly. Here are the fields it provides:

```
Scheme     string
Opaque     string    // encoded opaque data
User       *Userinfo // username and password information
Host       string    // host or host:port
Path       string
RawPath    string // encoded path hint
ForceQuery bool   // append a query ('?') even if RawQuery is empty
RawQuery   string // encoded query values, without '?'
Fragment   string // fragment for references, without '#'
```

## Developing yuri

### Updating dependencies

Run `make vendor` and check in updates.

### Releasing

* Set `EQUINOX_TOKEN` environment variable
* Ensure `equinox.key` file exists
* Run `make release version=<maj>.<min>.<patch>`
