package main

import (
  "fetch-app/application"
  "fetch-app/application/registry"
  "flag"
  "fmt"
)

func main() {

  appMap := map[string]func() application.RegistryContract{
    "fetchapp": registry.NewFetchapp(),
  }

  flag.Parse()

  app, exist := appMap[flag.Arg(0)]
  if exist {
    application.Run(app())
  } else {
    fmt.Println("You may try 'go run main.go <app_name>' :")
    for appName := range appMap {
      fmt.Printf(" - %s\n", appName)
    }
  }

}
