package main

import (
	"fmt"
	"github.com/albertlockett/sports-betting/sources/betdsi"
	"github.com/albertlockett/sports-betting/sources/dao"
	"github.com/albertlockett/sports-betting/sources/fivethirtyeight"
	"log"
)

func main() {
	err := dao.Init()
	if err != nil {
		panic(err)
	}

	// fetch lines from book
	lines, err := betdsi.FetchLines()
	if err != nil {
		log.Println("an error happened")
		panic(err)
	}
	log.Println(fmt.Sprintf("There were %d lines", len(lines) / 2))

	// save lines
	err = dao.ResetLineLatestCollected()
	//if err != nil {
	//	panic(err)
	//}
	for _, line := range lines {
		dao.SaveLine(line)
	}

	// fetch handicap daa
	handicaps, err := fivethirtyeight.FetchEvents()
	if err != nil {
		log.Println("an error happened")
		panic(err)
	}
	log.Println(fmt.Sprintf("There were %d events", len(handicaps) / 2))

	// save handicap data
	err = dao.ResetHandicapLatestCollected()
	//if err != nil {
	//	panic(err)
	//}
	for _, handicap := range handicaps {
		dao.SaveHandicap(handicap)
	}
}
