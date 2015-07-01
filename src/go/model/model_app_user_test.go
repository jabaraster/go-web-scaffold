package model

import (
	"testing"
    "fmt"
)

func Test_GetAllUsers(t *testing.T) {
    fmt.Println(GetAllUsers(_db))
}
