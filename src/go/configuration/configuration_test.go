package configuration
import (
    "testing"
)

func Test_Initialize(t *testing.T) {
    path := "./test-configuration.json"
    Initialize(path)
    t.Log(_config.Database.Kind)
    t.Log(_config.Database.SQLite)
    t.Log(_config.Database.PostgreSQL)

    t.Log(_config.Session.Kind)
    t.Log(_config.Session.MaxAge)
    t.Log(_config.Session.Cookie)
    t.Log(_config.Session.SQLite)
}
