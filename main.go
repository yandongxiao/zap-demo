package main

import (
	"time"

	"go.uber.org/zap"
)

// zap 要解决的问题：For applications that log in the hot path, reflection-based serialization
// and string formatting are prohibitively expensive
func main() {
	logger, _ := zap.NewProduction()
	defer logger.Sync() // flushes buffer, if any
	// In contexts where performance is nice, but not critical, use the SugaredLogger
	sugar := logger.Sugar()

	// 产生了一个key=world, value=123 的项
	sugar.With("world", 123).Info("message content")

	// 产生了一个key=logger, value=xxx 的项
	sugar.Named("reconciler").Info("log content")
	sugar.Named("dispatcher").Info("log content")

	// 格式化输出的key,value对
	sugar.Infow("failed to fetch URL", zap.String("url", "baidu.com"), zap.Any("attemp", 3))

	// When performance and type safety are critical, use the Logger.
	logger.Info("failed to fetch URL",
		// Structured context as strongly typed Field values.
		zap.String("url", "www.baidu.com"),
		zap.Int("attempt", 3),
		zap.Duration("backoff", time.Second),
	)
}