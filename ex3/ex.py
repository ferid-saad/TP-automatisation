import requests

r = requests.get('https://example.com')
if r.status_code == 200:
    print("Site accessible")
else:
    print("Site indisponible")