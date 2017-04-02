package cmd

import (
  "fmt"
  "encoding/base64"

  "github.com/urfave/cli"
)

func Base64Action(c *cli.Context) {
  isDecode := c.Bool("decode")
  if isDecode {
    DecodeAction(c)
  } else {
    EncodeAction(c)
  }
}

func EncodeAction(c *cli.Context) {
  var data = ""
  if len(c.Args()) > 0 {
    data = c.Args().First()
  }
  encodedStr := base64.StdEncoding.EncodeToString([]byte(data))
  fmt.Println(encodedStr)
}

func DecodeAction(c *cli.Context) {
  var data = ""
  if len(c.Args()) > 0 {
    data = c.Args().First()
  }
  decodedStr, _ := base64.StdEncoding.DecodeString(data)
  fmt.Println(string(decodedStr))
}
