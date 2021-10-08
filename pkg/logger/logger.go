package logger

import (
	"fmt"
	"github.com/go-kratos/kratos/v2/log"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
)

var _ log.Logger = (*ZapLogger)(nil)

type ZapLogger struct {
	Logger *zap.Logger
	Sync   func() error
}

// NewZapLogger return ZapLogger
func NewZapLogger(encoder zapcore.EncoderConfig, level zap.AtomicLevel, opts ...zap.Option) *ZapLogger {
	core := zapcore.NewCore(
		zapcore.NewConsoleEncoder(encoder),
		zapcore.NewMultiWriteSyncer(
			zapcore.AddSync(os.Stdout),
		), level)
	zapLogger := zap.New(core, opts...)
	return &ZapLogger{Logger: zapLogger, Sync: zapLogger.Sync}
}

// Log Implementation of logger interface
func (l *ZapLogger) Log(level log.Level, keyvals ...interface{}) error {
	if len(keyvals) == 0 || len(keyvals)%2 != 0 {
		l.Logger.Warn(fmt.Sprint("Keyvalues must appear in pairs: ", keyvals))
		return nil
	}

	// Zap.Field is used when keyvals pairs appear
	var data []zap.Field
	for i := 0; i < len(keyvals); i += 2 {
		data = append(data, zap.Any(fmt.Sprint(keyvals[i]), fmt.Sprint(keyvals[i+1])))
	}
	switch level {
	case log.LevelDebug:
		l.Logger.Debug("", data...)
	case log.LevelInfo:
		l.Logger.Info("", data...)
	case log.LevelWarn:
		l.Logger.Warn("", data...)
	case log.LevelError:
		l.Logger.Error("", data...)
	}
	return nil
}

func GetEncoder() zapcore.Encoder {
	return zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig())
}

func GetLogWriter() zapcore.WriteSyncer {
	file, _ := os.Create("./test.log")
	return zapcore.AddSync(file)
}
