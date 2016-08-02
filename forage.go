package forager

import (
"fmt"
gowit "github.com/snehacb/forage/gowit"
forager "github.com/snehacb/forage/forager"
)

func main() {
  fmt.Println("Forager at your service!")
  fmt.Println("Send user message to wit.ai app")
  witrequest := gowit.SendUserMessage("hi")
  witresponse := forager.ProcessJsonRequest(witrequest)
  if !witresponse {
    fmt.Println("We are done here!")
  }
}
