package main

import (
	"net/http"
	"net"
	"log"
	"fmt"

  "github.com/gin-gonic/contrib/static"
  "github.com/gin-gonic/gin"
)

func main() {
  log.Print("App Started...")

  router := gin.Default()

  router.Use(static.Serve("/", static.LocalFile("./public", true)))

  api := router.Group("/api")
  {
      api.GET("/whoami", func(c *gin.Context) {
				var ip = getIPAdress(c.Request);
				c.JSON(200, gin.H{
						"message": "Your IP is " + ip,
				})
		})
	}
	
  router.Run(":3000")
}

func getIPAdress(req *http.Request) string {

	ip, port, err := net.SplitHostPort(req.RemoteAddr)

	log.Print(fmt.Sprintf("%d-%d-%d", ip, port, err))
	if err != nil {
			return "Unknown"
	}

	userIP := net.ParseIP(ip)
	if userIP == nil {
		return "Unknown"
	}

	if userIP.IsLoopback() {
		return "localhost"
	}

	return userIP.String();
}
