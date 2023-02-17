package main

import (
	"log"
	"net/http"

	"github.com/algh-dev/microservices-demo/cache"
	"github.com/algh-dev/microservices-demo/cache/repository"
	"github.com/gin-gonic/gin"
)


func SetupMoviesRoutes(route *gin.Engine) {
	redisClient, ctx := cache.GetRedisClient()

	repository := repository.NewMovieRepository(redisClient, ctx)


	route.POST("/movies", func(ctx *gin.Context) {
		var movie cache.Movie

		if err := ctx.ShouldBind(&movie); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error" : err.Error()})
			log.Println("Failed to bind request body", err.Error())
			return
		}

		res, err := repository.CreateMovie(&movie)

		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error" : err.Error()})
			log.Println("Failed to save movie ", err.Error())

			return
		}

		ctx.JSON(http.StatusOK, gin.H{
			"id" : res.Id,
			"title" : res.Title,
			"description" : res.Description,
		})
	})

	route.GET("/movies/:id", func(ctx *gin.Context) {
		id := ctx.Param("id")

		val, err := repository.GetMovie(id)
		if err != nil {
			ctx.JSON(http.StatusNotFound, gin.H{"error" : err.Error()})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{
			"id" : val.Id,
			"title" : val.Title,
			"description" : val.Description,
		})
	})
}