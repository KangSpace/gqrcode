package logger

import (
	"fmt"
)

// Define log handle here

func Warn(msg interface{}){
	fmt.Println("[WARN]", msg)
}

func Info(msg interface{}) {
	fmt.Println("[INFO]", msg)
}

func Error(msg interface{}){
	fmt.Println("[ERROR] ", msg)
}