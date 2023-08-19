package config

import "os"

var DatabaseURL string
var CacheURL string
var PROFILING bool

func init() {
	DatabaseURL = envOrFatal("DATABASE_URL")
	CacheURL = envOrFatal("CACHE_URL")
	PROFILING = os.Getenv("PROFILING") == "true"
}

func envOrFatal(key string) string {
	value := os.Getenv(key)
	if value == "" {
		panic("missing required environment variable " + key)
	}

	return value
}