package binding

import (
	"log"
	"net/http"
	"time"

	"github.com/algh-dev/microservices-demo/validator"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
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

func BindPerson(c *gin.Context) {
	var person Person

	err := c.Bind(&person)

	if err != nil {
		log.Println("Failed to bind Person", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "Failed to bind Person",
		})

		return
	}

	log.Printf("Successfully bound Person, person=%v\n", person)

	c.JSON(http.StatusOK, gin.H{
		"name":     person.Name,
		"address":  person.Address,
		"birthday": person.Birthday,
	})
}

func BindUriParam(c *gin.Context) {
	var uriParam UriParam

	err := c.ShouldBindUri(&uriParam)

	if err != nil {
		log.Println("Failed to bing uri params", err)
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "message": err})
		return
	}

	c.JSON(http.StatusOK, gin.H{"name": uriParam.Name, "uuid": uriParam.Id})
}

func BindCustomValidator(c *gin.Context) {
	var b validator.Booking

	if err := c.ShouldBindWith(&b, binding.Query); err == nil {
		c.JSON(http.StatusOK, gin.H{"message": "Booking dates are valid!"})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
	}

}
