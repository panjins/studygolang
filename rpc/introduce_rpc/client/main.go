package main

import (
	"fmt"
	"github.com/kirinlabs/HttpRequest"
)

func main() {
	req := HttpRequest.NewRequest()
	res, err := req.Get("http://127.0.0.1:8000/add?a=1&b=2")
	if err != nil {
		panic(err)
	}
	body, _ := res.Body()
	fmt.Println(string(body))

}
