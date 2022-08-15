package main

import (
	"context"
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
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

/*
mongo
*/
type Author struct {
	ID   primitive.ObjectID `bson:"_id"`
	Name string             `json:"name"`
}

func MongoCreateAuthor(c *gin.Context) {
	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()

	var author Author
	authorCollection := DB.Database("library").Collection("authors")

	err := c.Bind(&author)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	author.ID = primitive.NewObjectID()

	res, err := authorCollection.InsertOne(ctx, author)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, res)
}

func MongoGetAllAuthors(c *gin.Context) {
	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()

	var authors []bson.M
	authorCollection := DB.Database("library").Collection("authors")

	cursor, err := authorCollection.Find(ctx, bson.M{})
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	err = cursor.All(ctx, &authors)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, authors)
}

func MongoGetAuthor(c *gin.Context) {
	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()

	ID := c.Params.ByName("id")
	authorID, _ := primitive.ObjectIDFromHex(ID)

	var author bson.M
	authorCollection := DB.Database("library").Collection("authors")

	err := authorCollection.FindOne(ctx, bson.M{"_id": authorID}).Decode(&author)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, author)
}
