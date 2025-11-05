package main

import (
    "os/exec"
)

func main() {
    exec.Command("sudo", "systemctl", "restart", "nginx").Run()
}
