package main

import (
	"fmt"
	"os"
	"time"
)

const (
	logFileName = "app.log"
	maxLines    = 200000
)

func main() {
	file, err := os.OpenFile(logFileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Printf("Error opening log file: %v\n", err)
		return
	}
	defer file.Close()

	for i := 1; i <= maxLines; i++ {
		logEntry := fmt.Sprintf("%s [INFO] This is a sample log entry number %d.\n", time.Now().Format("2006-01-02 15:04:05"), i)
		_, err := file.WriteString(logEntry)
		if err != nil {
			fmt.Printf("Error writing to log file: %v\n", err)
			return
		}
		if i%10000 == 0 {
			fmt.Printf("Generated %d lines...\n", i)
		}
	}

	fmt.Printf("Log file '%s' with %d lines generated successfully.\n", logFileName, maxLines)
}


