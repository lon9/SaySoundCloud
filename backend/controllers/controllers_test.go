package controllers

import (
	"database/sql"
	"os"
	"testing"

	"github.com/lon9/SaySoundCloud/backend/models"

	"github.com/go-testfixtures/testfixtures/v3"
	"github.com/lon9/SaySoundCloud/backend/config"
	"github.com/lon9/SaySoundCloud/backend/database"
)

var fixtures *testfixtures.Loader

func TestMain(m *testing.M) {
	config.Init("unit_test")
	database.Init(true, &models.User{}, &models.Application{})
	db, err := sql.Open("postgres", config.GetConfig().GetString("db.url"))
	if err != nil {
		panic(err)
	}

	fixtures, err = testfixtures.New(
		testfixtures.Database(db),
		testfixtures.Dialect("postgres"),
		testfixtures.Directory("fixtures"),
	)
	if err != nil {
		panic(err)
	}

	os.Exit(m.Run())
}

func prepareTestDatabase() {
	if err := fixtures.Load(); err != nil {
		panic(err)
	}
}
