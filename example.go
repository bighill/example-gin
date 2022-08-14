package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

/*
form
*/
type FormField struct {
	Apple string `form:"apple_field"`
}

func GetFormField(c *gin.Context) {
	var f FormField
	c.Bind(&f)
	fmt.Printf("apple message: %s\n", f.Apple)
	c.JSON(200, gin.H{"apple_message": f.Apple})
}

/*
checkbox
*/
type CheckboxField struct {
	Colors []string `form:"color_field"`
}

func GetCheckboxField(c *gin.Context) {
	var checks CheckboxField
	c.ShouldBind(&checks)
	fmt.Printf("check boxes: %s \n", checks.Colors)
	c.JSON(200, gin.H{"check_boxes": checks.Colors})
}

/*
post
*/
type JsonPostData struct {
	Star   string `form:"star"`
	Planet string `form:"planet"`
}

func PostData(c *gin.Context) {
	var data JsonPostData
	c.Bind(&data)
	c.JSON(200, gin.H{"star": data.Star, "planet": data.Planet})
	fmt.Printf("post data: Star is %s, Planet is %s \n", data.Star, data.Planet)
}

/*
dynamic URI
*/
type DynamicUri struct {
	ID   string `uri:"id" binding:"required,uuid"`
	Name string `uri:"name" binding:"required"`
}

func GetDynamicUri(c *gin.Context) {
	var uri DynamicUri
	err := c.ShouldBindUri(&uri)
	if err != nil {
		c.JSON(400, gin.H{"message": "invalid URI", "error": err})
		fmt.Println("invalid URI")
		return
	}
	c.JSON(200, gin.H{"name": uri.Name, "uuid": uri.ID})
	fmt.Printf("dynamic URI: Name is %s, UUID is %s \n", uri.Name, uri.ID)
}

func main() {
	r := initRouter()
	r.Run(":8080")
}
