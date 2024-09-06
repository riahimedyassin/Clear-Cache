package main

import (
	"fmt"

	"github.com/riahimedyassin/Clear-Cache/lib"
)

func init() {
	lib.InitConfig()
}

func main() {
	config, err := lib.GetConfig()
	if err != nil {
		fmt.Println("Could not load config")
		return
	}
	fmt.Println(*config)
}
