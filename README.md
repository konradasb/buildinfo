# buildinfo

[![Go Report Card](https://goreportcard.com/badge/github.com/konradasb/buildinfo)](https://goreportcard.com/report/github.com/konradasb/buildinfo)

Golang package which parses `debug.ReadBuildInfo()` BuildInfo.Settings slice into map

## Requires

* Golang >= 1.18

## Example

```golang
package app

import (
  "github.com/konradasb/buildinfo"
)

func main() {
  fmt.Println(buildinfo.Get("vcs.revision"))
  fmt.Println(buildinfo.Get("vcs.time"))
}
```

## License

MIT
