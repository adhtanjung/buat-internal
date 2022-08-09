package main

import (
	"log"
	"net/http"

	_ "github.com/adhtanjung/trackerstache_api/docs/trackerstache"
	"github.com/adhtanjung/trackerstache_api/pkg/config"
	"github.com/adhtanjung/trackerstache_api/pkg/db"
	"github.com/adhtanjung/trackerstache_api/pkg/modules/auth"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title           Buat Internal API
// @version         1.0
// @description     This is a sample server celler server.
// @termsOfService  http://swagger.io/terms/

// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io
// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:8080
// @BasePath  /api/v1
// @schemes http

// @securityDefinitions.basic  BasicAuth
func main() {
	c, err := config.LoadConfig()

	if err != nil {
		log.Fatalln("Failed at config", err)
	}
	// Initialize DB
	db.Init(c.DBUrl)

	route := gin.Default()
	v1 := route.Group("/api/v1")
	v1.GET("/", HealthCheck)

	// Initialize Auth
	authRepo := auth.NewRepository(&db.DB)
	authHandler := auth.NewAuthHandler(authRepo)
	authRoute := auth.NewAuthRoute(v1, authHandler)
	authRoute.Init(v1)
	url := ginSwagger.URL("http://localhost:8080/swagger/doc.json")
	route.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))

	route.Run(c.Port)

}

// HealthCheck godoc
// @Summary Show the status of server.
// @Description get the status of server.
// @Tags root
// @Accept */*
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router / [get]
func HealthCheck(c *gin.Context) {
	res := map[string]interface{}{
		"data": "Server is up and running",
	}

	c.JSON(http.StatusOK, res)
}
