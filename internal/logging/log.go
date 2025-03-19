package logging

import (
	"os"
	"os/exec"
	"strings"

	"github.com/sirupsen/logrus"
)

type ResponseLogRecord struct {
	ClientIP      string  `json:"client_ip"`
	ServerIP      string  `json:"server_ip"`
	Level         string  `json:"level"`
	Status        int     `json:"status"`
	HttpMethod    string  `json:"http_method"`
	Message       string  `json:"message"`
	SysErrMsg     string  `json:"sys_err_msg"`
	OperationName string  `json:"operation_name"`
	TimeStamp     string  `json:"time_stamp"`
	TimeZone      string  `json:"time_zone"`
	ResponseTime  float64 `json:"response_time"`
}

var logConfig = logrus.New()

func InitLogger() {
	// Set log format to JSON
	logConfig.SetFormatter(&logrus.JSONFormatter{})

	// Set log file path
	//localPath := getModuleRoot() + "/logs/app.log"
	dockerPath := "/app/logs/app.log"

	// Write logs to a file
	file, err := os.OpenFile(dockerPath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err == nil {
		logConfig.SetOutput(file)
	} else {
		logConfig.SetOutput(os.Stdout)
	}
}

func LogResponse(fields ResponseLogRecord) {
	entry := logConfig.WithFields(logrus.Fields{
		"client_ip":      fields.ClientIP,
		"server_ip":      fields.ServerIP,
		"status":         fields.Status,
		"http_method":    fields.HttpMethod,
		"sys_err_msg":    fields.SysErrMsg,
		"operation_name": fields.OperationName,
		"time_stamp":     fields.TimeStamp,
		"time_zone":      fields.TimeZone,
		"response_time":  fields.ResponseTime,
	})

	switch fields.Level {
	case "INFO":
		entry.Info(fields.Message)
	case "ERROR":
		entry.Error(fields.Message)
	default:
		entry.Debug(fields.Message)
	}

}

func getModuleRoot() string {
	cmd := exec.Command("go", "list", "-m", "-f", "{{.Dir}}")
	output, err := cmd.CombinedOutput()
	if err != nil {
		panic(err)
	}
	return strings.TrimSpace(string(output))
}
