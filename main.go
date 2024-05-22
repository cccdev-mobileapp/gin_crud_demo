package main

import (
	"fmt"
	"go-crud/config"
	"go-crud/middleware"
	"go-crud/routers"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	config.ConnectDB()
	config.Initialize()
	fmt.Println("after ConnectDB")
	// muxRouter := mux.NewRouter()
	ginRouter := gin.Default()

	ginRouter.LoadHTMLGlob("templates/*")
	ginRouter.Use(middleware.LoggerMiddleware())

	routers.RegisterUserRoutes(ginRouter)
	http.Handle("/", ginRouter)
	log.Fatal(http.ListenAndServe("localhost:3000", nil))
}
