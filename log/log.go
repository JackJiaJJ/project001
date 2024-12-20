package log

import (
	"fmt"
	"os"

	"github.com/cihub/seelog"
)

var (
	myLogger seelog.LoggerInterface
)

func init() {
	config := `
<?xml version="1.0" encoding="UTF-8"?>
<seelog minlevel="debug">
	<outputs formatid="main">  
		<console/>
	</outputs>
	<formats>
		<format id="main" format="[%UTCDate(2006-01-02 15:04:05.000)] [%LEVEL] %Msg%n" />
	</formats>
</seelog>	
`
	logger, _ := seelog.LoggerFromConfigAsBytes([]byte(config))
	err := seelog.ReplaceLogger(logger)
	if err != nil {
		fmt.Printf("seelog.ReplaceLogger error=%v\n", err)
	}
}

func ConfigureLogger(logDir string) {
	myLogger, err := seelog.LoggerFromConfigAsFile(logDir)
	if err != nil {
		panic(err)
	}

	err = seelog.ReplaceLogger(myLogger)
	if err != nil {
		fmt.Printf("seelog.ReplaceLogger error=%v\n", err)
	}
}

/******************************
Seelog functions:
******************************/

func discard(value string) {
}

func Tracef(format string, params ...interface{}) {
	seelog.Tracef(format, params...)
}

// Debugf formats message according to format specifier
// and writes to default logger with log level = Debug.
func Debugf(format string, params ...interface{}) {
	if os.Getenv("DEBUG_LOG") == "true" {
		seelog.Debugf(format, params...)
	}
}

// Infof formats message according to format specifier
// and writes to default logger with log level = Info.
func Infof(format string, params ...interface{}) {
	seelog.Infof(format, params...)
}

// Warnf formats message according to format specifier and writes to default logger with log level = Warn
func Warnf(format string, params ...interface{}) {
	err := seelog.Warnf(format, params...)
	discard(err.Error())
}

// Errorf formats message according to format specifier and writes to default logger with log level = Error
func Errorf(format string, params ...interface{}) {
	err := seelog.Errorf(format, params...)
	discard(err.Error())
}

// Criticalf formats message according to format specifier and writes to default logger with log level = Critical
func Criticalf(format string, params ...interface{}) error {
	return seelog.Criticalf(format, params...)
}

// Trace formats message using the default formats for its operands and writes to default logger with log level = Trace
func Trace(v ...interface{}) {
	seelog.Trace(v...)
}

// Debug formats message using the default formats for its operands and writes to default logger with log level = Debug
func Debug(v ...interface{}) {
	if os.Getenv("DEBUG_LOG") == "true" {
		seelog.Debug(v...)
	}
}

// Info formats message using the default formats for its operands and writes to default logger with log level = Info
func Info(v ...interface{}) {
	seelog.Info(v...)
}

// Warn formats message using the default formats for its operands and writes to default logger with log level = Warn
func Warn(v ...interface{}) {
	err := seelog.Warn(v...)
	discard(err.Error())
}

// Error formats message using the default formats for its operands and writes to default logger with log level = Error
func Error(v ...interface{}) {
	err := seelog.Error(v...)
	discard(err.Error())
}

// Critical formats message using the default formats for its operands and writes to default logger with log level = Critical
func Critical(v ...interface{}) error {
	return seelog.Critical(v...)
}

func Flush() {
	seelog.Flush()
}
