package main

import (
    "net/smtp"
)

func main() {
    msg := []byte("Subject: Alerte système\r\n\r\nAlerte: problème détecté.\r\n")
    smtp.SendMail("smtp.gmail.com:587",
        smtp.PlainAuth("", "admin@example.com", "motdepasse", "smtp.gmail.com"),
        "admin@example.com", []string{"user@example.com"}, msg)
}
