package main

import (
	"fmt"
	"os"
)

func start() error {
	return nil
}

func main() {
	fmt.Println("hello scale search")
	if start() != nil {
		os.Exit(1)
	}

	os.Exit(0)
}
