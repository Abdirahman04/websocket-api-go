package main

import "github.com/gin-gonic/gin"

func main() {
  r := gin.Default()

  r.GET("/", func(c *gin.Context) {
    c.String(200, "Hello")
  })

  r.Run(":9000")
}
