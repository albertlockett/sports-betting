package model

import (
  "crypto/sha1"
  "encoding/base64"
  "fmt"
  "time"
)

const SIDE_HOME = "home"
const SIDE_AWAY = "away"

const _ID_DATE_LAYOUT = "yyyy-MM-dd"

type Event struct {
  HomeTeam string
  AwayTeam string
  Time     time.Time
}

type Handicap struct {
  Event
  TimeCollected   time.Time
  LatestCollected bool
  Odds            float64
  Side            string
  Source          string
}

func (h Handicap) ComputeId() string {
  stringVal := fmt.Sprintf("%s%s%s%s%s", h.Side, h.HomeTeam, h.AwayTeam, h.Time.Format(time.RFC3339), h.TimeCollected.Format(time.RFC3339))
  hasher := sha1.New()
  hasher.Write([]byte(stringVal))
  return base64.URLEncoding.EncodeToString(hasher.Sum(nil))
}
