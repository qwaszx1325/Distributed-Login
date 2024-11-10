package config

import (
	"log"

	"example.com/simple-login/pkg/cfgloader"
)

type (
	Host struct {
		ServiceName           string   `env:"SERVICE_NAME"`
		ServiceUrl            string   `env:"SERVICE_URL"`
		ServiceDomains        []string `env:"SERVICE_DOMAINS"`
		RateLimitIntervalSecs int      `env:"RATE_LIMIT_INTERVAL_SECS"`
		RateLimitMaxRequests  int      `env:"RATE_LIMIT_MAX_REQUESTS"`
		EnableTLS             bool     `env:"ENABLE_TLS"`
		CertFilePath          string   `env:"CERT_FILE_PATH"`
		KeyFilePath           string   `env:"KEY_FILE_PATH"`
	}

	Auth struct {
		AuthUrl string `env:"AUTH_URL"`
	}

	User struct {
		UserUrl string `env:"USER_URL"`
	}

	Config struct {
		Host
		Auth
		User
	}
)

func NewConfig() *Config {
	config, err := cfgloader.LoadConfigFromEnv[Config]()
	if err != nil {
		log.Fatalf("load config from env failed: %v", err)
	}
	return config
}
