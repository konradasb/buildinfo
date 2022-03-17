package buildinfo

import (
	"runtime/debug"
)

var infoMap map[string]string

func init() {
	infoMap = make(map[string]string)

	info, ok := debug.ReadBuildInfo()
	if !ok {
		return
	}

	for _, v := range info.Settings {
		infoMap[v.Key] = v.Value
	}
}

// Get returns value from runtime/debug BuildInfo.Settings slice by key
// or returns empty string otherwise
func Get(k string) string {
	v, ok := infoMap[k]
	if !ok {
		return ""
	}

	return v
}
