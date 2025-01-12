package util_test

import (
	"fmt"
	session_util "jeanfo_mix/util/session"
	"testing"

	"github.com/fatih/structs"
)

type TmpSub struct {
	Age int
	Sex *string
}

type Tmp struct {
	Name  string
	Sub   TmpSub
	extra string
}

func TestT(t *testing.T) {
	session_util.ClearAllSession()
	fmt.Println(structs.Map(Tmp{Name: "ee", extra: "extradata"}))
}
