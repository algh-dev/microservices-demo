package main

import (
	"log"
	"net/http"
	"time"

	"github.com/algh-dev/microservices-demo/binding"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(route *gin.Engine) {

	route.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "pong",
		})
	})

	//bind query params
	route.GET("/bind", binding.BindPerson)

	//bind URI path
	route.GET("/:id/:name", binding.BindUriParam)

	//test custom middleware
	route.GET("/middleware", func(ctx *gin.Context) {
		example := ctx.MustGet("example")

		log.Println("Middleware variable: ", example)
	})

	//custom validator
	route.GET("/bookable", binding.BindCustomValidator)

	//go routines
	route.GET("/long_async", func(ctx *gin.Context) {
		cCtx := ctx.Copy()

		go func() {
			time.Sleep(5 * time.Second)

			log.Println("Done! in path " + cCtx.Request.URL.Path)
		}()
	})

	route.GET("/long_sync", func(ctx *gin.Context) {
		time.Sleep(5 * time.Second)

		log.Println("Done! in path " + ctx.Request.URL.Path)
	})

	//jsonp
	route.GET("/jsonp", func(ctx *gin.Context) {
		ctx.JSONP(http.StatusOK, gin.H{"foo": "bar"})
	})

	//movies endpoints
	SetupMoviesRoutes(route)
}
