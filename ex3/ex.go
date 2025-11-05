package main

import (
    "fmt"
    "net/http"
)

func main() {
    resp, err := http.Get("https://example.com")
    if err == nil && resp.StatusCode == 200 {
        fmt.Println("Site accessible")
    } else {
        fmt.Println("Site indisponible")
    }
}
