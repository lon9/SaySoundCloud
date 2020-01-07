package server

import "github.com/lon9/soundboard/backend/config"

// Init initialize server
func Init() error {
	c := config.GetConfig()
	r, err := NewRouter()
	if err != nil {
		return err
	}
	r.Logger.Fatal(r.Start(":" + c.GetString("server.port")))
	return nil
}
