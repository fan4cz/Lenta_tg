package handler

import "github.com/gin-gonic/gin"

type ExampleHandler struct {
}

func NewExampleHandler() *ExampleHandler  {
	return &ExampleHandler{}
}

func (h *ExampleHandler) GetExample(c *gin.Context) {
	c.JSON(200, gin.H{
        "message": "Hello from Gin!",
    })
}