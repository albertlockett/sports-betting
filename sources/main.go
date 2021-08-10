package main

import (
  "fmt"
  "github.com/albertlockett/sports-betting/sources/betdsi"
  "github.com/albertlockett/sports-betting/sources/config"
  "github.com/albertlockett/sports-betting/sources/dao"
  "github.com/albertlockett/sports-betting/sources/engine"
  "github.com/albertlockett/sports-betting/sources/fivethirtyeight"
  "log"
  "os"
)

func main() {
  config.Init()
  err := dao.Init()
  if err != nil {
    panic(err)
  }

  command := os.Args[1]

  switch command {
  case "expected-values":
    fetchExpectedValues()
  case "handicaps":
    fetchHandicaps()
  case "lines":
    fetchLines()
  case "daily-games":
    fetchDailySummarys()
  }
}

func fetchHandicaps() {
  handicaps, err := fivethirtyeight.FetchEvents()
  if err != nil {
    panic(err)
  }
  log.Println(fmt.Sprintf("There were %d events", len(handicaps)/2))

  for _, handicap := range handicaps {
    handicap.EventId = handicap.ComputeEventId()
    if err := dao.ResetHandicapLatestCollected(handicap.EventId); err != nil {
      panic(err)
    }
    if err := dao.SaveHandicap(handicap); err != nil {
      panic(err)
    }
  }
}

func fetchLines() {
  lines, err := betdsi.FetchLines()
  if err != nil {
    log.Println("an error happened")
    panic(err)
  }
  log.Println(fmt.Sprintf("There were %d lines", len(lines)/2))

  for _, line := range lines {
    line.EventId = line.ComputeEventId()
    if err := dao.ResetLineLatestCollected(line.EventId); err != nil {
      panic(err)
    }
    if err := dao.SaveLine(line); err != nil {
      panic(err)
    }
  }
}

func fetchExpectedValues() {
  results, err := dao.FetchEvents()
  if err != nil {
    panic(err)
  }

  for _, result := range results {
    result.EventId = result.ComputeEventId()
    if err := dao.ResetExpectedValueLatestCollected(result.EventId); err != nil {
      panic(err)
    }
    if err := dao.SaveExpectedValue(result); err != nil {
      panic(err)
    }
  }
}

func fetchDailySummarys () {
  ds, err := engine.CalculateDailySummary()
  if err != nil {
    panic(err)
  }
  if err = dao.SaveDailyValue(ds); err != nil {
    panic(err)
  }
}
