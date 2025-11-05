package main

import (
    "fmt"
    "os/exec"
)

func main() {
    out, _ := exec.Command("who").Output()
    fmt.Println(string(out))
}
