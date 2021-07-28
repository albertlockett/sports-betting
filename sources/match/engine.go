package main

import (
  "github.com/albertlockett/sports-betting/sources/dao"
  "github.com/albertlockett/sports-betting/sources/model"
)

func FindBetEVs() ([]*model.ExpectedValue, error) {
  return dao.FetchEvents()
}