package lib

import (
	"fmt"
	"os"
)

func DeleteFile(path string) {
	file, err := os.ReadDir(path)
	if err != nil {
		panic(fmt.Sprintf("Cannot open folder : %s", err.Error()))
	}
	for _, f := range file {
		fmt.Printf(f.Name(), " ", f.IsDir())
	}
}
