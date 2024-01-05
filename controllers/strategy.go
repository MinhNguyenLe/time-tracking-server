package controllers

import (
	"fmt"
	"net/http"

	"github.com/MinhNguyenLe/time-tracking-server/forms"
	"github.com/MinhNguyenLe/time-tracking-server/models"
	"github.com/gin-gonic/gin"
)

type StrategyController struct{}

var StrategyModel = new(models.StrategyModel)

func (sc *StrategyController) Insert(c *gin.Context) {
	var form forms.InsertStrategyForm

	if validationError := c.ShouldBindJSON(&form); validationError != nil {
		fmt.Println(validationError)
		c.AbortWithStatusJSON(http.StatusNotAcceptable, gin.H{"Message": "error"})
		return
	}

	msg, error := StrategyModel.Insert(form)

	if error != nil {
		fmt.Println(error)

		c.AbortWithStatusJSON(http.StatusNotAcceptable, gin.H{"message": "Error: Insert strategy"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": msg})
}

func (sc *StrategyController) GetList(c *gin.Context) {
	msg, error := StrategyModel.GetList()

	if error != nil {
		fmt.Println(error)

		c.AbortWithStatusJSON(http.StatusNotAcceptable, gin.H{"message": "Error: Get list strategies"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"strategies": msg})
}

func (sc *StrategyController) ChangeStatus(c *gin.Context) {
	var form forms.ChangeStrategyStatusForm
	if validationError := c.ShouldBindJSON(&form); validationError != nil {
		fmt.Println(validationError)
		c.AbortWithStatusJSON(http.StatusNotAcceptable, gin.H{"Message": "error"})
		return
	}

	error := StrategyModel.ChangeStatus(form)

	if error != nil {
		fmt.Println(error)

		c.AbortWithStatusJSON(http.StatusNotAcceptable, gin.H{"message": "Error: Change status strategy"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Successful"})
}

func (sc *StrategyController) TriggerCompleted(c *gin.Context) {
	var form forms.StrategyIdForm
	if validationError := c.ShouldBindJSON(&form); validationError != nil {
		fmt.Println(validationError)
		c.AbortWithStatusJSON(http.StatusNotAcceptable, gin.H{"Message": "error"})
		return
	}

	error := StrategyModel.TriggerCompleted(form)

	if error != nil {
		fmt.Println(error)

		c.AbortWithStatusJSON(http.StatusNotAcceptable, gin.H{"message": "Error: Strategy complete"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Successful"})
}

func (sc *StrategyController) GetStrategiesByStatus(c *gin.Context) {
	var form forms.StrategyStatusForm

	if validationError := c.ShouldBindJSON(&form); validationError != nil {
		fmt.Println(validationError)
		c.AbortWithStatusJSON(http.StatusNotAcceptable, gin.H{"Message": "error"})
		return
	}

	rows, error := StrategyModel.GetStrategiesByStatus(form)

	if error != nil {
		fmt.Println(error)

		c.AbortWithStatusJSON(http.StatusNotAcceptable, gin.H{"message": "Error: Get strategy by status"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": rows})
}
