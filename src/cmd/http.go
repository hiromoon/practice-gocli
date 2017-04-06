package cmd

import (
  "fmt"
  "net/http"
  "io/ioutil"

  "github.com/urfave/cli"
)

func GetAction(c* cli.Context) {

  req, err := http.NewRequest("GET", "https://www.google.com", nil)
  if err != nil {
    fmt.Println(err)
  }

  client := &http.Client{}
  resp, err := client.Do(req)
  if err != nil {
    fmt.Println(err)
  }
  fmt.Println(resp)
  defer resp.Body.Close()

  b, err := ioutil.ReadAll(resp.Body)
  if err == nil {
    fmt.Println(string(b))
  }
}
