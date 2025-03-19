package main

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	logging "github.com/namcnab/calculator/internal/logging"
	metrics "github.com/namcnab/calculator/internal/metrics"
	"github.com/namcnab/calculator/mathfuncs"
)

func main() {
	// Initialize the logger
	logging.InitLogger()

	// Create a new Gin router
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
	responseLog := CreateResponseLogRecord(c.ClientIP(), c.RemoteIP(), c.Request.Method, "AddTwo", time.Now().Format(time.RFC3339))

	// Unmarshal JSON into a map to validate fields
	data, err := validateAllowedKeys(c, map[string]bool{"a": true, "b": true}, responseLog)
	if err != nil {
		return
	}

	// Bind JSON to the struct
	if err := json.Unmarshal(data, &input); err != nil {
		handleError(c, 400, "AddTwoHandler: Invalid input", err.Error(), responseLog)
		return
	}

	// Perform the calculation
	startTime := time.Now().UnixNano()
	result := mathfuncs.AddTwo(input.A, input.B)
	endTime := time.Now().UnixNano()

	responseTime := metrics.CalculateResponseTime(startTime, endTime) // Response time in milliseconds

	SendLog("info", 200, responseTime, "", responseLog)
	c.JSON(200, gin.H{"result": result})
}

func AddListHandler(c *gin.Context) {
	var input mathfuncs.ListOfNums
	responseLog := CreateResponseLogRecord(c.ClientIP(), c.RemoteIP(), c.Request.Method, "AddList", time.Now().Format(time.RFC3339))

	if err := c.ShouldBindJSON(&input); err != nil {
		SendLog("error", 400, 0, err.Error(), responseLog)
		c.JSON(400, gin.H{"error": "Invalid input"})
		return
	}

	startTime := time.Now().UnixNano()
	result := mathfuncs.AddList(input.List) // Performs calculation
	endTime := time.Now().UnixNano()

	responseTime := metrics.CalculateResponseTime(startTime, endTime) // Response time in milliseconds

	SendLog("info", 200, responseTime, "", responseLog)
	c.JSON(200, gin.H{"result": result})
}

func DivideTwoHandler(c *gin.Context) {
	var input mathfuncs.TwoNums
	responseLog := CreateResponseLogRecord(c.ClientIP(), c.RemoteIP(), c.Request.Method, "DivideTwo", time.Now().Format(time.RFC3339))

	// Unmarshal JSON into a map to validate fields
	data, err := validateAllowedKeys(c, map[string]bool{"a": true, "b": true}, responseLog)
	if err != nil {
		return
	}

	// Bind JSON to the struct
	if err := json.Unmarshal(data, &input); err != nil {
		handleError(c, 400, "AddTwoHandler: Invalid input", err.Error(), responseLog)
		return
	}

	startTime := time.Now().UnixNano()
	result := mathfuncs.DivideTwo(input.A, input.B)
	endTime := time.Now().UnixNano()

	responseTime := metrics.CalculateResponseTime(startTime, endTime) // Response time in milliseconds

	SendLog("info", 200, responseTime, "", responseLog)
	c.JSON(200, gin.H{"result": result})
}

func MultiplyTwoHandler(c *gin.Context) {
	var input mathfuncs.TwoNums

	responseLog := CreateResponseLogRecord(c.ClientIP(), c.RemoteIP(), c.Request.Method, "MultiplyTwo", time.Now().Format(time.RFC3339))

	// Unmarshal JSON into a map to validate fields
	data, err := validateAllowedKeys(c, map[string]bool{"a": true, "b": true}, responseLog)
	if err != nil {
		return
	}

	// Bind JSON to the struct
	if err := json.Unmarshal(data, &input); err != nil {
		handleError(c, 400, "AddTwoHandler: Invalid input", err.Error(), responseLog)
		return
	}

	startTime := time.Now().UnixNano()
	result := mathfuncs.MultiplyTwo(input.A, input.B)
	endTime := time.Now().UnixNano()

	responseTime := metrics.CalculateResponseTime(startTime, endTime) // Response time in milliseconds

	SendLog("info", 200, responseTime, "", responseLog)
	c.JSON(200, gin.H{"result": result})
}

func SubtractTwoHandler(c *gin.Context) {
	var input mathfuncs.TwoNums
	responseLog := CreateResponseLogRecord(c.ClientIP(), c.RemoteIP(), c.Request.Method, "SubtractTwo", time.Now().Format(time.RFC3339))

	// Unmarshal JSON into a map to validate fields
	data, err := validateAllowedKeys(c, map[string]bool{"a": true, "b": true}, responseLog)
	if err != nil {
		return
	}

	// Bind JSON to the struct
	if err := json.Unmarshal(data, &input); err != nil {
		handleError(c, 400, "AddTwoHandler: Invalid input", err.Error(), responseLog)
		return
	}

	startTime := time.Now().UnixNano()
	result := mathfuncs.SubtractTwo(input.A, input.B)
	endTime := time.Now().UnixNano()

	responseTime := metrics.CalculateResponseTime(startTime, endTime) // Response time in milliseconds

	SendLog("info", 200, responseTime, "", responseLog)
	c.JSON(200, gin.H{"result": result})
}

func CreateResponseLogRecord(clientIp string, serverIP string, httpMethod string, opeartion string, timestamp string) logging.ResponseLogRecord {
	return logging.ResponseLogRecord{
		ClientIP:      clientIp,
		ServerIP:      serverIP,
		HttpMethod:    httpMethod,
		OperationName: opeartion,
		TimeStamp:     timestamp,
		TimeZone:      "GMT", // time.RFC3339
	}
}

func SendLog(level string, status int, responseTime float64, sysErrMsg string, logRecord logging.ResponseLogRecord) {
	if level == "info" {
		logRecord.Level = "INFO"
		logRecord.Status = 200
		logRecord.Message = "Operation completed successfully"
		logRecord.SysErrMsg = ""
		logRecord.ResponseTime = responseTime
		logging.LogResponse(logRecord)
	} else if level == "error" && status == 400 {
		logRecord.Level = "ERROR"
		logRecord.Status = 400
		logRecord.Message = "Invalid input"
		logRecord.SysErrMsg = sysErrMsg
		logRecord.ResponseTime = responseTime
		logging.LogResponse(logRecord)
	}
}

func validateAllowedKeys(c *gin.Context, allowedKeys map[string]bool, responseLog logging.ResponseLogRecord) ([]byte, error) {
	var rawData map[string]interface{}

	if err := c.ShouldBindJSON(&rawData); err != nil {
		handleError(c, 400, "validateAllowedKeys: Invalid input", err.Error(), responseLog)
		return nil, err
	}
	fmt.Println("Raw Data:", rawData)
	for key := range rawData {
		if !allowedKeys[key] {
			handleError(c, 400, "validateAllowedKeys: Unexpected field: "+key, "", responseLog)
			return nil, fmt.Errorf("unexpected field: %s", key)
		}
	}
	// convert rawData to JSON
	jsonData, err := json.Marshal(rawData)
	if err != nil {
		return nil, err
	}
	return jsonData, nil
}

func handleError(c *gin.Context, status int, message string, sysErrMsg string, responseLog logging.ResponseLogRecord) {
	SendLog("error", status, 0, sysErrMsg, responseLog)
	c.JSON(status, gin.H{"error": message})
}
