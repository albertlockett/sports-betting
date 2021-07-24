package model

import "time"

type Event struct {
	HomeTeam string
	AwayTeam string
	Time time.Time
}