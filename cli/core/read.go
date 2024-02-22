package core

import (
	"fmt"
	"log"
	"os"
)

func ReadFileContent(filePath string) ([]byte, error) {
    content, err := os.ReadFile(filePath)
    if err != nil {
        log.Fatal(err)
        return nil, err
    }
    fmt.Println(string(content))
    return content, nil
}
