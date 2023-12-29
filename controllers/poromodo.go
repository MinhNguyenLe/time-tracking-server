package controllers

import (
	"fmt"
	"net/http"

	"github.com/MinhNguyenLe/time-tracking-server/forms"
	"github.com/MinhNguyenLe/time-tracking-server/models"
	"github.com/gin-gonic/gin"
)

type PoromodoController struct{}

var PoromodoModel = new(models.PoromodoModel)

func (pc *PoromodoController) Insert(c *gin.Context) {
	var form forms.InsertPoromodoForm

	fmt.Println("in")

	if validationError := c.ShouldBindJSON(&form); validationError != nil {
		fmt.Println("error")

		fmt.Println(validationError)
		c.AbortWithStatusJSON(http.StatusNotAcceptable, gin.H{"Message": "error"})
		return
	}
	fmt.Println("pass")

	msg, error := PoromodoModel.Insert(form)
	if error != nil {
		fmt.Println(error)

		c.AbortWithStatusJSON(http.StatusNotAcceptable, gin.H{"message": "Error: Poromodo"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": msg})
}
