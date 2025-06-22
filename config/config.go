package config

import (
	"fmt"
	"time"

	"github.com/ilyakaznacheev/cleanenv"
)

type (
	Config struct {
		Internal     InternalConfig     `yaml:"internal"`
		Postgres     PostgresConfig     `yaml:"postgres"`
		Cron         CronConfig         `yaml:"cron"`
		HTTP         HTTPConfig         `yaml:"http"`
		Log          LogConfig          `yaml:"log"`
		Response     ResponseConfig     `yaml:"response"`
		Cache        CacheConfig        `yaml:"Cache"`
	}

	InternalConfig struct {
		Path               string `yaml:"path" env:"API_PATH"`
		Environment        string `yaml:"environment" env:"ENVIRONMENT"`
		LogInputParamOnErr bool   `yaml:"logInputParamOnErr" env:"LOG_INPUT_PARAM_ON_ERR"`
	}

	CacheConfig struct {
		Url      string `yaml:"url" env:"CACHE_URL"`
		Username string `yaml:"username" env:"CACHE_USERNAME"`
		Password string `yaml:"password" env:"CACHE_PASSWORD"`
	}

	ResponseConfig struct {
		ExportError bool `yaml:"exportError" env:"RESPONSE_EXPORT_ERROR"`
	}

	LogConfig struct {
		Level              string `yaml:"level" env:"LOG_LEVEL"`
		RequestLogEnabled  bool   `yaml:"requestLogEnabled" env:"LOG_REQUEST_ENABLED"`
		RequestLogWithBody bool   `yaml:"requestLogWithBody" env:"LOG_REQUEST_WITH_BODY"`
	}

	PostgresConfig struct {
		Host                  string        `yaml:"host" env:"DB_HOST"`
		Port                  string        `yaml:"port" env:"DB_PORT"`
		Username              string        `yaml:"username" env:"DB_USERNAME"`
		Password              string        `yaml:"password" env:"DB_PASSWORD"`
		DBName                string        `yaml:"dbname" env:"DB_NAME"`
		Schema                string        `yaml:"schema" env:"DB_SCHEMA"`
		SSLMode               string        `yaml:"sslmode" env:"DB_SSL_MODE"`
		SSLRootCert           string        `yaml:"sslrootcert" env:"DB_SSL_ROOT_CERT"`
		PoolMax               int           `yaml:"poolMax" env:"DB_POOL_MAX"`
		PoolMin               int           `yaml:"poolMin" env:"DB_POOL_MIN"`
		HealthCheckPeriod     time.Duration `yaml:"healthCheckPeriod" env:"DB_HEALTH_CHECK_PERIOD"`
		ConnectionMaxIdleTime time.Duration `yaml:"connectionMaxIdleTime" env:"DB_CONNECTION_MAX_IDLE_TIME"`
		ConnectionMaxLifeTime time.Duration `yaml:"connectionMaxLifeTime" env:"DB_CONNECTION_MAX_LIFE_TIME"`
	}

	CronConfig struct {
	}

	HTTPConfig struct {
		Host               string        `yaml:"host" env:"HTTP_HOST"`
		Port               string        `yaml:"port" env:"HTTP_PORT"`
		ReadTimeout        time.Duration `yaml:"readTimeout" env:"HTTP_READ_TIMEOUT"`
		WriteTimeout       time.Duration `yaml:"writeTimeout" env:"HTTP_WRITE_TIMEOUT"`
		MaxHeaderMegabytes int           `yaml:"maxHeaderBytes"`
	}
)

func NewConfig(configDir string) (*Config, error) {
	cfg := &Config{}

	err := cleanenv.ReadConfig(configDir, cfg)
	if err != nil {
		return nil, fmt.Errorf("config error: %w", err)
	}

	err = cleanenv.ReadEnv(cfg)
	if err != nil {
		return nil, err
	}

	return cfg, nil
}
