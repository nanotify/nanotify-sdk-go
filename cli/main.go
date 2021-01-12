package main

import (
	"fmt"

	"github.com/nanotify/nanotify-sdk-go/cli/cmd"
)

func main() {
	err := cmd.Execute()
	if err != nil {
		fmt.Println("Error with command: %s", err.Error())
	}
}
