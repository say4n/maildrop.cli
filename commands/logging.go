package commands

import (
	"log"
	"os"
	"sync"
)

var Logger *log.Logger
var once sync.Once

func GetLoggerInstance() *log.Logger {
	once.Do(func() {
		Logger = createLogger()
	})
	return Logger
}

func createLogger() *log.Logger {
	return log.New(os.Stderr, "maildrop:", log.Lshortfile)
}
