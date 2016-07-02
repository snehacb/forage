package gowit

import (
"fmt"
gorequest "github.com/parnurzeal/gorequest"
)

var url string

func GetRequest() {
  request := gorequest.New()
  resp, body, errs := request.Get(url).End()
  if resp != nil {
    fmt.Println(body)
  } else if errs != nil {
    fmt.Println(errs)
  }
}

func init() {
  url = "https://hub.docker.com/v2/search/repositories/?page=1&query=sc&is_official=1&page_size=20"
}
