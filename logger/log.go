package logger

import (
	"log"
	"os"
)

var Logfile, LogErr = os.Create("app.log")

func init() {
	if LogErr != nil {
		log.Fatal(LogErr)
	}
	log.SetOutput(Logfile)
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
}

var infoLogger = log.New(Logfile, "INFO: ", log.Ldate|log.Ltime)
var debugLogger = log.New(Logfile, "DEBUG: ", log.Ldate|log.Ltime)

type AggregatedLogger struct {
	InfoLogger  *log.Logger
	DebugLogger *log.Logger
}

var Logg = AggregatedLogger{
	InfoLogger:  infoLogger,
	DebugLogger: debugLogger,
}

func (l AggregatedLogger) Info(v ...interface{}) {
	l.InfoLogger.Println(v...)
}

func (l AggregatedLogger) Debug(v ...interface{}) {
	l.DebugLogger.Println(v...)
}
