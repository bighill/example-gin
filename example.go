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

type JsonPostData struct {
	Star   string `form:"star"`
	Planet string `form:"planet"`
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
	fmt.Printf("check boxes: %s \n", checks.Colors)
	c.JSON(200, gin.H{"check_boxes": checks.Colors})
}

func PostData(c *gin.Context) {
	var data JsonPostData
	c.Bind(&data)
	fmt.Printf("post data: Star is %s, Planet is %s \n", data.Star, data.Planet)
	c.JSON(200, gin.H{"star": data.Star, "planet": data.Planet})
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

	// curl -X POST localhost:8080/post-data --data '{"star":"Sun","planet":"Venus"}' -H "Content-Type:application/json"
	r.POST("/post-data", PostData)

	r.Run()
}
