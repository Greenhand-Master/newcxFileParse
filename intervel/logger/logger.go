package logger

import (
	"os"
	"strings"
	"time"

	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// InitLogger 日志初始化，warnFile、infoFile、debugFile 表示三种等级日志包含路径的文件名，可为空；newName 表示替换文件名，建议设置为"-%Y%m%d.log"；maxSaveTime 最长保存时间，建议设置为 time.Hour*24*30； rotationTime 日志切割时间，建议设置为 time.Hour*24
func InitLogger(warnFile, infoFile, debugFile, newName string, maxSaveTime, rotationTime time.Duration) {
	var cores []zapcore.Core

	if warnFile != "" {
		warnWriter := SetRotateRule(warnFile, newName, maxSaveTime, rotationTime) // "./log/warn/warn.log"
		warnCore := zapcore.NewCore(GetConfig(), warnWriter, zap.WarnLevel)
		cores = append(cores, warnCore)
	}

	if infoFile != "" {
		infoWriter := SetRotateRule(infoFile, newName, maxSaveTime, rotationTime) // "./log/info/info.log"
		infoCore := zapcore.NewCore(GetConfig(), infoWriter, zap.InfoLevel)
		cores = append(cores, infoCore)
	}

	if debugFile != "" {
		debugWriter := SetRotateRule(debugFile, newName, maxSaveTime, rotationTime) //"./log/cim.log"
		debugCore := zapcore.NewCore(GetConfig(), debugWriter, zap.DebugLevel)
		cores = append(cores, debugCore)
	}

	consoleCore := zapcore.NewCore(GetConfig(), zapcore.Lock(zapcore.AddSync(zapcore.AddSync(os.Stdout))), zap.DebugLevel)
	cores = append(cores, consoleCore)

	core := zapcore.NewTee(cores...)

	_logger := zap.New(core, zap.AddCaller()) // 将调用函数信息记录到日志

	zap.ReplaceGlobals(_logger)
}

func GetConfig() zapcore.Encoder {
	config := zap.NewDevelopmentEncoderConfig()

	config.EncodeTime = zapcore.ISO8601TimeEncoder
	config.EncodeLevel = zapcore.CapitalLevelEncoder // 让日志信息级别以大写输出

	return zapcore.NewConsoleEncoder(config)
}

// SetRotateRule 设置日志切割规则，fileName 包含路径的文件名，如"./cim.log"；newName 替换文件名，建议设置为"-%Y%m%d.log"或"-%Y%m%d%H%M"；maxSaveTime 最长保存时间，建议设置为 time.Hour*24*30； rotationTime 日志切割时间，建议设置为 time.Hour*24
func SetRotateRule(fileName, newName string, maxSaveTime, rotationTime time.Duration) zapcore.WriteSyncer {
	hook, err := rotatelogs.New(
		strings.Replace(fileName, ".log", "", -1)+newName,
		rotatelogs.WithMaxAge(maxSaveTime),
		rotatelogs.WithRotationTime(rotationTime),
	)

	if err != nil {
		zap.S().Panic(err)
	}

	return zapcore.AddSync(hook)
}
