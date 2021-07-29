package config

import (
  "os"
)

var Local Config

const CFG_ES_URL = "CFG_ES_URL"

func Init() {
  Local = Config{
    EsUrl: os.Getenv(CFG_ES_URL),
  }
}

type Config struct {
  EsUrl string
}