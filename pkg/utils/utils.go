package utils

import (
	"fmt"
	"time"
)

func PrintMessage(message string) {
	fmt.Println(message)
}

func GetLogFilename() string {
	currentTime := time.Now().Format("2006-01-02")
	return fmt.Sprintf("./logs/%s.log", currentTime)
}
