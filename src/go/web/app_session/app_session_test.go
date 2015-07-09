package app_session

import (
	"encoding/json"
	"github.com/jabaraster/go-web-scaffold/src/go/web/app_session"
	"testing"
)

func Test_(t *testing.T) {
	lu := app_session.LoginUser{}
	d, e := json.Marshal(lu)
	if e != nil {
		t.Error(e)
		return
	}
	s := string(d)
	t.Error(s)
}
