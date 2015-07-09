package model

import (
	"fmt"
	"testing"
)

func Test_GetAllAppUsers(t *testing.T) {
	fmt.Println(GetAllAppUsers())
}

func Test_AddAppUser(t *testing.T) {
	au, err := AddAppUser("kawatomo@city.co.jp", "hogehoge", "hogehoge")
	if err != nil {
		t.Error(err)
	}
	t.Log(au)
}
