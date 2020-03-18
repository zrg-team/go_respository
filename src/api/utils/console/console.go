package console

import (
	"encoding/json"
	"log"

	"github.com/fatih/color"
)

// Pretty json
func Pretty(data interface{}) {
	b, err := json.MarshalIndent(data, "", "")
	if err != nil {
		log.Println(err)
		return
	}
	color.Magenta(string(b))
}

// Log normal data
func Log(message string, options ...interface{}) {
	color.Blue(message, options...)
}

// Highlight data
func Highlight(message string, options ...interface{}) {
	color.Cyan(message, options...)
}

// Warning log data
func Warning(message string, options ...interface{}) {
	color.Yellow(message, options...)
}

// Info log data
func Info(message string, options ...interface{}) {
	color.Green(message, options...)
}

// Error log data
func Error(message string, options ...interface{}) {
	color.Red(message, options...)
}
