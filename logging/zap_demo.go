package main

import (
	"encoding/json"
	"go.uber.org/zap"
	"io/ioutil"
	"os"
	"time"
)

func init()  {
	
}

func main()  {
	sugarLogger()
	logger()
	confLogger()
}

func confLogger()  {
	file, _ := os.Open("logging/zap_logger.conf")
	defer file.Close()

	bytes, _ := ioutil.ReadAll(file)

	var cfg zap.Config
	json.Unmarshal(bytes, &cfg)

	logger, _ := cfg.Build()
	defer logger.Sync()

	logger.Info("configured log")

}

/**
 * When performance and type safety are critical, use the Logger.
 * It's even faster than the SugaredLogger and allocates far less,
 * but it only supports structured logging.
 */
func logger()  {
	logger, _ := zap.NewProduction()
	defer logger.Sync()
	logger.Info("failed to fetch URL",
		// Structured context as strongly typed Field values.
		zap.String("url", "url test"),
		zap.Int("attempt", 3),
		zap.Duration("backoff", time.Second),
	)
}


/**
 * In contexts where performance is nice, but not critical
 * It includes both structured and printf-style APIs.
 */
func sugarLogger()  {
	var url = "url test"
	logger, _ := zap.NewProduction()
	defer logger.Sync() // flushes buffer, if any
	sugar := logger.Sugar()
	// structured
	sugar.Infow("failed to fetch URL",
		"url", url,
		"attempt", 3,
		"backoff", time.Second,
	)
	// printf-style
	sugar.Infof("Failed to fetch URL: %s", url)
}