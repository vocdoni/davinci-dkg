package log

import (
	"cmp"
	"fmt"
	"io"
	"os"
	"path"
	"sync"
	"time"

	"github.com/rs/zerolog"
)

const (
	LogLevelDebug = "debug"
	LogLevelInfo  = "info"
	LogLevelWarn  = "warn"
	LogLevelError = "error"
)

var (
	global zerolog.Logger
	mu     sync.RWMutex
)

func init() {
	Init(cmp.Or(os.Getenv("LOG_LEVEL"), LogLevelError), "stderr", nil)
}

// Init configures the package logger.
func Init(level, output string, errorOutput io.Writer) {
	var out io.Writer
	switch output {
	case "stdout":
		out = os.Stdout
	case "stderr":
		out = os.Stderr
	default:
		f, err := os.OpenFile(output, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0o600)
		if err != nil {
			panic(fmt.Sprintf("open log output: %v", err))
		}
		out = f
	}

	writer := zerolog.ConsoleWriter{
		Out:        out,
		TimeFormat: time.RFC3339,
	}
	if errorOutput != nil {
		out = zerolog.MultiLevelWriter(writer, errorOutput)
	} else {
		out = writer
	}

	logger := zerolog.New(out).With().Timestamp().Caller().Logger()
	zerolog.CallerSkipFrameCount = 3
	zerolog.CallerMarshalFunc = func(pc uintptr, file string, line int) string {
		return fmt.Sprintf("%s/%s:%d", path.Base(path.Dir(file)), path.Base(file), line)
	}

	switch level {
	case LogLevelDebug:
		logger = logger.Level(zerolog.DebugLevel)
	case LogLevelInfo:
		logger = logger.Level(zerolog.InfoLevel)
	case LogLevelWarn:
		logger = logger.Level(zerolog.WarnLevel)
	default:
		logger = logger.Level(zerolog.ErrorLevel)
	}

	mu.Lock()
	global = logger
	mu.Unlock()
}

// Logger returns the configured logger.
func Logger() zerolog.Logger {
	mu.RLock()
	defer mu.RUnlock()
	return global
}

// Debugw logs a debug message with structured fields.
func Debugw(msg string, keysAndValues ...any) { event(zerolog.DebugLevel, msg, keysAndValues...) }

// Infow logs an info message with structured fields.
func Infow(msg string, keysAndValues ...any) { event(zerolog.InfoLevel, msg, keysAndValues...) }

// Warnw logs a warning message with structured fields.
func Warnw(msg string, keysAndValues ...any) { event(zerolog.WarnLevel, msg, keysAndValues...) }

// Errorw logs an error message with structured fields.
func Errorw(err error, msg string, keysAndValues ...any) {
	fields := append([]any{"error", err}, keysAndValues...)
	event(zerolog.ErrorLevel, msg, fields...)
}

// Info logs a plain info message.
func Info(msg string) {
	logger := Logger()
	logger.Info().Msg(msg)
}

// Warn logs a plain warning message.
func Warn(msg string) {
	logger := Logger()
	logger.Warn().Msg(msg)
}

func event(level zerolog.Level, msg string, keysAndValues ...any) {
	logger := Logger()
	e := logger.WithLevel(level)
	for i := 0; i+1 < len(keysAndValues); i += 2 {
		key := fmt.Sprint(keysAndValues[i])
		e = e.Interface(key, keysAndValues[i+1])
	}
	e.Msg(msg)
}
