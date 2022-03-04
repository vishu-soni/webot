package logger

import (
	"fmt"
	"os"
	"webot/config"

	log "github.com/sirupsen/logrus"
)

//inits new logrus instance
func init() {
	cfg := config.Logger()
	logDir := cfg.Dir
	logFile := cfg.Dir + "/" + cfg.File
	os.MkdirAll(logDir, 0755)
	file, err := os.OpenFile(logFile, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		log.Fatalf(fmt.Sprintf("Cannot initialise log file %s, %v", logFile, err))
	}
	log.SetLevel(log.InfoLevel)
	log.SetOutput(file)
}

// AddHook adds hook to global standard logger
func AddHook(hook log.Hook) {
	log.AddHook(hook)
}
