## ElasticSearch api use
`
import json
import requests
ret = requests.get("http://localhost:9200/_analyze", data=json.dumps({"analyzer": "english","text": "this word is hello world"}), headers=headers)
headers = {"Connection":"keep-alive","Accept-Encoding":"gzip, deflate, br","Accept":"*/*","User-Agent":"PostmanRuntime/7.26.3","Content-Type":"application/json","Cache-Control":"no-cache"}
ret = requests.get("http://localhost:9200/_analyze", data=json.dumps({"analyzer": "english","text": "this word is hello world"}), headers=headers)
`

