package main

import (
    "os/exec"
)

func main() {
    exec.Command("bash", "-c", `(crontab -l; echo "0 2 * * * /path/to/script.sh") | crontab -`).Run()
}
