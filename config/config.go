package config

import (
	"flag"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type config struct {
	Name     string         `mapstructure:"name"`
	Env      Environment    `mapstructure:"env"`
	Postgres PostgresConfig `mapstructure:"postgres"`
	Logger   LoggerConfig   `mapstructure:"logger"`
	Sentry   SentryConfig   `mapstructure:"sentry"`
}

var (
	env = flag.String("env", "dev", "deployment environment for config. Default `env`")
	cfg *config
)

func init() {
	flag.Parse()
	setDefaults()
	viper.AddConfigPath("$GOPATH/src/sanchar")
	viper.AddConfigPath(".")
	viper.SetConfigType("yaml")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		log.WithField("err", err).Fatal("couldn't read base config")
	}
	viper.SetConfigName("config-" + *env)
	if err := viper.MergeInConfig(); err != nil {
		log.WithField("err", err).Fatalf("couldn't read config file %s", "config-"+*env)
	}
	if err := viper.Unmarshal(&cfg); err != nil {
		log.WithField("err", err).Fatal("couldn't marshall config")
	}
	err := isValid(cfg)
	if err != nil {
		log.WithField("err", err).Fatal("invalid configuration")
	}
}

func setDefaults() {
	// viper.SetDefault("key", "value")
}

// can implement validation on config in future
func isValid(cfg *config) error {
	return nil
}

// Sentry returns config for sentry
func Sentry() SentryConfig {
	return cfg.Sentry
}

// Logger returns config for logger
func Logger() LoggerConfig {
	return cfg.Logger
}

// returns value associated with a key in string format
func Name() string {
	return viper.GetString("name")
}

//Return config for postgres DB
func Postgres() PostgresConfig {
	return cfg.Postgres
}

// Env returns environment that is used for config
func Env() Environment {
	return cfg.Env
}
