package main

import (
    "fmt"
    "syscall"
)

func main() {
    var stat syscall.Statfs_t
    _ = syscall.Statfs("/", &stat)
    free := stat.Bavail * uint64(stat.Bsize)
    total := stat.Blocks * uint64(stat.Bsize)
    if float64(free)/float64(total) < 0.1 {
        fmt.Println("Espace disque faible (<10%)")
    } else {
        fmt.Println("Espace disque OK")
    }
}
