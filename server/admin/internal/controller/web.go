package controller

import (
	"insulation/server/admin/internal/auth"
	"insulation/server/admin/internal/routes"
	"insulation/server/base/pkg/config"
	"insulation/server/base/pkg/database"

	// _ "insulation/apis/admin"

	"github.com/gin-gonic/gin"
	// swaggerFiles "github.com/swaggo/files"
	// ginSwagger "github.com/swaggo/gin-swagger"
)

func Start(address string) {
	auth.InitializeAuth()
	if config.IsRelease() {
		gin.SetMode(gin.ReleaseMode)
	} else {
		gin.SetMode(gin.DebugMode)
	}
	database.Initialize()

	app := gin.Default()

	routes.NewLoginRoute(app)

	app.Run(address)
}
