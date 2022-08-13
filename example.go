package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

type FormField struct {
	Apple string `form:"apple_field"`
}

func GetFormField(c *gin.Context) {
	var f FormField
	c.Bind(&f)
	fmt.Printf("apple message: %s\n", f.Apple)
	c.JSON(200, gin.H{"apple_message": f.Apple})
}

func main() {
	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	// curl "localhost:8080/form?apple_field=Apples_are_for_eating"
	r.GET("/form", GetFormField)

	r.Run()
}
