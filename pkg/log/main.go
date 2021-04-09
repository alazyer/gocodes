package main

import (
	"fmt"
	"go.uber.org/zap/zapcore"
)

func main() {
	fmt.Printf("Debug: %+v, Info: %+v, Warn: %+v, Error: %+v", zapcore.DebugLevel, zapcore.InfoLevel, zapcore.WarnLevel, zapcore.ErrorLevel)
}
