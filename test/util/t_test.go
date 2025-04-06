package util_test

import (
	"fmt"
	auth_service "jeanfo_mix/internal/service/auth"
	"testing"
)

func f(num *int) {
	// *num = 4
	fmt.Println(*num)
}

func TestT(t *testing.T) {
	fmt.Println(auth_service.HashPassword("073d4c9a9aafdefc4ab2f4b9184a93ba"))
}
