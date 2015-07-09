package env

import (
	"fmt"
	"os"
	"strings"
    "regexp"
)

const (
	KeyMode             = "APP_MODE"

	ModeProduction = "production"
	ModeDebug      = "debug"
)

var (
	mode             string
)

func init() {
	mode = getEnv(KeyMode, ModeDebug)
}

func Dump() {
	dump("mode", mode)
}

func Mode() string {
	return mode
}

func IsProductionMode() bool {
	return strings.EqualFold(ModeProduction, mode)
}

func IsDebugMode() bool {
	return strings.EqualFold(ModeDebug, mode)
}

func dump(key string, val string) {
	fmt.Printf("%s -> %s\n", key, val)
}

func ResolveEnv(value string) string {
    reg := regexp.MustCompile("\\$\\{(.+)\\}")
    matches := reg.FindSubmatch([]byte(value))
    if len(matches) < 2 {
        return value
    }
    key := string(matches[1])
    return getEnv(key, "")
}

func getEnv(key string, defaultValue string) string {
	ret := os.Getenv(key)
	if ret == "" {
		return defaultValue
	}
	return ret
}
