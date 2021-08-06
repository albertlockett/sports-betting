package model

import (
  "crypto/sha1"
  "encoding/base64"
  "fmt"
  "time"
)

const SIDE_HOME = "home"
const SIDE_AWAY = "away"

const LINE_TYPE_MONEYLINE = "moneyline"

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

func (h Handicap) ComputeEventId() string {
  stringVal := fmt.Sprintf("%s%s%s%s", h.Side, h.HomeTeam, h.AwayTeam, h.Time.Format(time.RFC3339))
  hasher := sha1.New()
  hasher.Write([]byte(stringVal))
  return base64.URLEncoding.EncodeToString(hasher.Sum(nil))
}

func (h Handicap) ComputeId() string {
  stringVal := fmt.Sprintf("%s%s", h.ComputeEventId(), h.TimeCollected.Format(time.RFC3339))
  hasher := sha1.New()
  hasher.Write([]byte(stringVal))
  return base64.URLEncoding.EncodeToString(hasher.Sum(nil))
}

type Line struct {
  Event
  TimeCollected   time.Time
  LatestCollected bool
  LineAmerican    int32
  LineDecimal     float64
  Side            string
  Source          string
  Type            string
}

func (l Line) ComputeEventId() string {
  stringVal := fmt.Sprintf("%s%s%s%s", l.Side, l.HomeTeam, l.AwayTeam, l.Time.Format(time.RFC3339))
  hasher := sha1.New()
  hasher.Write([]byte(stringVal))
  return base64.URLEncoding.EncodeToString(hasher.Sum(nil))
}

func (l Line) ComputeId() string {
  stringVal := fmt.Sprintf("%s%s", l.ComputeEventId(), l.TimeCollected.Format(time.RFC3339))
  hasher := sha1.New()
  hasher.Write([]byte(stringVal))
  return base64.URLEncoding.EncodeToString(hasher.Sum(nil))
}

type ExpectedValue struct {
  Event
  Handicap
  Line
  LatestCollected bool
  Side            string
  ExpectedValue   float64
}

func (e ExpectedValue) ComputeEventId() string {
  stringVal := fmt.Sprintf("%s%s%s%s%s", e.Line.Side, e.HomeTeam, e.AwayTeam, e.Time.Format(time.RFC3339), e.Handicap.TimeCollected.Format(time.RFC3339))
  hasher := sha1.New()
  hasher.Write([]byte(stringVal))
  return base64.URLEncoding.EncodeToString(hasher.Sum(nil))
}

func (e ExpectedValue) ComputeId() string {
  stringVal := fmt.Sprintf("%s%s", e.Line.Side, e.HomeTeam, e.AwayTeam, e.Time.Format(time.RFC3339), e.Handicap.TimeCollected.Format(time.RFC3339), e.Line.TimeCollected.Format(time.RFC3339))
  hasher := sha1.New()
  hasher.Write([]byte(stringVal))
  return base64.URLEncoding.EncodeToString(hasher.Sum(nil))
}
