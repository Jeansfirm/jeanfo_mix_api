package main

import (
	"fmt"
	"log"
)

type MyError struct {
	Code    int
	Message string
}

func (e *MyError) Error() string {
	return fmt.Sprintf("MyError: %d, %s", e.Code, e.Message)
}

func main() {
	// log.Println("ee")

	log.Fatalln("ff")

	// log.Panicln("gg")

}
