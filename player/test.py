import requests
import json

data = {'name': 'phuongaz', "join_date": "10-10-2022", "coin": 100}
url = "http://localhost:8080/register"
headers = {'Content-type': 'application/json', 'Accept': 'text/plain'}
r = requests.post(url, data=json.dumps(data), headers=headers)

print(r.text)

get = requests.get("http://localhost:8080/player/phuongaz")
print(get.text)
