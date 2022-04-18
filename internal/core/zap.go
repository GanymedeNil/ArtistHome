package core

import (
	"fmt"
	"os"
	"time"

	"github.com/GanymedeNil/GoFrameworkBase/internal/global"
	"github.com/GanymedeNil/GoFrameworkBase/internal/util"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var level zapcore.Level

func Zap() {
	if ok, _ := util.PathExists(global.CONFIG.Zap.Directory); !ok {
		fmt.Printf("create %v directory\n", global.CONFIG.Zap.Directory)
		_ = os.Mkdir(global.CONFIG.Zap.Directory, os.ModePerm)
	}
	switch global.CONFIG.Zap.Level {
	case "debug":
		level = zap.DebugLevel
	case "info":
		level = zap.InfoLevel
	case "warn":
		level = zap.WarnLevel
	case "error":
		level = zap.ErrorLevel
	case "dpanic":
		level = zap.DPanicLevel
	case "panic":
		level = zap.PanicLevel
	case "fatal":
		level = zap.FatalLevel
	default:
		level = zap.InfoLevel

	}

	var logger *zap.Logger
	if level == zap.DebugLevel || level == zap.ErrorLevel {
		logger = zap.New(getEncoderCore(), zap.AddStacktrace(level))

	} else {
		logger = zap.New(getEncoderCore())
	}
	if global.CONFIG.Zap.ShowLine {
		logger = logger.WithOptions(zap.AddCaller())
	}
	global.LOGGER = logger
}

// getEncoderConfig get zapcore.EncoderConfig
func getEncoderConfig() (config zapcore.EncoderConfig) {
	config = zapcore.EncoderConfig{
		MessageKey:     "message",
		LevelKey:       "level",
		TimeKey:        "time",
		NameKey:        "logger",
		CallerKey:      "caller",
		StacktraceKey:  global.CONFIG.Zap.StacktraceKey,
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.LowercaseLevelEncoder,
		EncodeTime:     CustomTimeEncoder,
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.FullCallerEncoder,
	}
	switch {
	case global.CONFIG.Zap.EncodeLevel == "LowercaseLevelEncoder":
		config.EncodeLevel = zapcore.LowercaseLevelEncoder
	case global.CONFIG.Zap.EncodeLevel == "LowercaseColorLevelEncoder":
		config.EncodeLevel = zapcore.LowercaseColorLevelEncoder
	case global.CONFIG.Zap.EncodeLevel == "CapitalLevelEncoder":
		config.EncodeLevel = zapcore.CapitalLevelEncoder
	case global.CONFIG.Zap.EncodeLevel == "CapitalColorLevelEncoder":
		config.EncodeLevel = zapcore.CapitalColorLevelEncoder
	default:
		config.EncodeLevel = zapcore.LowercaseLevelEncoder
	}
	return config
}

// getEncoder get zapcore.Encoder
func getEncoder() zapcore.Encoder {
	if global.CONFIG.Zap.Format == "json" {
		return zapcore.NewJSONEncoder(getEncoderConfig())
	}
	return zapcore.NewConsoleEncoder(getEncoderConfig())
}

// getEncoderCore get Encoder zapcore.Core
func getEncoderCore() (core zapcore.Core) {
	writer := util.GetWriteSyncer(fmt.Sprintf("./%s/server.log", global.CONFIG.Zap.Directory))
	return zapcore.NewCore(getEncoder(), writer, level)
}

func CustomTimeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(t.Format(global.CONFIG.Zap.Prefix + "2006/01/02 - 15:04:05.000"))
}
