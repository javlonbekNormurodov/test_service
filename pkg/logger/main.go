package logger

import (
	"os"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type LoggingConfig struct {
	Level         string        `json:"level" yaml:"level" envDefault:"info"`
	LogFirstN     int           `json:"log_first_n" yaml:"log_first_n" envDefault:"3"`
	LogThereAfter int           `json:"log_there_after" yaml:"log_there_after" envDefault:"10"`
	LogInterval   time.Duration `json:"log_interval" yaml:"log_interval" envDefault:"1s"`
	ProjectName   string        `json:"project_name" yaml:"project_name"`
}

func lowerCaseLevelEncoder(
	level zapcore.Level,
	enc zapcore.PrimitiveArrayEncoder,
) {
	if level == zap.PanicLevel || level == zap.DPanicLevel {
		enc.AppendString("error")
		return
	}

	zapcore.LowercaseLevelEncoder(level, enc)
}

func New(config *LoggingConfig) *zap.Logger {
	stdout := zapcore.AddSync(os.Stdout)

	globalLogLevel, err := zapcore.ParseLevel(config.Level)
	if err != nil {
		panic(err)
	}

	level := zap.NewAtomicLevelAt(globalLogLevel)

	productionCfg := zap.NewProductionEncoderConfig()
	productionCfg.TimeKey = "timestamp"
	productionCfg.EncodeTime = zapcore.ISO8601TimeEncoder
	productionCfg.EncodeLevel = lowerCaseLevelEncoder
	productionCfg.StacktraceKey = "stack"

	jsonEncoder := zapcore.NewJSONEncoder(productionCfg)

	jsonOutCore := zapcore.NewCore(jsonEncoder, stdout, level)

	samplingCore := zapcore.NewSamplerWithOptions(
		jsonOutCore,
		config.LogInterval,
		config.LogFirstN,
		config.LogThereAfter,
	)

	return zap.New(samplingCore, zap.AddCaller(), zap.AddStacktrace(zap.ErrorLevel)).Named(config.ProjectName)
}
