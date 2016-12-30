# yuri
CLI tool to parse URIs into their components

## Examples

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

## Developing yuri

### Updating dependencies

Run `make vendor` and check in updates.

### Releasing

* Set `EQUINOX_TOKEN` environment variable
* Ensure `equinox.key` file exists
* Run `make release version=<maj>.<min>.<patch>`
