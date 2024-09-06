package lib

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
)

func DeleteFile(path string, info fs.FileInfo, err error) error {
	if err != nil {
		fmt.Println("Cannot load file")
		return nil
	}
	if info.IsDir() {
		return nil
	}
	err = os.Remove(path)
	if err != nil {
		fmt.Printf("Could not delete file %s", err.Error())
		return nil
	}
	fmt.Println("File deleted successfully")
	return nil
}

func DeleteDirectoryFiles(p string) {
	err := filepath.Walk(p, DeleteFile)
	if err != nil {
		fmt.Printf("Could not open folder : %s", err.Error())
	}
}
