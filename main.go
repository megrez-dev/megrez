package main

import (
	"github.com/megrez/app"
	_ "github.com/megrez/docs"
	"log"
)

// @title           Megrez backend API
// @version         1.0
// @description     This is a megrez backend server.
// @termsOfService  http://megrez.run

// @contact.name   Megrez
// @contact.url    http://megrez.run
// @contact.email  alkaidchen@qq.com

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:8080
// @BasePath  /

func main() {
	megrez := app.New()
	if err := megrez.Init(); err != nil {
		log.Fatal("application init failed, ", err)
	}
	if err := megrez.Run(); err != nil {
		log.Fatal("application run failed, ", err)
	}
}
