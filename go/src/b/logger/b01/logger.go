package logger

import (
	"fmt"
	"io"
	"os"
	"runtime"
	"strconv"
	"strings"
	"time"
)

// Error Levels that can be used to differentiate logged messages
// and also set the verbosity of logs to display.
const (
	// For unrecoverable errors where you would be unable to
	// continue the current scope of code.
	LogError int = iota

	// For non-critical errors that do not require you to abort/exit
	// from the current scope of code.
	LogWarn

	// For non-error "informational" logging.
	LogInfo

	// For any type of verbose debug specific logging.
	LogDebug
)

var (
	// PrefixError is the custom prefix for errors.
	PrefixError string

	// PrefixWarn is the custom prefix for warnings.
	PrefixWarn string

	// PrefixInfo is the custom prefix for informational messages.
	PrefixInfo string

	// PrefixDebug is the custom prefix for debug messages.
	PrefixDebug string

	// Prefix is a string that is added to the start of any logged messages.
	Prefix string

	// LogLevel is the level of messages to be logged.
	LogLevel int

	// Writer is the output io.Writer where messages are wrote to.
	Writer io.Writer
)

func init() {
	Prefix = `Log`
	LogLevel = 0
	Writer = os.Stderr

	PrefixError = strconv.Itoa(LogError)
	PrefixWarn = strconv.Itoa(LogWarn)
	PrefixInfo = strconv.Itoa(LogInfo)
	PrefixDebug = strconv.Itoa(LogDebug)
}

func getPrefix(level int) string {
	if level == LogError {
		return PrefixError
	} else if level == LogWarn {
		return PrefixWarn
	} else if level == LogInfo {
		return PrefixInfo
	} else {
		return PrefixDebug
	}
}

// Error logs a Error level message.
func Error(format string, args ...interface{}) {
	if LogError > LogLevel {
		return
	}
	Logger(Writer, LogError, 2, format, args...)
}

// Warn logs a Warn level message.
func Warn(format string, args ...interface{}) {
	if LogWarn > LogLevel {
		return
	}
	Logger(Writer, LogWarn, 2, format, args...)
}

// Info logs a Info level message.
func Info(format string, args ...interface{}) {
	if LogInfo > LogLevel {
		return
	}
	Logger(Writer, LogInfo, 2, format, args...)
}

// Debug logs a Debug level message.
func Debug(format string, args ...interface{}) {
	if LogDebug > LogLevel {
		return
	}
	Logger(Writer, LogDebug, 2, format, args...)
}

// Logger formats and writes the provided message to the defined io.Writer output
// as long as the passed level is less than or equal to the Logger.LogLevel.
func Logger(out io.Writer, level int, depth int, format string, args ...interface{}) {
	if level > LogLevel {
		return
	}
	now := time.Now()
	pc, file, line, _ := runtime.Caller(depth)
	files := strings.Split(file, "/")
	file = files[len(files)-1]

	name := runtime.FuncForPC(pc).Name()
	fns := strings.Split(name, ".")
	name = fns[len(fns)-1]
	message := fmt.Sprintf(format, args...)
	fmt.Fprintf(
		out,
		"%s [%s%s] %s:%d:%s() %s\n",
		now.Format("2006-01-02 15:04:05"),
		Prefix,
		getPrefix(level),
		file,
		line,
		name,
		message,
	)
}
