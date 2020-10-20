## ElasticSearch api use
###安装插件 ./bin/elasticsearch-plugin install analysis-icu
`
import json
import requests
headers = {"Connection":"keep-alive","Accept-Encoding":"gzip, deflate, br","Accept":"*/*","User-Agent":"PostmanRuntime/7.26.3","Content-Type":"application/json","Cache-Control":"no-cache"}
data = {"analyzer": "icu_analyzer","text": "他说的确实在理"}
ret = requests.get("http://localhost:9200/_analyze", data=json.dumps(), headers=headers)
`

### search:
 `get kibana_sample_data_ecommerce/_search?q=EUR
             {
               "profile":"true"
             }`
`

`get /_cat/nodes
 get kibana_sample_data_ecommerce
 get _cluster/health
 
 put users/_create/1
 {
 	"firstName":"Jack",
 	"lastName":"Johnson",
 	"tags":["guitar", "skateboard"]
 }
 
 get users/_doc/1
 
 put users/_doc/1
 {
 	"tags":["guitar", "skateboard", "reading"]
 }
 
 post users/_update/1/
 {
 	"doc":{
 	"post_date":"2020-10-19T17:58:12",
 	"message":"trying out Elastixsearch"
 	}
 }
 
 GET _analyze?
 {"analyzer": "english",
 "text": "this word is hello world"
 }
 
 
 
 get kibana_sample_data_ecommerce
 
 //泛查询，对所有字段进行检索
 get kibana_sample_data_ecommerce/_search?q=EUR
 {
   "profile":"true"
 }
 
 //指定字段currency
 get kibana_sample_data_ecommerce/_search?q=currendcy:EUR
 {
   "profile":"true"
 }
 
 //bool 查询
 get kibana_sample_data_ecommerce/_search?q=customer_full_name:(Eddie %2BWeber)
 {
   "profile":"true"
 }
 
 //AND
 get kibana_sample_data_ecommerce/_search?q=customer_full_name:(Eddie AND Weber)
 {
   "profile":"true"
 }
 
 //OR
 get kibana_sample_data_ecommerce/_search?q=customer_full_name:(Eddie OR Weber)
 {
   "profile":"true"
 }
 
 //NOT
 get kibana_sample_data_ecommerce/_search?q=customer_full_name:(Eddie NOT Weber)
 {
   "profile":"true"
 }
 
 //相似查询 模糊查询
 get kibana_sample_data_ecommerce/_search?q=order_date:"2020"~2
 {
   "profile":"true"
 }
 
 get kibana_sample_data_ecommerce/_search?q=created_on:"2016-12-26T09:28:48+00:00"
 {
   "profile":"true"
 }
 
 get kibana_sample_data_ecommerce/_search?q=category:Men~2
 {
   "profile":"true"
 }
 
 //通配符查询
 get kibana_sample_data_ecommerce/_search?q=category:Men*
 {
   "profile":"true"
 }
 
 

 
`