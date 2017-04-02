package main

import (
  "fmt"
  "os"

  "github.com/urfave/cli"

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
           Name: "hello",
           Aliases: []string{"h"},
           Usage: "print 'Hello, World!!'",
           Action: helloAction,
        },
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
  }

  app.Before = func(c *cli.Context) error {
    fmt.Println("Start")
    return nil
  }

  app.After = func(c *cli.Context) error {
    fmt.Println("Finish")
    return nil
  }

  app.Run(os.Args)
}

func helloAction(c *cli.Context) {
  var isDry = c.GlobalBool("dryrun")
  if isDry {
    fmt.Println("this is dry-run")
  }
  var paramFirst = ""
  if len(c.Args()) > 0 {
    paramFirst = c.Args().First()
  }

  fmt.Printf("Hello World!! %s\n", paramFirst)
}
