package cmd

import (
  "fmt"
  "encoding/base64"

  "github.com/urfave/cli"
)

func EncodeAction(c *cli.Context) {
  var paramFirst = ""
  if len(c.Args()) > 0 {
    paramFirst = c.Args().First()
  }
  encodedStr := base64.StdEncoding.EncodeToString([]byte(paramFirst))
  fmt.Println(encodedStr)
}

func DecodeAction(data string) {
  decodedStr, _ := base64.StdEncoding.DecodeString(data)
  fmt.Println(decodedStr)
}
