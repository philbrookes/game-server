package config

import "errors"

//Config is a set of config values used by this program
type Config struct {
	DatabaseName      string
	UserTable         string
	MongoHost         string
	MongoPort         string
	OutputFormat      string
	WebsocketListener string
	HTTPPort          string
}

//Get a config for the provided environment
func Get(env string) (*Config, error) {
	switch env {
	case "production":
		return GetProduction(), nil
	case "development":
		return GetDevelopment(), nil
	default:
		return nil, errors.New("Unknown game environment requested: " + env)
	}
}

//GetProduction creates a config suitable for production
func GetProduction() *Config {
	config := Config{
		DatabaseName:      "game",
		UserTable:         "user",
		MongoHost:         "mongo",
		OutputFormat:      "json",
		WebsocketListener: "connect",
		HTTPPort:          "80",
	}

	return &config
}

//GetDevelopment creates a config suitable for development
func GetDevelopment() *Config {
	config := Config{
		DatabaseName:      "game",
		UserTable:         "user",
		MongoHost:         "localhost",
		OutputFormat:      "json",
		WebsocketListener: "connect",
		HTTPPort:          "8080",
	}

	return &config
}
