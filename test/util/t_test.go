package util_test

import (
	"fmt"
	"jeanfo_mix/util"
	"path/filepath"
	"testing"
)

func TestT(t *testing.T) {
	fmt.Println(util.GenRandomString(8, false))
	fmt.Println(util.GenRandomString(6, true))
	fmt.Println(util.GenTimeBasedUUID(28))
	fmt.Println(filepath.Ext("/abc/hahha"))
}
