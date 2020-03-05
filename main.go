package main

import (
	"errors"
	"fmt"
)

func main() {
	errors.New("will be ignored")
	err := fmt.Errorf("will be handled")
	if err != nil {
		panic(err)
	}
}
