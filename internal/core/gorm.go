package core

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/GanymedeNil/GoFrameworkBase/internal/global"

	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	gormLogger "gorm.io/gorm/logger"
)

func Gorm() {
	var err error
	m := global.CONFIG.Mysql
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		m.Username, m.Password,
		m.Host, m.Port,
		m.Database)
	global.DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: new(global.LOGGER).LogMode(gormLogger.Info),
	})
	if err != nil {
		panic("failed to connect database")
	}
	if global.CONFIG.App.Debug {
		global.DB.Debug()
	}
	sqlDB, _ := global.DB.DB()
	sqlDB.SetMaxIdleConns(m.MaxIdleConns)
	sqlDB.SetMaxOpenConns(m.MaxOpenConns)
}

func new(zapLogger *zap.Logger) gormLogger.Interface {
	return &logger{
		ZapLogger:                 zapLogger,
		LogLevel:                  gormLogger.Warn,
		SlowThreshold:             100 * time.Millisecond,
		IgnoreRecordNotFoundError: false,
	}
}

type logger struct {
	ZapLogger                 *zap.Logger
	LogLevel                  gormLogger.LogLevel
	SlowThreshold             time.Duration
	IgnoreRecordNotFoundError bool
}

func (l *logger) LogMode(level gormLogger.LogLevel) gormLogger.Interface {
	newLogger := *l
	newLogger.LogLevel = level
	return &newLogger
}
func (l logger) Info(ctx context.Context, msg string, data ...interface{}) {
	if l.LogLevel >= gormLogger.Info {
		l.ZapLogger.Sugar().Infof(msg, data...)
	}
}
func (l logger) Warn(ctx context.Context, msg string, data ...interface{}) {
	if l.LogLevel >= gormLogger.Warn {
		l.ZapLogger.Sugar().Warnf(msg, data...)
	}
}
func (l logger) Error(ctx context.Context, msg string, data ...interface{}) {
	if l.LogLevel >= gormLogger.Error {
		l.ZapLogger.Sugar().Errorf(msg, data...)
	}
}
func (l logger) Trace(ctx context.Context, begin time.Time, fc func() (string, int64), err error) {
	if l.LogLevel <= gormLogger.Silent {
		return
	}
	elapsed := time.Since(begin)

	switch {
	case err != nil && l.LogLevel >= gormLogger.Error && (!errors.Is(err, gormLogger.ErrRecordNotFound) || !l.IgnoreRecordNotFoundError):
		sql, rows := fc()
		l.ZapLogger.Error("trace",
			zap.Error(err),
			zap.Duration("elapsed", elapsed),
			zap.Int64("rows", rows),
			zap.String("sql", sql))
	case elapsed > l.SlowThreshold && l.SlowThreshold != 0 && l.LogLevel >= gormLogger.Warn:
		sql, rows := fc()
		l.ZapLogger.Warn("trace",
			zap.Duration("elapsed", elapsed),
			zap.Int64("rows", rows),
			zap.String("sql", sql))
	case l.LogLevel == gormLogger.Info:
		sql, rows := fc()
		l.ZapLogger.Info("trace",
			zap.Duration("elapsed", elapsed),
			zap.Int64("rows", rows),
			zap.String("sql", sql))
	}
}
