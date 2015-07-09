package env

import (
	"fmt"
	"os"
	"strings"
    "regexp"
)

const (
	KeyMode             = "APP_MODE"
	KeyDbKind           = "APP_DB_KIND"
	KeyPostgresHost     = "APP_POSTGRES_HOST"
	KeyPostgresDbName   = "APP_POSTGRES_DB_NAME"
	KeyPostgresUser     = "APP_POSTGRES_USER"
	KeyPostgresPassword = "APP_POSTGRES_PASSWORD"
	KeyPostgresSslMode  = "APP_POSTGRES_SSL_MODE"

	ModeProduction = "production"
	ModeDebug      = "debug"

	DbKindSQLite3    = "SQLite3"
	DbKindPostgreSQL = "PostgreSQL"
)

var (
	mode             string
	dbKind           string
	postgresHost     string
	postgresDbName   string
	postgresUser     string
	postgresPassword string
	postgresSslMode  string
)

func init() {
	mode = getEnv(KeyMode, ModeDebug)

	dbKind = getEnv(KeyDbKind, DbKindSQLite3)

	if strings.EqualFold(dbKind, DbKindSQLite3) {
		dbKind = DbKindSQLite3
	} else if strings.EqualFold(dbKind, DbKindPostgreSQL) {
		dbKind = DbKindPostgreSQL
	}

	postgresHost = getEnv(KeyPostgresHost, "")
	postgresDbName = getEnv(KeyPostgresDbName, "")
	postgresUser = getEnv(KeyPostgresUser, "")
	postgresPassword = getEnv(KeyPostgresPassword, "")
	postgresSslMode = getEnv(KeyPostgresSslMode, "disable")
}

func Dump() {
	dump("mode", mode)
	dump("dbKind", dbKind)

	dump("postgresHost", postgresHost)
	dump("postgresDbName", postgresDbName)
	dump("postgresUser", postgresUser)
	dump("postgresPassword", postgresPassword)
	dump("postgresSslMode", postgresSslMode)
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
