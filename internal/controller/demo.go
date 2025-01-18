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

// @Summary Demo Say Hello
// @Description Demo api for getting hello message
// @Tags Demo
// @Success 200 {string} string "ok"
// @Router /api/demos/hello [get]
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

type CreateDemoReq struct {
	Title   string `json:"title" binding:"required"`
	Content string `json:"content" binding:"required"`
}

// @Summary Demo Create
// @Tags Demo
// @Success 200 {string} string "ok"
// @Param demo body CreateDemoReq true "create demo"
// @Router /api/demos [post]
func (c *DemoController) CreateDemo(ctx *gin.Context) {
	var req CreateDemoReq

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

// @Summary Demo Delete
// @Tags Demo
// @Success 200 {string} string "ok"
// @Param id path int true "DemoID"
// @Router /api/demos/{id} [delete]
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
