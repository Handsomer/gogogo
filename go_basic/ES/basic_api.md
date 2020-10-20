## ElasticSearch api use
###安装插件 ./bin/elasticsearch-plugin install analysis-icu
`
import json
import requests
headers = {"Connection":"keep-alive","Accept-Encoding":"gzip, deflate, br","Accept":"*/*","User-Agent":"PostmanRuntime/7.26.3","Content-Type":"application/json","Cache-Control":"no-cache"}
ret = requests.get("http://localhost:9200/_analyze", data=json.dumps({"analyzer": "icu_analyzer","text": "他说的确实在理"}), headers=headers)
`

