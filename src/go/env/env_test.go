package env
import "testing"

func Test_ResolveEnv(t *testing.T) {
    var s string

    s = ResolveEnv("${PATH}")
    t.Log(s)

    s = ResolveEnv("localhost")
    if s != "localhost" {
        t.Error("expected value is [localhost], but ["+s+"]")
    }
}