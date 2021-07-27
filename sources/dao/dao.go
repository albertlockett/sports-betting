package dao

import (
  "context"
  "encoding/json"
  "errors"
  "fmt"
  "github.com/albertlockett/sports-betting/sources/model"
  "github.com/elastic/go-elasticsearch/v7"
  "github.com/elastic/go-elasticsearch/v7/esapi"
  "log"
  "strings"
)

const IDX_HANDICAP = "handicaps"
const IDX_LINES = "lines"

type dao struct {
  client *elasticsearch.Client
}

var local = dao{}

func Init() error {
  config := elasticsearch.Config{
    Addresses: []string{
      "http://localhost:9200",
    },
    //Transport: &LoggingTransport{},
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
  var r  map[string]interface{}
  if err := json.NewDecoder(res.Body).Decode(&r); err != nil {
    log.Fatalf("Error parsing the response body: %s", err)
  }
  // Print client and server version numbers.
  log.Printf("Client: %s", elasticsearch.Version)
  log.Printf("Server: %s", r["version"].(map[string]interface{})["number"])
  log.Println("Connection success")

  return nil
}

func SaveHandicap(handicap *model.Handicap) error {
  document, err := json.Marshal(handicap);
  if err != nil {
    return err
  }
  return saveRecord(IDX_HANDICAP, handicap.ComputeId(), document)
}

func SaveLine(line *model.Line) error {
  document, err := json.Marshal(line);
  if err != nil {
    return err
  }
  return saveRecord(IDX_LINES, line.ComputeId(), document)
}

func saveRecord(index string, id string, document []byte) error {
  req := esapi.IndexRequest{
    Index: index,
    DocumentID: id,
    Body: strings.NewReader(string(document)),
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

func resetLatestCollectedFlag(index string) error {
  req := esapi.UpdateByQueryRequest{
    Index: []string{index},
    Body: strings.NewReader(`{
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


