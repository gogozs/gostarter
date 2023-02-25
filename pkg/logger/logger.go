package logger

import (
	"os"
	"strings"
	"time"

	"github.com/gogozs/zlib/xlog"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"

	"github.com/gogozs/gostarter/configs"
)

// InitLogger global logger init
func InitLogger(config *configs.Config) {
	w := zapcore.AddSync(&lumberjack.Logger{
		Filename:   config.LogConfig.Filename,
		MaxSize:    config.LogConfig.MaxSize, // MB
		MaxBackups: config.LogConfig.MaxBackups,
		MaxAge:     config.LogConfig.MaxAge, // days
	})
	core := zapcore.NewCore(
		zapcore.NewConsoleEncoder(NewEncoderConfig()),
		zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout),
			w),
		getLevel(config.LogConfig.LogLevel),
	)
	l := zap.New(core, zap.AddCaller(), zap.AddCallerSkip(1)) // 跳过调用
	xlog.SetGlobalLogger(xlog.NewLogger(l.Sugar()))
}

func NewEncoderConfig() zapcore.EncoderConfig {
	return zapcore.EncoderConfig{
		// Keys can be anything except the empty string.
		TimeKey:        "T",
		LevelKey:       "L",
		NameKey:        "N",
		CallerKey:      "C",
		MessageKey:     "M",
		StacktraceKey:  "S",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.CapitalLevelEncoder,
		EncodeTime:     TimeEncoder,
		EncodeDuration: zapcore.StringDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	}
}

func TimeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(t.Format("2006-01-02 15:04:05.000"))
}

func getLevel(s string) (level zapcore.Level) {
	s = strings.ToUpper(s)
	switch s {
	case "DEBUG":
		level = zap.DebugLevel
	case "INFO":
		level = zap.InfoLevel
	case "WARN":
		level = zap.WarnLevel
	case "ERROR":
		level = zap.ErrorLevel
	case "PANIC":
		level = zap.PanicLevel
	default:
		level = zap.InfoLevel
	}
	return
}
