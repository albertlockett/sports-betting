package main

import (
  "encoding/json"
  "fmt"
  "github.com/albertlockett/sports-betting/sources/config"
  "github.com/albertlockett/sports-betting/sources/dao"
  "github.com/gorilla/mux"
  "io/ioutil"
  "log"
  "net/http"
)

type reqBody struct {
  LatestCollected bool
  Time            string
}

func queryFromBody(req *http.Request) (dao.Query, error) {
  bodyBytes, err := ioutil.ReadAll(req.Body)
  if err != nil {
    return nil, err
  }
  if len(bodyBytes) == 0 {
    return &dao.MatchAllQuery{}, nil
  }

  rb := reqBody{}
  if err := json.Unmarshal(bodyBytes, &rb); err != nil {
    return nil, err
  }

  queries := make([]dao.Query, 0)
  if rb.LatestCollected {
    queries = append(queries, dao.TermQuery{Field: "LatestCollected", Value: true,})
  }

  if rb.Time != "" {
    queries = append(queries, dao.TermQuery{Field: "Time", Value: rb.Time})
  }

  var query dao.Query = dao.BoolQuery{
    Must: queries,
  }
  return query, nil
}

func sendErrorResponse(w http.ResponseWriter) {
  w.WriteHeader(500)
  w.Write([]byte(`{ "error": "Internal Server Error" }`))
}

func getExpectedValues(w http.ResponseWriter, r *http.Request) {
  w.Header().Set("Content-Type", "application/json")

  query, err := queryFromBody(r)
  if err != nil {
    sendErrorResponse(w)
    return
  }

  results, err := dao.SearchExpectedValues(&dao.SearchRequestBody{Query: query})
  if err != nil {
    sendErrorResponse(w)
    return
  }

  bytes, err := json.Marshal(results)
  if err != nil {
    sendErrorResponse(w)
    return
  }

  w.Write(bytes)
}

func main() {
  if err := config.Init(); err != nil {
    panic(err)
  }
  if err := dao.Init(); err != nil {
    panic(err)
  }
  r := mux.NewRouter()
  r.HandleFunc("/expected-values", getExpectedValues).Methods("POST")
  log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", config.Local.Port), r))
}
