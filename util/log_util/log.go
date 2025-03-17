package log_util

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"sync"
	"time"

	"jeanfo_mix/config"
	"jeanfo_mix/util"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

var (
	logger      *zap.Logger
	errorLogger *zap.Logger
	once        sync.Once
)

type LogLevel zapcore.Level

const (
	DebugLevel LogLevel = iota - 1
	InfoLevel
	WarnLevel
	ErrorLevel
	DPanicLevel
	PanicLevel
	FatalLevel
)

func Init() error {
	var err error
	once.Do(func() {
		cfg := config.GetConfig().Log

		// 获取项目根目录
		rootDir := util.GetProjRoot()

		// 创建日志目录
		logDir := cfg.Dir
		if logDir == "" || logDir == "./log" {
			logDir = filepath.Join(rootDir, "log")
		} else if !filepath.IsAbs(logDir) {
			logDir = filepath.Join(rootDir, logDir)
		}
		fmt.Println("Init log dir: " + logDir)
		if err := os.MkdirAll(logDir, os.ModePerm); err != nil {
			panic(fmt.Sprintf("Create log dir failed: %v", err))
		}

		// 配置日志级别
		level := getZapLevel(cfg.Level)

		// 普通日志配置
		normalWriter := &lumberjack.Logger{
			Filename:   filepath.Join(logDir, "app.log"),
			MaxSize:    cfg.Normal.MaxSize,
			MaxBackups: cfg.Normal.MaxBackups,
			Compress:   true,
		}

		// 错误日志配置
		errorWriter := &lumberjack.Logger{
			Filename:   filepath.Join(logDir, "error.log"),
			MaxSize:    cfg.Error.MaxSize,
			MaxBackups: cfg.Error.MaxBackups,
			Compress:   true,
		}

		// 创建核心
		cores := []zapcore.Core{
			newCore(normalWriter, level, cfg.Console),
			newErrorCore(errorWriter, zapcore.ErrorLevel, cfg.Console),
		}

		// 创建logger
		logger = zap.New(zapcore.NewTee(cores...), zap.AddCaller(), zap.AddStacktrace(zapcore.ErrorLevel))
		errorLogger = logger.WithOptions(zap.AddStacktrace(zapcore.PanicLevel))

		// 替换标准库logger
		zap.ReplaceGlobals(logger)
		zap.RedirectStdLog(logger)
	})

	return err
}

func newCore(writer io.Writer, level zapcore.Level, console bool) zapcore.Core {
	encoderConfig := zapcore.EncoderConfig{
		TimeKey:       "time",
		LevelKey:      "level",
		NameKey:       "logger",
		CallerKey:     "caller",
		FunctionKey:   zapcore.OmitKey,
		MessageKey:    "msg",
		StacktraceKey: "stacktrace",
		LineEnding:    zapcore.DefaultLineEnding,
		EncodeTime: func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
			enc.AppendString(fmt.Sprintf("[%s]", t.Format("2006-01-02 15:04:05")))
		},
		EncodeLevel: func(l zapcore.Level, enc zapcore.PrimitiveArrayEncoder) {
			enc.AppendString(fmt.Sprintf("[%s]", l.CapitalString()))
		},
		EncodeDuration:   zapcore.StringDurationEncoder,
		ConsoleSeparator: " ",
	}

	var ws zapcore.WriteSyncer
	if console {
		ws = zapcore.NewMultiWriteSyncer(zapcore.AddSync(writer), zapcore.AddSync(os.Stdout))
	} else {
		ws = zapcore.AddSync(writer)
	}

	return zapcore.NewCore(
		zapcore.NewConsoleEncoder(encoderConfig),
		ws,
		level,
	)
}

func newErrorCore(writer io.Writer, level zapcore.Level, console bool) zapcore.Core {
	encoderConfig := zapcore.EncoderConfig{
		TimeKey:       "time",
		LevelKey:      "level",
		NameKey:       "logger",
		CallerKey:     "caller",
		FunctionKey:   zapcore.OmitKey,
		MessageKey:    "msg",
		StacktraceKey: "stacktrace",
		LineEnding:    zapcore.DefaultLineEnding,
		EncodeTime: func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
			enc.AppendString(fmt.Sprintf("[%s]", t.Format("2006-01-02 15:04:05")))
		},
		EncodeLevel: func(l zapcore.Level, enc zapcore.PrimitiveArrayEncoder) {
			enc.AppendString(fmt.Sprintf("[%s]", l.CapitalString()))
		},
		// EncodeCaller: func(caller zapcore.EntryCaller, enc zapcore.PrimitiveArrayEncoder) {
		// 	enc.AppendString(fmt.Sprintf("[%s:%d]", filepath.Base(caller.File), caller.Line))
		// },
		EncodeDuration:   zapcore.StringDurationEncoder,
		ConsoleSeparator: " ",
	}

	ws := zapcore.AddSync(writer)

	return zapcore.NewCore(
		zapcore.NewConsoleEncoder(encoderConfig),
		ws,
		level,
	)
}

func getZapLevel(level string) zapcore.Level {
	switch strings.ToUpper(level) {
	case "DEBUG":
		return zapcore.DebugLevel
	case "INFO":
		return zapcore.InfoLevel
	case "WARN":
		return zapcore.WarnLevel
	case "ERROR":
		return zapcore.ErrorLevel
	case "DPANIC":
		return zapcore.DPanicLevel
	case "PANIC":
		return zapcore.PanicLevel
	case "FATAL":
		return zapcore.FatalLevel
	default:
		return zapcore.DebugLevel
	}
}

func getCallerInfo() string {
	_, file, line, ok := runtime.Caller(3)
	if !ok {
		return ""
	}
	return fmt.Sprintf("%s:%d", filepath.Base(file), line)
}

func fillCallerInfo(msg string) string {
	caller := getCallerInfo()
	msg = fmt.Sprintf("[%s] %s", caller, msg)
	return msg
}

// Debug 打印调试日志
func Debug(format string, a ...any) {
	msg := fmt.Sprintf(format, a...)
	msg = fillCallerInfo(msg)
	logger.Debug(msg)
}

// Info 打印信息日志
func Info(format string, a ...any) {
	msg := fmt.Sprintf(format, a...)
	msg = fillCallerInfo(msg)
	logger.Info(msg)
}

// Warn 打印警告日志
func Warn(format string, a ...any) {
	msg := fmt.Sprintf(format, a...)
	msg = fillCallerInfo(msg)
	logger.Warn(msg)
}

// Error 打印错误日志
func Error(format string, a ...any) {
	msg := fmt.Sprintf(format, a...)
	msg = fillCallerInfo(msg)
	// fields = append(fields, zap.String("caller", caller))
	errorLogger.Error(msg)
}

// Panic 打印panic日志
func Panic(format string, a ...any) {
	msg := fmt.Sprintf(format, a...)
	msg = fillCallerInfo(msg)
	errorLogger.Panic(msg)
}

// Fatal 打印致命错误日志
func Fatal(format string, a ...any) {
	msg := fmt.Sprintf(format, a...)
	msg = fillCallerInfo(msg)
	errorLogger.Fatal(msg)
}

// Sync 刷新日志缓冲区
func Sync() {
	_ = logger.Sync()
	_ = errorLogger.Sync()
}
