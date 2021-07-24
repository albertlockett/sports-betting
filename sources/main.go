package main

import (
	"fmt"
	"github.com/albertlockett/sports-betting/sources/dao"
	"github.com/albertlockett/sports-betting/sources/fivethirtyeight"
)

func main() {
	fmt.Println("Hello, world.")
	err := dao.Init()
	if err != nil {
		panic(err)
	}

	events, err := fivethirtyeight.FetchEvents()

	if err != nil {
		fmt.Println("an error happened")
		panic(err)
	}

	fmt.Println(fmt.Sprintf("There were %d events", len(events)))
}
