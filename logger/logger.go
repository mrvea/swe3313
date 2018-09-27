package logger

import (
	"fmt"
	"io"
	"log"
	"os"
	"sync"
)

const (
	// Info and Debug are the two possible logging levels
	// The only difference is the Debug method does nothing on info level
	_ = iota
	LevelInfo
	LevelDebug

	callDepth = 3
)

var (
	// defaults
	config = &Config{
		os.Stdout,
		log.Ldate | log.Ltime | log.Lshortfile,
		LevelInfo,
		false,
		"[DEBUG] ",
	}

	infoConfig = &Config{
		os.Stdout,
		log.Ldate | log.Ltime | log.Lshortfile,
		LevelInfo,
		false,
		"[INFO]  ",
	}

	// TODO - tj this is hacked to accomodate old API -
	// let's take a look and improve it for posterity
	// global logger singleton - for this project, it can be DEBUG
	instance      *Logger
	instantiation = sync.Once{}

	infoInstance      *Logger
	infoInstantiation = sync.Once{}
)

// Config allows selection of logger output, content and level (debug!)
// new - create a logger that discards all output
type Config struct {
	W           io.Writer
	Flag, Level int
	Discard     bool
	Prefix      string
}

// Logger wraps the standard logger and adds a debug level
type Logger struct {
	*log.Logger
	Level int
}

type discardWriter struct{}

func (*discardWriter) Write(p []byte) (int, error) {
	return len(p), nil
}

// Configure sets up the global logger.  This should be called from the main thread
// before the logger is created with Get
func Configure(c *Config) {
	if c.Discard {
		config.Discard = c.Discard // which would be true
		config.W = &discardWriter{}
	}
	// The presence of a writer overrules the discard option
	if c.W != nil {
		config.W = c.W
	}
	if c.Flag != 0 {
		// flags start at 1
		config.Flag = c.Flag
	}
	config.Level = c.Level
}

func ConfigureInfo(c *Config) {
	if c.Discard {
		infoConfig.Discard = c.Discard // which would be true
		infoConfig.W = &discardWriter{}
	}
	// The presence of a writer overrules the discard option
	if c.W != nil {
		infoConfig.W = c.W
	}
	if c.Flag != 0 {
		// flags start at 1
		infoConfig.Flag = c.Flag
	}
	infoConfig.Level = c.Level
}

// New makes a new logger with config.
func New(c *Config) *Logger {
	return &Logger{
		log.New(c.W, c.Prefix, c.Flag),
		c.Level,
	}
}

// Get returns the global singleton logger
func Get() *Logger {
	instantiation.Do(func() {
		instance = New(config)
	})

	return instance
}

func GetInfo() *Logger {
	infoInstantiation.Do(func() {
		infoInstance = New(infoConfig)
	})

	return infoInstance
}

func (l *Logger) Debug(v ...interface{}) {
	if l.Level < LevelDebug {
		return
	}
	l.Logger.Print(v...)
}

func (l *Logger) Debugf(format string, v ...interface{}) {
	if l.Level < LevelDebug {
		return
	}
	l.Logger.Printf(format, v...)
}

func (l *Logger) Debugln(v ...interface{}) {
	if l.Level < LevelDebug {
		return
	}

	l.Logger.Println(v...)
}

func Debug(v ...interface{}) {
	Get().Output(callDepth, fmt.Sprintln(v...))
}

func Debugf(format string, v ...interface{}) {
	Get().Output(callDepth, fmt.Sprintf(format, v...))
}

func Info(v ...interface{}) {
	GetInfo().Output(callDepth, fmt.Sprintln(v...))
}

func Infof(format string, v ...interface{}) {
	GetInfo().Output(callDepth, fmt.Sprintf(format, v...))
}
