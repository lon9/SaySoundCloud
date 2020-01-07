package main

import (
	"flag"

	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/lon9/soundboard/backend/config"
	"github.com/lon9/soundboard/backend/database"
	"github.com/lon9/soundboard/backend/models"
	"github.com/lon9/soundboard/backend/server"
)

func main() {

	env := flag.String("e", "development", "")
	flag.Parse()

	config.Init(*env)
	database.Init(false, &models.User{}, &models.Application{})
	defer database.Close()
	if err := server.Init(); err != nil {
		panic(err)
	}
}
