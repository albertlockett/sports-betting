package fivethirtyeight

import (
  "encoding/csv"
  "github.com/albertlockett/sports-betting/sources/model"
  "net/http"
  "time"
)

const FILE_URL = "https://projects.fivethirtyeight.com/mlb-api/mlb_elo_latest.csv"

func FetchEvents() ([]*model.Event, error) {
  data, err := readFromCsv()
  if err != nil {
    return nil, err
  }

  results, err := unmarshalCsvData(data)
  if err != nil {
    return nil, err
  }

  return results, nil
}

// download the raw data
func readFromCsv() ([][]string, error) {
  resp, err := http.Get(FILE_URL)
  if err != nil {
    return nil, err
  }

  defer resp.Body.Close()
  reader := csv.NewReader(resp.Body)
  reader.Comma = ','
  data, err := reader.ReadAll()
  if err != nil {
    return nil, err
  }

  return data, nil
}

func unmarshalCsvData(data [][]string) ([]*model.Event, error) {
  layout := "2006-01-02"
  today := time.Now().Format(layout)

  results := make([]*model.Event, 0)
  for i, row := range(data) {
    if i == 0 {
      continue // skip header
    }

    // only return todays games
    timeRaw := row[0]
    if timeRaw != today {
      continue
    }

    gameTime, err := time.Parse(layout, timeRaw)
    if err != nil {
      return nil, err
    }

    event := model.Event{
      HomeTeam: row[4],
      AwayTeam: row[5],
      Time: gameTime,
    }
    results = append(results, &event)
  }

  return results, nil
}