package configuration
import (
    "io/ioutil"
    "encoding/json"
)

var (
    _config Configuration
)

func init() {
    _config = loadFrom("./app-configuration.json")
}

const (
    DbKind_SQLite = "SQLite"
    DbKind_PostgreSQL = "PostgreSQL"

    SessionKind_SQLite = "SQLite"
    SessionKind_Cookie = "Cookie"
)

type Configuration struct {
    Database DatabaseConfiguration
    Session SessionConfiguration
}

type DatabaseConfiguration struct {
    Kind string
    SQLite *SQLiteConfiguration
    PostgreSQL *PostgreSQLConfiguration
}

type SessionConfiguration struct {
    MaxAge int
    Kind string
    Cookie *CookieSessionConfiguration
    SQLite *SQLiteSessionConfiguration
}

func Get() Configuration {
    return _config
}

type SQLiteConfiguration struct {
    DatabaseFilePath string
}

type SQLiteSessionConfiguration struct {
    DatabaseFilePath string
    TableName string
    SecretPhrase string
}

type PostgreSQLConfiguration struct {
    Host string
    Port string
    User string
    Password string
    Database string
}

type CookieSessionConfiguration struct {
    SecretPhrase string
}

func loadFrom(path string) Configuration {
    data := mustReadData(path)
    c := Configuration {}
    err := json.Unmarshal(data, &c)
    if err != nil {
        panic(err)
    }
    return c
}

func mustReadData(path string) []byte {
    data, err := ioutil.ReadFile(path)
    if err != nil {
        panic(err)
    }
    return data
}