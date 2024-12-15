package logging

import (
	"log"
	"os"
)

var logger *log.Logger

func InitLogger() {
	file, err := os.OpenFile("build/app.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatalf("Error opening log file: %v", err)
	}

	logger = log.New(file, "", log.LstdFlags)
}

func Info(msg string) {
	logger.Println("INFO: " + msg)
}

func Error(msg string) {
	logger.Println("ERROR: " + msg)
}
