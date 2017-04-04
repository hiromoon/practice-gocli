package cmd

import (
  "os"
  "fmt"

  "github.com/urfave/cli"
)

func ConfigAction(c* cli.Context) {
  message := fmt.Sprintf("user_name=%s password=%s", os.Getenv("USER_NAME"), os.Getenv("PASSWORD"))
  fmt.Println(message)
}
