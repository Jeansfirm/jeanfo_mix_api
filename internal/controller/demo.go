package controller

import (
	"jeanfo_mix/internal/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type DemoController struct {
	Service *service.DemoService
}

func (c *DemoController) HelloWorld(ctx *gin.Context) {
	message := c.Service.GetHelloWord()
	ctx.JSON(http.StatusOK, gin.H{"message": message})
}

func (c *DemoController) GetDemos(ctx *gin.Context) {
	title := ctx.Query("tigle")
	page, _ := strconv.Atoi(ctx.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(ctx.DefaultQuery("page_size", "10"))

	demos, err := c.Service.GetDemos(title, page, pageSize)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	ctx.JSON(http.StatusOK, gin.H{"data": demos})
}
