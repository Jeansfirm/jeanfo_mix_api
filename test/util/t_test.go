package util_test

import (
	"fmt"
	"testing"
)

func TestT(t *testing.T) {
	cases := []struct {
		Name string
		Age  int
	}{
		{},
		{Age: 13, Name: "kk"},
		{"jeanfo", 18},

		{Name: "zz"},
		{Age: 19},
	}

	for _, c := range cases {
		fmt.Println("Name:", c.Name, "Age:", c.Age)
	}
}
