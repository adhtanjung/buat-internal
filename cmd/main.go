package main

import (
	"log"

	"github.com/adhtanjung/trackerstache_api/pkg/config"
	"github.com/adhtanjung/trackerstache_api/pkg/db"
	"github.com/adhtanjung/trackerstache_api/pkg/modules/auth"
	"github.com/gin-gonic/gin"
)

func main() {
	c, err := config.LoadConfig()

	if err != nil {
		log.Fatalln("Failed at config", err)
	}
	// Initialize DB
	db.Init(c.DBUrl)

	route := gin.Default()
	v1 := route.Group("/api/v1")

	// Initialize Auth
	authRepo := auth.NewRepository(&db.DB)
	authHandler := auth.NewAuthHandler(authRepo)
	authRoute := auth.NewAuthRoute(v1, authHandler)
	authRoute.Init(v1)

	route.Run(c.Port)

}
