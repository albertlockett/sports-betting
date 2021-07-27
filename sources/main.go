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

	_, err = betdsi.FetchLines()
	if err != nil {
		log.Println("an error happened")
		panic(err)
	}

	handicaps, err := fivethirtyeight.FetchEvents()

	if err != nil {
		log.Println("an error happened")
		panic(err)
	}

	log.Println(fmt.Sprintf("There were %d events", len(handicaps) / 2))
	log.Println("Saving events")

	for _, handicap := range handicaps {
		dao.SaveHandicap(handicap)
	}
}
