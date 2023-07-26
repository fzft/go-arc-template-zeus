package template

import (
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

var logger *zap.Logger

func InitLogger() {
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}

	config := zap.Config{
		Level:       zap.NewAtomicLevelAt(getLogLevel(viper.GetString("logging.level"))),
		Development: false, // if true, DPanicLevel logs will panic
		Encoding:    viper.GetString("logging.encoding"),
		EncoderConfig: zapcore.EncoderConfig{
			TimeKey:        "ts",
			LevelKey:       "level",
			NameKey:        "logger",
			CallerKey:      "caller",
			MessageKey:     "msg",
			StacktraceKey:  "stacktrace",
			LineEnding:     zapcore.DefaultLineEnding,
			EncodeLevel:    zapcore.LowercaseLevelEncoder,
			EncodeTime:     zapcore.ISO8601TimeEncoder,
			EncodeDuration: zapcore.SecondsDurationEncoder,
			EncodeCaller:   zapcore.ShortCallerEncoder,
		},
		OutputPaths:      viper.GetStringSlice("logging.outputPaths"),
		ErrorOutputPaths: viper.GetStringSlice("logging.errorOutputPaths"),
	}

	var err error
	logger, err = config.Build()
	if err != nil {
		panic(err)
	}

	// rotate log
	rotateLog()
}

func rotateLog() {
	w := zapcore.AddSync(&lumberjack.Logger{
		Filename:   "/var/log/myapp.log",
		MaxSize:    viper.GetInt("logging.rotate.max_size"), // megabytes
		MaxBackups: viper.GetInt("logging.rotate.max_backups"),
		MaxAge:     viper.GetInt("logging.rotate.max_age"),   // days
		Compress:   viper.GetBool("logging.rotate.compress"), // disabled by default
	})

	core := zapcore.NewCore(
		zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig()),
		w,
		zap.InfoLevel,
	)

	logger := zap.New(core)
	defer logger.Sync() // flushes buffer, if any

	logger.Info("Zap logger initialized.")
}

func getLogLevel(level string) zapcore.Level {
	switch level {
	case "debug":
		return zapcore.DebugLevel
	case "info":
		return zapcore.InfoLevel
	case "warn":
		return zapcore.WarnLevel
	case "error":
		return zapcore.ErrorLevel
	default:
		return zapcore.InfoLevel
	}
}
