package dao

import (
  "context"
  "encoding/json"
  "fmt"
  "github.com/albertlockett/sports-betting/sources/model"
  "github.com/elastic/go-elasticsearch/v7/esapi"
  "strings"
)

type Query interface {
  Build() (map[string]interface{}, error)
}

type BoolQuery struct {
  Must []Query
}

func (t BoolQuery) Build() (map[string]interface{}, error) {
  mb := make([]map[string]interface{}, 0)
  for _, m := range t.Must {
    if b, err := m.Build(); err != nil {
      return nil, err
    } else {
      mb = append(mb, b)
    }
  }
  
  return map[string]interface{}{
    "bool": map[string]interface{}{
      "must": mb,
    },
  }, nil
}

type TermQuery struct {
  Field string
  Value interface{}
}

func (t TermQuery) Build() (map[string]interface{}, error) {
  return map[string]interface{}{
    "term": map[string]interface{}{
      t.Field: map[string]interface{}{
        "value": t.Value,
      },
    },
  }, nil
}

type MatchAllQuery struct{}

func (m MatchAllQuery) Build() (map[string]interface{}, error) {
 return map[string]interface{} {
   "match_all": map[string]interface{}{},
 }, nil
}

type SearchRequestBody struct {
  Size  int64
  Query Query
}

func stringifyQuery(query Query) ([]byte, error) {
  data, err := query.Build()
  if err != nil {
    return nil, err
  }
  return json.Marshal(data)
}

func Search(index string, search *SearchRequestBody) (*esapi.Response, error) {
  q, err := stringifyQuery(search.Query)
  if err != nil {
    return nil, err
  }

  size := search.Size
  if size == 0 {
    size = 100
  }

  body := fmt.Sprintf(`{
    "size": %d,
    "query":  %s
  }`, size, q)

  req := esapi.SearchRequest{
    Index: []string{index},
    Body:  strings.NewReader(body),
  }
  res, err := req.Do(context.Background(), local.client)
  if err != nil {
    return nil, err
  }
  defer res.Body.Close()
  return res, nil
}

func SearchExpectedValues(search *SearchRequestBody) ([]*model.ExpectedValue, error) {
  res, err := Search(IDX_EXPECTED_VALUES, search)
  if err != nil {
    return nil, err
  }

  resBody := struct {
    Hits struct {
      Hits []struct {
        Source model.ExpectedValue `json:"_source"`
      }
    }
  }{}
  if err := json.NewDecoder(res.Body).Decode(&resBody); err != nil {
    return nil, err
  }

  evs := make([]*model.ExpectedValue, 0)
  for i, _ := range resBody.Hits.Hits {
    //evs = append(evs, &result.Source)
    evs = append(evs, &resBody.Hits.Hits[i].Source)
  }

  return evs, nil
}

func SearchDailySumamrys(search *SearchRequestBody) ([]*model.DailySummary, error) {
  res, err := Search(IDX_DAILY_SUMMARY, search)
  if err != nil {
    return nil, err
  }

  resBody := struct {
    Hits struct {
      Hits []struct {
        Source model.DailySummary `json:"_source"`
      }
    }
  }{}
  if err := json.NewDecoder(res.Body).Decode(&resBody); err != nil {
    return nil, err
  }

  evs := make([]*model.DailySummary, 0)
  for i, _ := range resBody.Hits.Hits {
    evs = append(evs, &resBody.Hits.Hits[i].Source)
  }

  return evs, nil
}