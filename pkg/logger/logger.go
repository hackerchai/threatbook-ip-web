package logger

import (
	"fmt"
	"github.com/hackerchai/threatbook-ip-web/internal/app/global"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"strings"
)

func Init() {
	switch strings.ToLower(global.CONFIG.Log.Type) {
	case "file":
		w := zapcore.AddSync(&lumberjack.Logger{
			Filename:   global.CONFIG.Log.Filename,
			MaxSize:    global.CONFIG.Log.MaxSize,
			MaxAge:     global.CONFIG.Log.MaxAge,
			MaxBackups: global.CONFIG.Log.MaxBackups,
			LocalTime:  global.CONFIG.Log.LocalTime,
			Compress:   global.CONFIG.Log.Compress,
		})
		// Create a zap core object for JSON encoding
		core := zapcore.NewCore(
			zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig()),
			w,
			global.CONFIG.Log.GetLogLevel(),
		)
		// Create a zap logger object
		global.Logger = zap.New(core)
		break
	// Use Access Logger in console based on environment by default
	default:
		var err error
		if strings.ToLower(global.CONFIG.Log.Environment) == "production" {
			global.Logger, err = zap.NewProduction()
		} else {
			global.Logger, err = zap.NewDevelopment()
		}
		if err != nil {
			fmt.Println(err.Error())
		}
		break
	}

	// Flush logger buffers, if any
	defer global.Logger.Sync()
}

func Debug(msg string, fields ...zap.Field) {
	global.Logger.Debug(msg, fields...)
}

func Info(msg string, fields ...zap.Field) {
	global.Logger.Info(msg, fields...)
}

func Warn(msg string, fields ...zap.Field) {
	global.Logger.Warn(msg, fields...)
}

func Error(msg string, fields ...zap.Field) {
	global.Logger.Error(msg, fields...)
}

func DPanic(msg string, fields ...zap.Field) {
	global.Logger.DPanic(msg, fields...)
}

func Panic(msg string, fields ...zap.Field) {
	global.Logger.Panic(msg, fields...)
}

func Fatal(msg string, fields ...zap.Field) {
	global.Logger.Fatal(msg, fields...)
}
