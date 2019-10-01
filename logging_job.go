package main

import (
	"fmt"
	"os"
	"strings"
	"sync"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/diode"
)

//ErrorLog is used to map error log
//EventLog is used to map event log
//MessageLog is used to map message log
var (
	Log ILog
)

// ILog is used to create interface logging
type ILog interface {
	Error(text ...string)
	Event(text ...string)
	Message(text ...string)
	Fatal(wg *sync.WaitGroup, text ...string)
}

//Logger is used to create object logging
type Logger struct {
	ErrorLog   zerolog.Logger
	EventLog   zerolog.Logger
	MessageLog zerolog.Logger
}

// Error is use to write error log
func (logger *Logger) Error(text ...string) {
	log := strings.Join(text, "][")
	timenow := "[" + time.Now().Format("2006-01-02 15:04:05") + "]"
	logger.ErrorLog.Error().Msg(timenow + log)
}

// Event is use to write event log
func (logger *Logger) Event(text ...string) {
	log := strings.Join(text, "][")
	timenow := "[" + time.Now().Format("2006-01-02 15:04:05") + "]"
	logger.EventLog.Info().Msg(timenow + log)

}

// Message is use to write Message log
func (logger *Logger) Message(text ...string) {
	log := "[" + strings.Join(text, "][") + "]"
	timenow := "[" + time.Now().Format("2006-01-02 15:04:05") + "]"
	logger.MessageLog.Info().Msg(timenow + log)
}

// Fatal is use to write error log and stop program
func (logger *Logger) Fatal(wg *sync.WaitGroup, text ...string) {
	log := "[" + strings.Join(text, "][") + "]"
	timenow := "[" + time.Now().Format("2006-01-02 15:04:05") + "]"
	logger.MessageLog.Info().Msg(timenow + log)
	wg.Done()
}

// SetupLogging is used to set up logging system
func SetupLogging(path string) {
	errorOutput := GetOutputLog(path+"/event/", "error.log")
	eventOutput := GetOutputLog(path+"/error/", "event.log")
	messageOutput := GetOutputLog(path+"/message/", "message.log")

	logger := Logger{
		ErrorLog:   zerolog.New(errorOutput).With().Logger(),
		EventLog:   zerolog.New(eventOutput).With().Logger(),
		MessageLog: zerolog.New(messageOutput).With().Logger(),
	}
	Log = &logger
}

//GetOutputLog is used handling log output result
func GetOutputLog(path, filename string) zerolog.ConsoleWriter {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		os.MkdirAll(path, 0775)
	}
	logfile, err := os.OpenFile(path+filename, os.O_APPEND|os.O_CREATE|os.O_RDWR, 0775)
	if err != nil {
		fmt.Printf("error opening file: %v" + err.Error())
		os.Exit(2)
	}
	writeLog := diode.NewWriter(logfile, 1000, 10*time.Millisecond, func(missed int) {
		fmt.Printf("Logger Dropped %d messages", missed)
	})

	output := zerolog.ConsoleWriter{Out: writeLog}
	output.FormatLevel = func(i interface{}) string { return "" }
	output.FormatTimestamp = func(i interface{}) string { return "" }
	output.FormatMessage = func(i interface{}) string { return fmt.Sprintf("%s", i) }
	output.FormatFieldName = func(i interface{}) string { return "" }
	output.FormatFieldValue = func(i interface{}) string { return "" }
	return output
}
