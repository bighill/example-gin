package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

type FormField struct {
	Apple string `form:"apple_field"`
}

type CheckboxField struct {
	Colors []string `form:"color_field"`
}

func GetFormField(c *gin.Context) {
	var f FormField
	c.Bind(&f)
	fmt.Printf("apple message: %s\n", f.Apple)
	c.JSON(200, gin.H{"apple_message": f.Apple})
}

func GetCheckboxField(c *gin.Context) {
	var checks CheckboxField
	c.ShouldBind(&checks)
	fmt.Printf("check boxes: %s\n", checks.Colors)
	c.JSON(200, gin.H{"check_boxes": checks.Colors})
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

	// curl "localhost:8080/form-colors?color_field=red&color_field=blue"
	r.GET("/form-colors", GetCheckboxField)

	r.Run()
}
