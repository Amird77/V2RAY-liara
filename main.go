package main

import (
  "log"
  "os"
  "github.com/v2ray/v2ray-core/v5/core"
  "github.com/v2ray/v2ray-core/v5/infra/conf"
)

func main() {
  configFile := "config.json"
  if len(os.Args) > 1 {
    configFile = os.Args[1]
  }

  configContent, err := os.ReadFile(configFile)
  if err != nil {
    log.Fatal("Cannot read config file:", err)
  }

  config, err := conf.LoadConfig(configContent)
  if err != nil {
    log.Fatal("Cannot load config:", err)
  }

  server, err := core.New(config)
  if err != nil {
    log.Fatal("Cannot create V2Ray server:", err)
  }

  if err := server.Start(); err != nil {
    log.Fatal("Cannot start server:", err)
  }

  log.Println("V2Ray server started")
  select {}
}
