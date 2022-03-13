package main

import (
	"log"

	"github.com/megrez/app"
)

func main() {
	megrez := app.New()
	if err := megrez.Init(); err != nil {
		log.Fatal("application init failed, ", err)
	}
	if err := megrez.Run(); err != nil {
		log.Fatal("application run failed, ", err)
	}
}
