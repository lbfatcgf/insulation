package password_test

import (
	"fmt"
	"testing"

	"insulation/server/base/pkg/password"
)

func TestPassword(t *testing.T) {
	hashed, err := password.Gen("123456")
	if err != nil {
		t.Error(err)
		return
	}
	fmt.Println(hashed)
	ok := password.Compare(hashed, "123456")
	fmt.Printf("%v\n", ok)
}
