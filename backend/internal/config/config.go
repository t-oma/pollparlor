// Package config provides configuration for the application
package config

import (
	"log"
	"strings"

	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)

// App represents the application configuration
type App struct {
	Env  string `mapstructure:"env"`
	Addr string `mapstructure:"addr"`
	Port string `mapstructure:"port"`
}

// Log represents the log configuration
type Log struct {
	Level  string `mapstructure:"level"`
	Format string `mapstructure:"format"`
}

// Mongo represents the mongoDB configuration
type Mongo struct {
	URI string `mapstructure:"uri"`
	DB  string `mapstructure:"db"`
}

// Config represents the application configuration
type Config struct {
	App   App   `mapstructure:"app"`
	Log   Log   `mapstructure:"log"`
	Mongo Mongo `mapstructure:"mongo"`
}

// Load loads the configuration from the file
func Load() (*Config, error) {
	_ = godotenv.Load("configs/.env")

	v := viper.New()
	v.SetConfigName("config")
	v.SetConfigType("yaml")
	v.AddConfigPath("configs/")
	v.AutomaticEnv()
	v.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	v.SetEnvPrefix("APP") // APP_APP_PORT, APP_LOG_LEVEL, ...

	v.SetDefault("app.env", "dev")
	v.SetDefault("app.addr", "127.0.0.1")
	v.SetDefault("app.port", "8080")
	v.SetDefault("log.level", "info")
	v.SetDefault("log.format", "console")
	v.SetDefault("mongo.uri", "mongodb://localhost:27017")
	v.SetDefault("mongo.db", "pollparlor")

	if err := v.ReadInConfig(); err != nil {
		// don't panic if config file not found (e.g. in dev mode)
		log.Println("config file not found")
	}

	var cfg Config
	if err := v.Unmarshal(&cfg); err != nil {
		return nil, err
	}
	return &cfg, nil
}
