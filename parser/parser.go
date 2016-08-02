package parser

import (
	"fmt"
)

/* SAMPLE JSON REQUEST
string, {
  "msg_id" : "084a2ad7-eee0-4b28-9be4-7203d363c1eb",
  "_text" : "hi",
  "entities" : {
    "intent" : [ {
      "confidence" : 0.9642521445544848,
      "value" : "greeting"
    } ]
  }
} */

type Value struct {
	Expressions []string `json:"expressions"`
	Value       string   `json:"value"`
}

type Entity struct {
	Lang    string   `json:"lang"`
	Closed  bool     `json:"closed"`
	Exotic  bool     `json:"exotic"`
	Value   string   `json:"value"`
	Values  []*Value `json:"values"`
	Builtin bool     `json:"builtin"`
	Doc     string   `json:"doc"`
	Name    string   `json:"name"`
	ID      string   `json:"id"`
}

type Message struct {
	ID       string    `json:"msg_id"`
	Text     string    `json:"_text"`
	Entities []*Entity `json:"entities"`
}

func init() {

}

func ProcessJsonRequest(request string) bool {
	if request != "" {
		fmt.Printf("To Marshal: %v", request)
	}
	return false
}
