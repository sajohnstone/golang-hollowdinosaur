package main

import (
  //"net/http"
  "log"

  "github.com/gin-gonic/contrib/static"
  "github.com/gin-gonic/gin"
)

func main() {
  log.Print("App Started...")

  router := gin.Default()

  router.Use(static.Serve("/", static.LocalFile("./public", true)))

  api := router.Group("/api")
  {
      api.GET("/events", func(c *gin.Context) {
          c.JSON(200, gin.H{
              "message": "pong",
          })
      })
  }

  router.Run(":3000")
}