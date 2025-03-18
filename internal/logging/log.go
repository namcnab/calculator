package logging

import (
	"log"
	"os"
	"os/exec"
	"strings"
)

func SendLog(msg string, level string) {
	// Open the log file. Create it if it doesn't exist; append if it does.
	root := getModuleRoot()
	localPath := root + "/logs/app.log"
	//dockerPath := "/app/logs/app.log"
	logFile, err := os.OpenFile(localPath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer logFile.Close()

	// Set the output of the default logger to the file.
	log.SetOutput(logFile)

	// Log messages
	switch level {
	case "info":
		log.Println("INFO: " + msg)
	case "error":
		log.Println("ERROR: " + msg)
	case "debug":
		log.Println("DEBUG: " + msg)
	default:
		log.Println("UNKNOWN LEVEL: " + msg)
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
