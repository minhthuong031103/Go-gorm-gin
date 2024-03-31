package main

import (
	"fmt"
	"go-gorm-gin/config"
	"go-gorm-gin/middleware"
	"log"

	_ "go-gorm-gin/docs"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title           Cake ATS API
// @description     This is a sample server Cake ATS API server.
// @version         1.0

// @contact.name   Minh Thuong

// @host localhost:8080
// @BasePath  /

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
func main() {

	cfg, err := config.LoadConfig()
	if err != nil {
		panic(err)
	}
	fmt.Println("Connecting to database...")
	db, err := config.ConnectDatabaseWithRetryIn60s(cfg)
	if err != nil {
		log.Fatal("Error when connecting to database", err)
	}
	if cfg.Env == "dev" {
		db.Debug()
	}

	// appCtx := appctx.NewAppContext(db, cfg.SecretKey, cfg.ServerHost)
	r := gin.Default()
	r.Use(middleware.CORSMiddleware())
	r.Use(middleware.Recover())
	apiRoute := r.Group("/")
	{
		apiRoute.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
		log.Println("Swagger is running on http://localhost:8080/docs/index.html")
	}

	errorInit := r.Run(fmt.Sprintf(":%s", cfg.Port))
	if errorInit != nil {
		log.Fatalln("Error running server:", errorInit)

	}

}
