package main

import (
	"fmt"

	"github.com/riahimedyassin/Clear-Cache/lib"
)

func main() {
	res, err := lib.GetConfig()
	if err != nil {
		fmt.Println("Cannot read config : %s", err.Error())
	}
	fmt.Println(*res)
}
