package main

import "github.com/gin-gonic/gin"


func main() {
	r := gin.Default()
	r.LoadHTMLGlob("assets/*.html")
	
	r.GET("/", func(c *gin.Context) {
		c.HTML(200, "index.html", gin.H{})
	})

	r.Run(":3000")
}