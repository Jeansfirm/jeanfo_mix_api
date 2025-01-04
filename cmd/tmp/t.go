package main

import (
	"errors"
	"fmt"
)

type MyError struct {
	Code    int
	Message string
}

func (e *MyError) Error() string {
	return fmt.Sprintf("MyError: %d, %s", e.Code, e.Message)
}

func main() {
	mye := MyError{Code: 20, Message: "myErrorMsg"}
	mye2 := mye
	mye2.Code = 40

	fmt.Println(mye, mye2)

	mye3 := &mye
	mye3.Code = 50
	fmt.Println(&mye, &mye2)

	errors.New()
}
