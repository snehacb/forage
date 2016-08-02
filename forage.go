package main

import (
	"fmt"

	gowit "github.com/snehacb/forage/gowit"
	parser "github.com/snehacb/forage/parser"
)

func main() {
	fmt.Println("Forager at your service!")
	fmt.Println("Send user message to wit.ai app")
	witrequest := gowit.PingBot("hi")
	witresponse := parser.ProcessJsonRequest(witrequest)
	if !witresponse {
		fmt.Println("We are done here!")
	}
}
