package core

import (
	"fmt"
	"os"
	"path/filepath"
)

func INIT(file string) {
	dir, err := os.Getwd()
	if err != nil {
		fmt.Println("Error getting current directory:", err)
		return
	}

	fullpath := filepath.Join(dir, file)
	exec(fullpath)
}