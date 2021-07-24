package dao

import (
  "encoding/json"
  "github.com/elastic/go-elasticsearch/v7"
  "log"
  "strings"
)

type dao struct {
  client *elasticsearch.Client
}

var local = dao{}

func Init() error {
  config := elasticsearch.Config{
    Addresses: []string{
      "http://localhost:9200",
    },
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
  log.Println(strings.Repeat("~", 37))

  return nil
}