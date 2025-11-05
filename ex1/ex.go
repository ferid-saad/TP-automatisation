package main

import (
    "fmt"
    "os"
    "path/filepath"
    "time"
)

func main() {
    dir := "/var/log/myapp"
    cutoff := time.Now().Add(-7 * 24 * time.Hour)
    filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
        if err == nil && !info.IsDir() && info.ModTime().Before(cutoff) {
            os.Remove(path)
            fmt.Println("Deleted:", path)
        }
        return nil
    })
}
