package model

import (
	"fmt"
	"testing"
)

func Test_GetAllUsers(t *testing.T) {
	fmt.Println(GetAllUsers(_db))
}
