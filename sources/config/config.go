package config

import (
  "errors"
  "fmt"
  "os"
  "strings"
)

var Local Config

const CFG_ES_URL = "CFG_ES_URL"
const CFG_PORT = "CFG_PORT"

func Init() error {
  Local = Config{
    EsUrl: os.Getenv(CFG_ES_URL),
    Port: os.Getenv(CFG_PORT),
  }
  return validate(Local)
}

func validate(config Config) error {
  errs := make([]string, 0)

  if config.Port == "" {
    errs = append(errs, fmt.Sprintf("%s must be set", CFG_PORT))
  }

  if len(errs) > 0 {
    err := errors.New(strings.Join(errs, "\n"))
    return err
  }

  return nil
}

type Config struct {
  EsUrl string
  Port string
}