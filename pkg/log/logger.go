package log

import (
	"errors"
	"fmt"
	"os"
	"strings"
	"time"
)

type Level int

const (
	Trace Level = iota
	Debug
	Info
	Warning
	Error
	Critical
)

type Log interface {
	ChangeLogLevel(lvl Level) error
	GetLogLevel() Level
	Trace(format string, a ...interface{})
	Debug(format string, a ...interface{})
	Info(format string, a ...interface{})
	Warn(format string, a ...interface{})
	Error(format string, a ...interface{})
	Critical(format string, a ...interface{})
}

type log struct {
	logLevel Level
	name string
}

func Create(lvl Level) (Log, error) {
	if lvl < Trace || lvl > Critical {
		return nil, errors.New("invalid log level")
	}
	return &log{logLevel: lvl}, nil
}

func CreateNamed(lvl Level, name string) (Log, error) {
	if lvl < Trace || lvl > Critical {
		return nil, errors.New("invalid log level")
	}
	if len(name) > 20 {
		return nil, errors.New("name cannot be longer than 20 characters")
	}
	if name == "" {
		return nil, errors.New("name may not be empty")
	}
	return &log{logLevel: lvl, name: name}, nil
}

func (log *log) GetLogLevel() Level {
	return log.logLevel
}

func (log *log) ChangeLogLevel(lvl Level) error {
	if lvl < Trace || lvl > Critical {
		return errors.New("invalid log level")
	}

	log.logLevel = lvl
	return nil
}

func (log *log) Trace(format string, a ...interface{}) {
	if log.logLevel <= Trace {
		log.stndPrint(Trace, format, a)
	}
}
func (log *log) Debug(format string, a ...interface{}) {
	if log.logLevel <= Debug {
		log.stndPrint(Debug, format, a)
	}
}

func (log *log) Info(format string, a ...interface{}) {
	if log.logLevel <= Info {
		log.stndPrint(Info, format, a)
	}
}

func (log *log) Warn(format string, a ...interface{}) {
	if log.logLevel <= Warning {
		log.errPrint(Warning, format, a)
	}
}

func (log *log) Error(format string, a ...interface{}) {
	if log.logLevel <= Error {
		log.errPrint(Error, format, a)
	}
}

func (log *log) Critical(format string, a ...interface{}) {
	if log.logLevel <= Critical {
		log.errPrint(Critical, format, a)
		os.Exit(-1)
	}
}

func (log *log) stndPrint(lvl Level, format string, a []interface{}) {
	s := log.generateString(lvl, format, a)
	_, err := os.Stdout.WriteString(s)
	// TODO find a better way to handle this lol
	if err != nil {
		panic(err)
	}
}

func (log *log) errPrint(lvl Level, format string, a []interface{}) {
	s := log.generateString(lvl, format, a)
	_, err := os.Stderr.WriteString(s)
	// TODO find a better way to handle this lol
	if err != nil {
		panic(err)
	}
}

func (log *log) generateString(lvl Level, format string, a []interface{}) string {
	lvlText := getLogLvlText(lvl)
	t := time.Now()
	var sb strings.Builder
	sb.WriteString(t.Format("2006-01-02 15:04:05"))
	sb.WriteString(" [")
	sb.WriteString(lvlText)
	sb.WriteString("]")
	if log.name != "" {
		sb.WriteString(" [")
		sb.WriteString(log.name)
		sb.WriteString("]")
	}
	sb.WriteString(" : ")

	if len(a) == 0 {
		sb.WriteString(format)
	} else {
		s := fmt.Sprintf(format, a...)
		sb.WriteString(s)
	}
	sb.WriteString("\n")
	return sb.String()
}

func getLogLvlText(lvl Level) string {
	switch lvl {
	case Trace:
		return "TRC"
	case Debug:
		return "DBG"
	case Info:
		return "INF"
	case Warning:
		return "WRN"
	case Error:
		return "ERR"
	case Critical:
		return "CRT"
	default:
		return "UNK"
	}
}
