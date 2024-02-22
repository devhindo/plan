package core

import (
    "os"
    "log"
)

func ReadFileContent(filePath string) ([]byte, error) {
    content, err := os.ReadFile(filePath)
    if err != nil {
        log.Fatal(err)
        return nil, err
    }
    return content, nil
}
