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

func (c *DemoController) CreateDemo(ctx *gin.Context) {
	var req struct {
		Title   string `json:"title" binding:"required"`
		Content string `json:"content" binding:"required"`
	}

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	demo, err := c.Service.CreateDemo(req.Title, req.Content)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"data": demo})
}

func (c *DemoController) DeleteDemo(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid id"})
		return
	}

	if err := c.Service.DeleteDemo(uint(id)); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"msg": "Delete successfully"})
}
