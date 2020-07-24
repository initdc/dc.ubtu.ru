package main

import (
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
)

func main() {
	all, err := ioutil.ReadFile("./sh/1key-docker-compose-ubuntu.sh")
	if err != nil {
		panic(err)
	}
	doc, err := ioutil.ReadFile("./sh/1key-docker-ubuntu.sh")
	if err != nil {
		panic(err)
	}
	com, err := ioutil.ReadFile("./sh/1key-compose-ubuntu.sh")
	if err != nil {
		panic(err)
	}

	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("$PORT must be set")
	}

	router := gin.New()
	router.Use(gin.Logger())
	router.LoadHTMLGlob("templates/index.html")

	router.StaticFile("/favicon.ico", "./favicon.png")

	router.GET("/", func(a *gin.Context) {
		ua := a.Request.Header.Get("User-Agent")
		if strings.Contains(ua, "Gecko") {
			a.HTML(http.StatusOK, "index.html", gin.H{
				"sh": string(all),
			})
		} else {
			a.File("./sh/1key-docker-compose-ubuntu.sh")
		}
	})

	router.GET("/d", func(d *gin.Context) {
		ua := d.Request.Header.Get("User-Agent")
		if strings.Contains(ua, "Gecko") {
			d.HTML(http.StatusOK, "index.html", gin.H{
				"sh": string(doc),
			})
		} else {
			d.File("./sh/1key-docker-ubuntu.sh")
		}
	})

	router.GET("/c", func(c *gin.Context) {
		ua := c.Request.Header.Get("User-Agent")
		if strings.Contains(ua, "Gecko") {
			c.HTML(http.StatusOK, "index.html", gin.H{
				"sh": string(com),
			})
		} else {
			c.File("./sh/1key-compose-ubuntu.sh")
		}
	})

	router.Run(":" + port)
}
