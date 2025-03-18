package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/namcnab/calculator/mathfuncs"
)

func main() {
	ginEngine := gin.Default()

	// Add CORS middleware
	ginEngine.Use(cors.Default())

	// Define routes
	ginEngine.POST("/addtwo", AddTwoHandler)
	ginEngine.POST("/addList", AddListHandler)
	ginEngine.POST("/dividetwo", DivideTwoHandler)
	ginEngine.POST("/multiplytwo", MultiplyTwoHandler)
	ginEngine.POST("/subtracttwo", SubtractTwoHandler)

	// Start the server on port 8080
	ginEngine.Run(":8080")
}

func AddTwoHandler(c *gin.Context) {
	var input mathfuncs.TwoNums
	if err := c.BindJSON(&input); err != nil {
		c.JSON(400, gin.H{"error": "Invalid input"})
		return
	}
	result := mathfuncs.AddTwo(input.A, input.B)
	c.JSON(200, gin.H{"result": result})
}

func AddListHandler(c *gin.Context) {
	var input mathfuncs.ListOfNums
	if err := c.BindJSON(&input); err != nil {
		c.JSON(400, gin.H{"error": "Invalid input"})
		return
	}
	result := mathfuncs.AddList(input.List)
	c.JSON(200, gin.H{"result": result})
}

func DivideTwoHandler(c *gin.Context) {
	var input mathfuncs.TwoNums
	if err := c.BindJSON(&input); err != nil {
		c.JSON(400, gin.H{"error": "Invalid input"})
		return
	}
	result := mathfuncs.DivideTwo(input.A, input.B)
	c.JSON(200, gin.H{"result": result})
}

func MultiplyTwoHandler(c *gin.Context) {
	var input mathfuncs.TwoNums
	if err := c.BindJSON(&input); err != nil {
		c.JSON(400, gin.H{"error": "Invalid input"})
		return
	}
	result := mathfuncs.MultiplyTwo(input.A, input.B)
	c.JSON(200, gin.H{"result": result})
}

func SubtractTwoHandler(c *gin.Context) {
	var input mathfuncs.TwoNums
	if err := c.BindJSON(&input); err != nil {
		c.JSON(400, gin.H{"error": "Invalid input"})
		return
	}
	result := mathfuncs.SubtractTwo(input.A, input.B)
	c.JSON(200, gin.H{"result": result})
}
