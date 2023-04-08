# saiStorageMongo

### API
#### Get
- request:

curl --location --request GET 'http://localhost:8801/get' \
&emsp;    --header 'Token: $SomeToken' \
&emsp;    --header 'Content-Type: application/json' \
&emsp;    --data-raw '{"collection":"$CollectionName",select":{"$KeyToSelect":"$ValueToSelect"}}'

- response example: {"result":[{"_id":"6431a25da6123e99c1598432","ch_time":1680974429,"cr_time":1680974429,"internal_id":"37f06418-34d3-4cff-9eb4-24d52786b371","sssss":"636037b417cde3a8fea98735"}]}

#### Save
- request:

curl --location --request GET 'http://localhost:8801/save' \
&emsp;    --header 'Token: $SomeToken' \
&emsp;    --header 'Content-Type: application/json' \
&emsp;    --data-raw '{"collection":"$CollectionName","data":{"KeyToSave":"$ValueToSave"}}'

- response: {"Status":"Ok"} 

#### Update
- request:

curl --location --request GET 'http://localhost:8801/update' \
&emsp;    --header 'Token: $SomeToken' \
&emsp;    --header 'Content-Type: application/json' \
&emsp;    --data-raw '{"collection":"$CollectionName","data":{"$Key":"$Value"}}'

- response: {"Status":"Ok"} 


#### Upsert
- request:

curl --location --request GET 'http://localhost:8801/upsert' \
&emsp;    --header 'Token: $SomeToken' \
&emsp;    --header 'Content-Type: application/json' \
&emsp;    --data-raw '{"collection":"$CollectionName","select":{"$Key":$Value},"data":{"$inc":{"$Key":$Value}}}'

- response: {"Status":"Ok"} 

#### Remove
- request:

curl --location --request GET 'http://localhost:8801/remove' \
&emsp;    --header 'Token: $SomeToken' \
&emsp;    --header 'Content-Type: application/json' \
&emsp;    --data-raw '{"collection":"$CollectionName","select":{"$Keys":"$Value"}}'

- response: {"Status":"Ok"} 

### Run in Docker
`make up`

### Run as standalone application
`microservices/saiStorageMongo/build/sai-storage` 


## Profiling
 host:port/debug/pprof`