package main

import (
	"fmt"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/namcnab/calculator/mathfuncs"
)

func main() {
	ginEngine := gin.Default()

	// Add CORS middleware
	ginEngine.Use(cors.Default())

	ginEngine.POST("/addtwo", AddHandler)
	ginEngine.Run(":8080")
	fmt.Println(mathfuncs.AddTwo(1, 2))
}

func AddHandler(c *gin.Context) {
	var input struct {
		A int `json:"a"`
		B int `json:"b"`
	}
	if err := c.BindJSON(&input); err != nil {
		c.JSON(400, gin.H{"error": "Invalid input"})
		return
	}
	result := mathfuncs.AddTwo(input.A, input.B)
	c.JSON(200, gin.H{"result": result})
}
