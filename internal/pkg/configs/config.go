package configs

import (
	log "github.com/sirupsen/logrus"
	"os"
	"sync"
)

const (
	configPath = "configs/config.yaml"
)

type config struct {
	Log struct {
		Level string `yaml:"level"`
	} `yaml:"log"`
	Db struct {
		Host     string `yaml:"host"`
		Port     int    `yaml:"port"`
		Username string `yaml:"username"`
		Password string `yaml:"password"`
	} `yaml:"db"`
}

var (
	once     sync.Once
	instance *config
)

func New() *config {
	once.Do(func() {
		file, err := os.Open(configPath)
		if err != nil {

		}
	})
}

func (c *config) initLog() {
	logLevel := map[string]log.Level {
		"DEBUG": log.DebugLevel,
		"INFO": log.InfoLevel,
		"WARN": log.WarnLevel,
		"ERROR": log.ErrorLevel,
		"FATAL": log.FatalLevel,
		"PANIC": log.PanicLevel,
	}
	log.SetLevel(logLevel[c.Log.Level])
}
