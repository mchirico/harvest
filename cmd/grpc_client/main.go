package main

import (
	"github.com/levigross/grequests"
	"log"
)

// go get -u github.com/levigross/grequests
// Test: http://httpbin.org/post
func main() {

	json := []byte(`{"method": "JSONServer.GiveBookDetail",
                 "params": [{
                 "Id": "1234"
                           }],
                  "id": "1"

}`)

	resp, err := grequests.Post("http://localhost:8082/rpc",
		&grequests.RequestOptions{
			JSON: json,
		})

	if err != nil {
		log.Println("Cannot post: ", err)
	}

	if resp.Ok != true {
		log.Println("Request did not return OK")
	} else {
		log.Println("resp:\n", resp)
	}

}
