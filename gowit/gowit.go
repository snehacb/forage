package gowit

import (
"fmt"
gorequest "github.com/parnurzeal/gorequest"
"os"

)

var url string
var token string

func init() {
  url = "https://api.wit.ai/message?v=20160702&q="
  token = os.Getenv("WITTOKEN")
}

func PingBot(msg string) string {
  request := gorequest.New()
  /* initialize token and url */
  tokenstr := "Bearer "
  tokenstr += token
  urlstr := url
  urlstr += msg
  fmt.Printf("Token = %v, Url Name = %v\n", tokenstr, urlstr)
  /* send message from user to wit.ai app */
  resp, body, errs := request.Get(urlstr).Set("Authorization", tokenstr).End()

  if resp != nil {
    fmt.Printf("Response %T, %v \n", resp, resp)
  } else if errs != nil {
    fmt.Printf("Errors %T, %v \n", errs, errs)
  }
  return body
}

func ProcessUserMessage(msg string) {

}
