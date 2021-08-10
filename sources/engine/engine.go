package engine

import (
  "github.com/albertlockett/sports-betting/sources/dao"
  "github.com/albertlockett/sports-betting/sources/model"
  "time"
)

func CalculateDailySummary() (*model.DailySummary, error) {
  time := time.Now().Truncate(24 * time.Hour)

  query := dao.BoolQuery{
    Must: []dao.Query{
      dao.TermQuery{
        Field: "Time",
        Value: time.Format("2006-01-02") + "T00:00:00Z",
      },
    },
  }

  evs, err := dao.SearchExpectedValues(&dao.SearchRequestBody{
    Query: query,
    Size:  10000,
  })

  if err != nil {
    return nil, err
  }

  return &model.DailySummary{
    Time:     time,
    NumGames: uint8(len(evs)),
  }, err
}
