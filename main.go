package main

import (
	"log"

	"github.com/megrez/app"

	"github.com/spf13/viper"
)

var config *viper.Viper

func main() {
	app := app.NewMegrez()
	if err := app.Init(); err != nil {
		log.Fatal("application init failed, ", err)
	}
	if err := app.Run(); err != nil {
		log.Fatal("application run failed, ", err)
	}

}
