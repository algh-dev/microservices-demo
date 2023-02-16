package main

import (
	"io"
	"log"
	"os"
	"time"

	"github.com/algh-dev/microservices-demo/middleware"
	"github.com/algh-dev/microservices-demo/validator"
	"github.com/gin-gonic/gin"
)

type Person struct {
	Name     string    `form:"name"`
	Address  string    `form:"address"`
	Birthday time.Time `form:"birthday" time_format:"2006-01-02" time_utc:"2"`
}

type UriParam struct {
	Id   string `uri:"id" binding:"required,uuid"`
	Name string `uri:"name" binding:"required"`
}

func main() {
	//write logs to a file
	f, _ := os.Create("gin.log")
	mw := io.MultiWriter(f, os.Stdout)
	gin.DefaultWriter = mw

	log.SetOutput(mw)

	route := gin.Default()

	//route := gin.New()
	route.Use(middleware.Logger())

	//register custom validators
	validator.RegisterBookableValidator()

	SetupRoutes(route)

	route.Run()
}
