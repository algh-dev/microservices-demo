package validator

import (
	"log"
	"time"

	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

type Booking struct {
	CheckIn  time.Time `form:"check_in" binding:"required,bookabledate" time_format:"2006-01-02"`
	CheckOut time.Time `form:"check_out" binding:"required,gtfield=CheckIn,bookabledate" time_format:"2006-01-02"`
}

var bookableDate validator.Func = func(fl validator.FieldLevel) bool {
	date, ok := fl.Field().Interface().(time.Time)

	if ok {
		today := time.Now()
		if today.After(date) {
			return false
		}
	}

	return true
}

func RegisterBookableValidator() {
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		err := v.RegisterValidation("bookabledate", bookableDate)
		if err != nil {
			log.Println("Failed to register bookabledate validator", err.Error())
		}
	}
}

func GetBookableDate() validator.Func {
	return bookableDate
}
