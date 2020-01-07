package server

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/lon9/soundboard/backend/config"
	"github.com/lon9/soundboard/backend/controllers"
	mymiddleware "github.com/lon9/soundboard/backend/middleware"
)

// NewRouter is constructor for router
func NewRouter() (*echo.Echo, error) {
	c := config.GetConfig()
	router := echo.New()
	router.Use(middleware.Logger())
	router.Use(middleware.Recover())
	authMiddleware, err := mymiddleware.NewFireBaseAuthMiddleware(c.GetString("server.firebase_config"), nil)
	if err != nil {
		return nil, err
	}

	healthController := controllers.NewHealthController()

	router.GET("/health", healthController.Index)
	router.GET("/secret", healthController.Secret, authMiddleware.Verify)

	version := router.Group("/v1")

	userController := controllers.NewUserController()

	version.GET("/user/:uid", userController.Show)
	version.PUT("/user/:uid", userController.Update, authMiddleware.Verify)
	version.DELETE("/user/:uid", userController.Destroy, authMiddleware.Verify)

	return router, nil
}
