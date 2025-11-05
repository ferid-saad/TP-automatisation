package main

import (
    "fmt"
    "os/exec"
)

func main() {
    user := "username"
    out, _ := exec.Command("bash", "-c", fmt.Sprintf("ps -u %s | wc -l", user)).Output()
    fmt.Printf("%s has %s processes running.\n", user, out)
}
