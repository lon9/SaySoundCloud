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

	version.GET("/users/:id", userController.Show)
	version.POST("/users", userController.Create, authMiddleware.Verify)
	version.PUT("/users/:id", userController.Update, authMiddleware.Verify)
	version.DELETE("/users/:id", userController.Destroy, authMiddleware.Verify)

	applicationController := controllers.NewApplicationController()

	version.GET("/apps", applicationController.Index)
	version.GET("/apps/:id", applicationController.Show)
	version.POST("/apps", applicationController.Create, authMiddleware.Verify)
	version.PUT("/apps/:id", applicationController.Update, authMiddleware.Verify)
	version.PUT("/apps/:id/renewtoken", applicationController.RenewToken, authMiddleware.Verify)
	version.DELETE("/apps/:id", applicationController.Destroy, authMiddleware.Verify)

	version.POST("/apps/:id/ws", applicationController.WSAuth)
	version.GET("/apps/:id/ws", applicationController.WS)
	version.POST("/apps/:id/cmd", applicationController.Cmd)

	return router, nil
}
