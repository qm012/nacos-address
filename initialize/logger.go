package initialize

import (
	"fmt"
	"github.com/natefinch/lumberjack"
	"github/qm012/nacos-adress/global"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
)

func InitLogger() {
	logger := global.Server.Zap
	writeSyncer := zapcore.AddSync(&lumberjack.Logger{
		Filename:   logger.Filename,
		MaxSize:    logger.MaxSize,
		MaxAge:     logger.MaxAge,
		MaxBackups: logger.MaxBackup,
	})
	l := new(zapcore.Level)
	if err := l.UnmarshalText([]byte(logger.Level)); err != nil {
		panic(fmt.Sprintf("init logger error: %v", err.Error()))
	}
	core := zapcore.NewTee(
		zapcore.NewCore(getEncoder(), writeSyncer, l),
		zapcore.NewCore(getEncoder(), os.Stdout, l))
	global.Log = zap.New(core, zap.AddCaller())
	zap.ReplaceGlobals(global.Log)
}

func getEncoder() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.TimeKey = "time"
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	encoderConfig.EncodeDuration = zapcore.SecondsDurationEncoder
	encoderConfig.EncodeCaller = zapcore.ShortCallerEncoder
	return zapcore.NewJSONEncoder(encoderConfig)
}
