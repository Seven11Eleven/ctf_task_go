package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/Seven11Eleven/ctf_task_go/internal/database"
)

func main() {

	fmt.Println("Hello, Go!")

	router := gin.Default()

	router.Static("/static", "./static")
	router.LoadHTMLGlob("templates/*")

	land := router.Group("land")
	{
		land.GET("/ru", func(c *gin.Context) {
			c.HTML(200, "land.html", nil)
		})

		land.GET("/fr", func(c *gin.Context) {
			c.HTML(200, "landfr.html", nil)
		})
	}

	router.Run(":1337")
}
