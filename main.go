package main

import (
    "github.com/gin-gonic/gin"
    "gin_test/sampleapi"
)

var DB = make(map[string]string)

func main() {
    router := gin.Default()

    router.GET("/sampleapi", sampleapi.RssApi)

    router.Run(":8080")

}

