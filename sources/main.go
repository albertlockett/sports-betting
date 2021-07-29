package main

import (
	"fmt"
	"github.com/albertlockett/sports-betting/sources/betdsi"
	"github.com/albertlockett/sports-betting/sources/config"
	"github.com/albertlockett/sports-betting/sources/dao"
	"github.com/albertlockett/sports-betting/sources/fivethirtyeight"
	"log"
	"os"
)

func main() {
	config.Init()
	err := dao.Init()
	if err != nil {
		panic(err)
	}

	command := os.Args[1]

	switch command {
	case "expected-values":
		fetchExpectedValues()
	case "handicaps":
		fetchHandicaps()
	case "lines":
		fetchLines()
	}
}


func fetchHandicaps() {
	err := dao.ResetHandicapLatestCollected()
	if err != nil {
		panic(err)
	}
	handicaps, err := fivethirtyeight.FetchEvents()
	if err != nil {
		panic(err)
	}
	log.Println(fmt.Sprintf("There were %d events", len(handicaps) / 2))
	for _, handicap := range handicaps {
		dao.SaveHandicap(handicap)
	}
}

func fetchLines() {
	lines, err := betdsi.FetchLines()
	if err != nil {
		log.Println("an error happened")
		panic(err)
	}
	log.Println(fmt.Sprintf("There were %d lines", len(lines) / 2))

	// save lines
	err = dao.ResetLineLatestCollected()
	if err != nil {
		panic(err)
	}
	for _, line := range lines {
		dao.SaveLine(line)
	}
}

func fetchExpectedValues() {
	results, err := dao.FetchEvents()
	if err != nil {
		panic(err)
	}

	err = dao.ResetExpectedValueLatestCollected()
	if err != nil {
		panic(err)
	}

	for _, result := range results {
		dao.SaveExpectedValue(result)
	}
}