package main

import (
	"flag"

	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/lon9/SaySoundCloud/backend/config"
	"github.com/lon9/SaySoundCloud/backend/database"
	"github.com/lon9/SaySoundCloud/backend/models"
	"github.com/lon9/SaySoundCloud/backend/server"
)

func main() {

	env := flag.String("e", "development", "")
	flag.Parse()

	config.Init(*env)
	database.Init(false, &models.User{}, &models.Application{}, &models.Sounds{})
	defer database.Close()
	if err := server.Init(); err != nil {
		panic(err)
	}
}
