package util

import (
	"os"

	"github.com/GanymedeNil/GoFrameworkBase/internal/global"

	"github.com/natefinch/lumberjack"
	"go.uber.org/zap/zapcore"
)

//@author: [SliverHorn](https://github.com/SliverHorn)
//@function: GetWriteSyncer
//@description: zap logger use file-rotatelogs
//@return: zapcore.WriteSyncer, error

func GetWriteSyncer(file string) zapcore.WriteSyncer {
	lumberJackLogger := &lumberjack.Logger{
		Filename:   file, // Location of log files
		MaxSize:    10,   // Maximum size (in MB) of log files before making cuts
		MaxBackups: 200,  // Maximum number of old files to keep
		MaxAge:     30,   // Maximum number of days to keep old files
		Compress:   true, // Whether to compress and archive old files
	}

	if global.CONFIG.Zap.LogInConsole {
		return zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout), zapcore.AddSync(lumberJackLogger))
	}
	return zapcore.AddSync(lumberJackLogger)
}
