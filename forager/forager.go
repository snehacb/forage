package forager

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
type Json struct {

}

func init() {

}

func ProcessJsonRequest( request string) bool{
  if request != "" {
    fmt.Println(request)
  }
  return false
}
