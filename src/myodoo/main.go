package main

import (
    //"fmt"
    //"log"
    //"myodoo/cmd"
    "github.com/gin-gonic/gin"
    "myodoo/redis"
    "myodoo/routes"
)

func main() {

	redis.InitRedis()
	router := gin.Default()
	routes.SetupRoutes(router)
	router.Run(":8080")

    //fmt.Println("🚀 Starting MyOdoo Framework...")
    //err := cmd.RunServer()
    //if err != nil {
    //    log.Fatalf("Server failed to start: %v", err)
    }
