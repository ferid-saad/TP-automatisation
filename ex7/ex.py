import smtplib

server = smtplib.SMTP('smtp.example.com', 587)
server.starttls()
server.login("user@example.com", "password")
server.sendmail("from@example.com", "to@example.com", 
    "Subject: Alerte\n\nMessage d'alerte")