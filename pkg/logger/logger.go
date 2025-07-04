package logger

import (
	"encoding/json"
	"fmt"
	"os"
	"runtime"
	"strings"
	"time"

	"github.com/fatih/color"
	"github.com/rs/zerolog"
)

type Logger struct {
	logger  zerolog.Logger
	service string
}

// New создает логгер с JSON-выводом и цветным форматированием для консоли
func New(serviceName string, prettyPrint bool) *Logger {
	var output zerolog.ConsoleWriter
	if prettyPrint {
		output = zerolog.ConsoleWriter{
			Out:        os.Stdout,
			TimeFormat: time.RFC3339,
			FormatLevel: func(i interface{}) string {
				return colorizeLevel(fmt.Sprintf("%-5s", i))
			},
			FormatMessage: func(i interface{}) string {
				return color.CyanString("message=%s", i)
			},
			FormatFieldName: func(i interface{}) string {
				return color.BlueString("%s=", i)
			},
			FormatFieldValue: func(i interface{}) string {
				return color.WhiteString("%v", i)
			},
		}
	} else {
		output = zerolog.ConsoleWriter{Out: os.Stdout, NoColor: true}
	}

	log := zerolog.New(output).
		With().
		Timestamp().
		Str("service", serviceName).
		Logger()

	return &Logger{
		logger:  log,
		service: serviceName,
	}
}

func colorizeLevel(level interface{}) string {
	l := fmt.Sprintf("%-5s", level)
	switch strings.ToLower(strings.TrimSpace(l)) {
	case "debug":
		return color.BlueString(l)
	case "info":
		return color.GreenString(l)
	case "warn":
		return color.YellowString(l)
	case "error", "fatal":
		return color.RedString(l)
	default:
		return l
	}
}

// getCallerInfo возвращает имя функции и файл, откуда был вызван логгер
func getCallerInfo() (string, string) {
	pc, file, line, ok := runtime.Caller(2) // 2 = глубина стека (пропускаем методы логгера)
	if !ok {
		return "unknown", "unknown"
	}

	// Получаем имя функции
	fnName := runtime.FuncForPC(pc).Name()
	lastSlash := strings.LastIndex(fnName, "/")
	if lastSlash >= 0 {
		fnName = fnName[lastSlash+1:]
	}

	// Укорачиваем путь к файлу
	lastSlash = strings.LastIndex(file, "/")
	if lastSlash >= 0 {
		file = file[lastSlash+1:]
	}

	return fmt.Sprintf("%s()", fnName), fmt.Sprintf("%s:%d", file, line)
}

// Методы логирования с JSON-форматированием
func (l *Logger) Debug(msg string, fields map[string]interface{}) {
	funcName, caller := getCallerInfo()
	l.logger.Debug().
		Str("func", funcName).
		Str("caller", caller).
		Fields(fields).
		Msg(msg)
}

func (l *Logger) Info(msg string, fields map[string]interface{}) {
	funcName, caller := getCallerInfo()
	l.logger.Info().
		Str("func", funcName).
		Str("caller", caller).
		Fields(fields).
		Msg(msg)
}

func (l *Logger) Warn(msg string, fields map[string]interface{}) {
	funcName, caller := getCallerInfo()
	l.logger.Warn().
		Str("func", funcName).
		Str("caller", caller).
		Fields(fields).
		Msg(msg)
}

func (l *Logger) Error(msg string, fields map[string]interface{}) {
	funcName, caller := getCallerInfo()
	l.logger.Error().
		Str("func", funcName).
		Str("caller", caller).
		Fields(fields).
		Msg(msg)
}

func (l *Logger) Fatal(msg string, fields map[string]interface{}) {
	funcName, caller := getCallerInfo()
	l.logger.Fatal().
		Str("func", funcName).
		Str("caller", caller).
		Fields(fields).
		Msg(msg)
	os.Exit(1)
}

// PrettyPrint красиво форматирует JSON (для debug)
func PrettyPrint(data interface{}) string {
	b, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return fmt.Sprintf("%v", data)
	}
	return string(b)
}
