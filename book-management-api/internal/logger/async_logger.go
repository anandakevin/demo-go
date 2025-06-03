package logger

import (
	"fmt"
	"os"
	"sync"
	"time"
)

type LogLevel string

const (
	InfoLevel  LogLevel = "INFO"
	ErrorLevel LogLevel = "ERROR"
	DebugLevel LogLevel = "DEBUG"
)

type AsyncLogger struct {
	channel chan string
	once    sync.Once
	file    *os.File
	mutex   sync.Mutex
}

func NewAsyncLogger() *AsyncLogger {
	file, err := os.OpenFile("app.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		panic(fmt.Sprintf("Failed to open log file: %v", err))
	}

	logger := &AsyncLogger{
		channel: make(chan string, 100),
		file:    file,
	}

	logger.once.Do(func() {
		go logger.worker()
	})

	return logger
}

func (l *AsyncLogger) Close() error {
	close(l.channel)
	return l.file.Close()
}

func (l *AsyncLogger) worker() {
	for msg := range l.channel {
		l.writeToFile(msg)
	}
}

func (l *AsyncLogger) writeToFile(msg string) {
	l.mutex.Lock()
	defer l.mutex.Unlock()

	timestamp := time.Now().Format("2006/01/02 15:04:05")
	fmt.Fprintf(l.file, "%s [LOG] %s\n", timestamp, msg)
	l.file.Sync()
}

func (l *AsyncLogger) log(level LogLevel, msg string) {
	logMsg := fmt.Sprintf("%s: %s", level, msg)
	select {
	case l.channel <- logMsg:
		// Message sent successfully
	default:
		// Channel is full, write directly to file
		l.writeToFile(logMsg)
	}
}

func (l *AsyncLogger) Info(msg string) {
	l.log(InfoLevel, msg)
}

func (l *AsyncLogger) Error(msg string) {
	l.log(ErrorLevel, msg)
}

func (l *AsyncLogger) Debug(msg string) {
	l.log(DebugLevel, msg)
}
