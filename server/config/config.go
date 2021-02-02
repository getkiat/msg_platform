package config

import (
	"log"
	"os"
	"strconv"
	"sync"
	"time"
)

const (
	defaultHTTPPort = 6001
	sessionDuration = time.Minute * 20
)

// Config holds necessary config values.
type Config struct {
	HTTPPort int
}

var (
	once   sync.Once
	config Config
)

func LoadConfig() Config {
	once.Do(func() {
		config = Config{
			HTTPPort: parseHTTPPort(),
		}
	})

	return config
}

func parseHTTPPort() int {
	httpPortStr, ok := os.LookupEnv("HTTP_PORT")
	if !ok || httpPortStr == "" {
		return defaultHTTPPort
	}

	httpPort, err := strconv.ParseInt(httpPortStr, 10, 32)
	if err != nil {
		log.Fatalf("HTTP_PORT value '%s' could not be parsed due to %s", httpPortStr, err.Error())
	}

	return int(httpPort)
}