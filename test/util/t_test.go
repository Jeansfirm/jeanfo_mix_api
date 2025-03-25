package util_test

import (
	"fmt"
	"testing"
)

func f(num *int) {
	// *num = 4
	fmt.Println(*num)
}

func TestT(t *testing.T) {
	f(nil)
}
