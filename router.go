package main

import "github.com/gin-gonic/gin"

func initRouter() *gin.Engine {
	r := gin.Default()
	r.SetTrustedProxies(nil)
	r.HTMLRender = initRenderer()
	r.Use(SpecialMiddleware())

	r.GET("/", func(c *gin.Context) {
		specialMessage := c.MustGet("specialVariable").(string)
		c.JSON(200, gin.H{"message": "root route", "special_mesage": specialMessage})
	})

	r.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})

	// curl "localhost:8080/form?apple_field=Apples_are_for_eating"
	r.GET("/form", GetFormField)

	// curl "localhost:8080/form-colors?color_field=red&color_field=blue"
	r.GET("/form-colors", GetCheckboxField)

	// curl -X POST localhost:8080/post-data --data '{"star":"Sun","planet":"Venus"}' -H "Content-Type:application/json"
	r.POST("/post-data", PostData)

	// curl -v localhost:8080/dynamic-uri/lex/brokenuuid
	// curl localhost:8080/dynamic-uri/lex/987fbc97-4bed-5078-9f07-9141ba07c9f3
	r.GET("/dynamic-uri/:name/:id", GetDynamicUri)

	r.GET("/html", func(c *gin.Context) {
		c.HTML(200, "default", gin.H{"templateVar": "foo"})
	})
	r.GET("/html2", func(c *gin.Context) {
		c.HTML(200, "secondary", gin.H{"templateVar": "bar"})
	})

	r.POST("/mongo/authors", MongoCreateAuthor)
	r.GET("/mongo/authors", MongoGetAllAuthors)
	r.GET("/mongo/authors/:id", MongoGetAuthor)

	return r
}
