package main

import (
    "archive/zip"
    "io"
    "os"
    "path/filepath"
)

func main() {
    src := "/etc"
    dst := "/tmp/backup.zip"
    zipfile, _ := os.Create(dst)
    defer zipfile.Close()

    zipWriter := zip.NewWriter(zipfile)
    defer zipWriter.Close()

    filepath.Walk(src, func(path string, info os.FileInfo, err error) error {
        if info.IsDir() {
            return nil
        }
        relPath, _ := filepath.Rel(src, path)
        f, _ := zipWriter.Create(relPath)
        file, _ := os.Open(path)
        defer file.Close()
        io.Copy(f, file)
        return nil
    })
}
