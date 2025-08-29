package main

import "fmt"

// Adapter pattern

type OldLogger struct{} // adaptee

func (o *OldLogger) LegacyLog(message string) {
	fmt.Println("Legacy log:", message)
}

type NewLogger interface { // target
	Log(message string)
}

type LoggerAdapter struct { // adapter
	oldLogger *OldLogger
}

func (l *LoggerAdapter) Log(message string) {
	l.oldLogger.LegacyLog(message)
}

func main() {
	oldLogger := &OldLogger{}
	adapter := &LoggerAdapter{oldLogger: oldLogger}
	var logger NewLogger = adapter

	logger.Log("This is a new log message.")
}
