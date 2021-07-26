package main

import (
	"fmt"
	"github.com/albertlockett/sports-betting/sources/dao"
	"github.com/albertlockett/sports-betting/sources/fivethirtyeight"
	"log"
)

func main() {
	err := dao.Init()
	if err != nil {
		panic(err)
	}

	events, err := fivethirtyeight.FetchEvents()

	if err != nil {
		log.Println("an error happened")
		panic(err)
	}

	log.Println(fmt.Sprintf("There were %d events", len(events)))
	log.Println("Saving events")

	for _, event := range events {
		dao.SaveEvent(event)
	}
}
