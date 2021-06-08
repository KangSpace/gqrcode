package logger

import (
	"fmt"
	"runtime"
)

// Define log handle here

func Warn(msg interface{}){
	fmt.Println("[WARN]", msg)
}

func Info(msg interface{}) {
	fmt.Println("[INFO]", msg)
}

func Error(msg interface{}){
	if _,ok:= msg.(error); ok {
		for i := 1; ; i++ {
			pc, file, line, ok := runtime.Caller(i)
			if !ok {
				break
			}
			f := runtime.FuncForPC(pc)
			if f.Name() != "runtime.main" && f.Name() != "runtime.goexit" {
				fmt.Printf("[ERROR] %s %s (%d) %s \n",file,f.Name(), line, msg)
			}
		}
	}else{
		fmt.Println("[ERROR] ", msg)
	}
}