package main

/*
curl -X POST \
   http://localhost:8082/rpc \
   -H 'cache-control: no-cache' \
   -H 'content-type: application/json' \
   -d '{
   "method": "JSONServer.GiveBookDetail",
   "params": [{
   "Id": "1234"
   }],
   "id": "1"
}'




curl -X POST \
   http://httpbin.org/post \
   -H 'cache-control: no-cache' \
   -H 'content-type: application/json' \
   -d '{
   "method": "JSONServer.GiveBookDetail",
   "params": [{
   "Id": "1234"
   }],
   "id": "1"
}'




 */


import (
	"github.com/mchirico/harvest/rpkg"
)

func main() {
	a := rpkg.App{}
	a.Initilize()
	a.Run("8082", 15, 15)
}
