package main

import (
  "fmt"
  "github.com/albertlockett/sports-betting/sources/dao"
  "log"
)

func main() {
  err := dao.Init()
  if err != nil {
    panic(err)
  }

  results, err := FindBetEVs()
  log.Println(fmt.Sprintf("There were %s evs", len(results)))

  for _, result := range results {
    log.Println(fmt.Sprintf(
      "%s @ %s : side: %s, handicap: %f, line %f, ev %f",
      result.AwayTeam,
      result.HomeTeam,
      result.Handicap.Side,
      result.Odds,
      result.LineDecimal,
      result.ExpectedValue,
    ))
  }
}
