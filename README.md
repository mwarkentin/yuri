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
$ yuri "https://username:password@stage.example.com:443/path+to+foo?query1=1&query2=2#FRAG" | jq
{
  "fragment": "FRAG",
  "host": "stage.example.com:443",
  "hostname": "stage.example.com",
  "port": "443",
  "opaque": "",
  "password": "password",
  "path": "/path+to+foo",
  "rawpath": "/path+to+foo",
  "rawquery": "query1=1&query2=2",
  "scheme": "https",
  "username": "username"
}
```

### Available fields

Here are the JSON fields provided by yuri:

* `scheme`: type of URI
* `opaque`: encoded opaque data
* `username`: basic auth username
* `password`: basic auth password
* `host`: host or host:port
* `hostname`: host without port
* `port`: port
* `path`: path
* `rawpath`: encoded path
* `rawquery`: encoded query values, without `?`
* `fragment`: fragment for references, without `#`

## Developing yuri

### Running tests

Run `make test`.

### Updating dependencies

Run `make vendor` and check in updates.

### Releasing

* Push tag to Github: `git tag x.y.z && git push origin x.y.z`
* Create release: `gh release create x.y.z`

Binaries for multiple architectures (linux / darwin) will be built in [Github Actions](.github/workflows/release.yml).
