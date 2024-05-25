package config

import (
	"log"

	"github.com/knadh/koanf"
	"github.com/knadh/koanf/parsers/json"
	"github.com/knadh/koanf/providers/env"
	"github.com/knadh/koanf/providers/file"
)

var k = koanf.New(".")

// Init loads configuration using Koanf into the global 'k' instance.
func Init() {
	// Load JSON config
	if err := k.Load(file.Provider("/var/www/jor.at/config.json"), json.Parser()); err != nil {
		log.Fatalf("error loading config: %v", err)
	}

	// Load environment variables and merge into the loaded config
	if err := k.Load(env.Provider("MYAPP_", ".", func(s string) string {
		return s
	}), nil); err != nil {
		log.Fatalf("error loading env vars: %v", err)
	}
}

// GetString fetches a string value by key from the loaded configuration
func GetString(key string) string {
	return k.String(key)
}

// GetInt fetches an integer value by key from the loaded configuration
func GetInt(key string) int {
	return k.Int(key)
}
