package main

import (
    "fmt"
    "io/fs"
    "os"
    "sort"
)

func main() {
    files, _ := os.ReadDir(".")
    type f struct{ name string; size int64 }
    var fl []f
    for _, file := range files {
        if info, _ := file.Info(); !info.IsDir() {
            fl = append(fl, f{file.Name(), info.Size()})
        }
    }
    sort.Slice(fl, func(i, j int) bool { return fl[i].size > fl[j].size })
    for i, file := range fl[:5] {
        fmt.Printf("%d. %s: %.2f KB\n", i+1, file.name, float64(file.size)/1024)
    }
}
