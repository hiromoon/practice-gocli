package main

import (
  "fmt"
  "os"

  "github.com/urfave/cli"
  "github.com/joho/godotenv"

  "./cmd"
)

func main() {
  app := cli.NewApp()
  app.Name = "boom"
  app.Usage = "make an explosive entrance"
  app.Version = "0.0.1"

  app.Flags = []cli.Flag{
          cli.BoolFlag{
            Name: "urlsafe, u",
            Usage: "URLSafe flag",
          },
  }

  app.Commands = []cli.Command{
        {
           Name: "base64",
           Aliases: []string{"b"},
           Usage: "base64 encoding(decoding) command",
           Action: cmd.Base64Action,
           Flags: []cli.Flag{
             cli.BoolFlag{
               Name: "decode, d",
               Usage: "decoding base64",
             },
           },
        },
        {
           Name: "config",
           Usage: "go tool config command",
           Action: cmd.ConfigAction,
        },
        {
           Name: "get",
           Usage: "http get request",
           Action: cmd.GetAction,
        },
  }

  app.Before = func(c *cli.Context) error {
    fmt.Println("Start")
    return nil
  }

  app.After = func(c *cli.Context) error {
    fmt.Println("Finish")
    return nil
  }
  loadEnv()
  app.Run(os.Args)
}

func loadEnv() {
  err := godotenv.Load()
  if err != nil {
    fmt.Println("Error loading .env file")
  }
}
