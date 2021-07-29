package dao

import (
  "context"
  "encoding/json"
  "errors"
  "fmt"
  config2 "github.com/albertlockett/sports-betting/sources/config"
  "github.com/albertlockett/sports-betting/sources/model"
  "github.com/elastic/go-elasticsearch/v7"
  "github.com/elastic/go-elasticsearch/v7/esapi"
  "log"
  "sort"
  "strings"
)

const IDX_HANDICAP = "handicaps"
const IDX_LINES = "lines"
const IDX_EXPECTED_VALUES = "expected-values"

type dao struct {
  client *elasticsearch.Client
}

var local = dao{}

func Init() error {
  config := elasticsearch.Config{
    Addresses: []string{
      config2.Local.EsUrl,
    },
    Transport: &LoggingTransport{},
  }
  es, err := elasticsearch.NewClient(config)
  if err != nil {
    return err
  }

  local.client = es

  err = testConnection()
  if err != nil {
    return err
  }

  return nil
}

func testConnection() error {
  log.Println("Testing database connection")
  res, err := local.client.Info()
  if err != nil {
    log.Fatalf("Error getting response: %s", err)
  }
  defer res.Body.Close()
  // Check response status
  if res.IsError() {
    log.Fatalf("Error: %s", res.String())
  }
  // Deserialize the response into a map.
  var r map[string]interface{}
  if err := json.NewDecoder(res.Body).Decode(&r); err != nil {
    log.Fatalf("Error parsing the response body: %s", err)
  }
  // Print client and server version numbers.
  log.Printf("Client: %s", elasticsearch.Version)
  log.Printf("Server: %s", r["version"].(map[string]interface{})["number"])
  log.Println("Connection success")

  return nil
}

func SaveExpectedValue(ev *model.ExpectedValue) error {
  document, err := json.Marshal(ev)
  if err != nil {
    return err
  }
  return saveRecord(IDX_EXPECTED_VALUES, ev.ComputeId(), document)
}

func SaveHandicap(handicap *model.Handicap) error {
  document, err := json.Marshal(handicap)
  if err != nil {
    return err
  }
  return saveRecord(IDX_HANDICAP, handicap.ComputeId(), document)
}

func SaveLine(line *model.Line) error {
  document, err := json.Marshal(line)
  if err != nil {
    return err
  }
  return saveRecord(IDX_LINES, line.ComputeId(), document)
}

func saveRecord(index string, id string, document []byte) error {
  req := esapi.IndexRequest{
    Index:      index,
    DocumentID: id,
    Body:       strings.NewReader(string(document)),
  }

  res, err := req.Do(context.Background(), local.client)
  if err != nil {
    return err
  }
  defer res.Body.Close()

  if res.IsError() {
    errMsg := fmt.Sprintf("[%s] Error indexing document", res.Status())
    err = errors.New(errMsg)
    return err
  }

  return nil
}

func ResetLineLatestCollected() error {
  return resetLatestCollectedFlag(IDX_LINES)
}

func ResetHandicapLatestCollected() error {
  return resetLatestCollectedFlag(IDX_HANDICAP)
}

func ResetExpectedValueLatestCollected() error {
  return resetLatestCollectedFlag(IDX_EXPECTED_VALUES)
}

func resetLatestCollectedFlag(index string) error {
  req := esapi.UpdateByQueryRequest{
    Index: []string{index},
    Body: strings.NewReader(`{
      "size": 1000,
      "query":{
        "term":{
          "LatestCollected":{"value":"true"}
        }
      },
      "script":{
        "source": "ctx._source.LatestCollected = false",
        "lang": "painless"
      }
    }`),
  }

  res, err := req.Do(context.Background(), local.client)
  if err != nil {
    return err
  }
  defer res.Body.Close()

  if res.IsError() {
    errMsg := fmt.Sprintf("[%s] Error resetting LastCollected document", res.Status())
    err = errors.New(errMsg)
    return err
  }

  return nil
}

type fetchHandicapsResult struct {
  Hits struct {
    Hits []struct {
      Source model.Handicap `json:"_source"`
    }
  }
}

type fetchLinesResult struct {
  Hits struct {
    Hits []struct {
      Source model.Line `json:"_source"`
    }
  }
}

func FetchEvents() ([]*model.ExpectedValue, error) {
  // TODO make times the same aka today

  req := esapi.SearchRequest{
    Index: []string{IDX_HANDICAP},

    //
    Body: strings.NewReader(`{
      "size": 1000,
      "query": {
        "bool": {
          "must": [
            { "term": { "LatestCollected":{"value":"true"} } }
          ]
        }
      }
    }`),
  }

  res, err := req.Do(context.Background(), local.client)
  if err != nil {
    return nil, err
  }
  defer res.Body.Close()

  resBody := &fetchHandicapsResult{}
  if err := json.NewDecoder(res.Body).Decode(&resBody); err != nil {
    log.Fatalf("Error parsing the response body: %s", err)
  }

  if err != nil {
    return nil, err
  }

  handicaps := make([]*model.Handicap, 0)
  for _, result := range resBody.Hits.Hits {
    handicaps = append(handicaps, &result.Source)
  }

  results := make([]*model.ExpectedValue, 0)

  for _, hit := range resBody.Hits.Hits {
    handicap := hit.Source
    query := fmt.Sprintf(`{
      "query": {
        "bool": {
          "must": [
            { "term": { "LatestCollected" : { "value": "true" } } },
            { "term": { "HomeTeam.keyword": { "value":"%s" } } },
            { "term": { "AwayTeam.keyword": { "value": "%s" } } },
            { "term": { "Side": { "value": "%s" } } }
          ]
        }
      }
    }`, handicap.HomeTeam, handicap.AwayTeam, handicap.Side)
    fmt.Println(query)

    req := esapi.SearchRequest{
      Index: []string{IDX_LINES},
      Body:  strings.NewReader(query),
    }
    res, err := req.Do(context.Background(), local.client)
    if err != nil {
      return nil, err
    }
    defer res.Body.Close()

    resBody := &fetchLinesResult{}
    if err := json.NewDecoder(res.Body).Decode(&resBody); err != nil {
      log.Fatalf("Error parsing the response body: %s", err)
    }

    for _, hit := range resBody.Hits.Hits {
      line := hit.Source
      ev := model.ExpectedValue{
        Event:           line.Event,
        Line:            line,
        Handicap:        handicap,
        LatestCollected: true,
        Side:            line.Side,
        ExpectedValue:   handicap.Odds * line.LineDecimal,
      }
      results = append(results, &ev)
    }
  }

  sort.Slice(results[:], func(i, j int) bool {
    return results[i].ExpectedValue > results[j].ExpectedValue
  })

  return results, nil
}
