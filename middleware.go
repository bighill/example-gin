package main

import (
	"log"
	"time"

	"github.com/gin-gonic/gin"
)

func SpecialMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		t := time.Now()

		c.Set("specialVariable", "foo")

		// before request
		c.Next()
		// after request

		latency := time.Since(t)
		log.Printf("special middleware latency was %s", latency)
	}
}
