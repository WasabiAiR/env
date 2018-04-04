package env

import (
	"os"
	"strconv"
	"strings"
	"time"
)

// GetenvWithDefault returns a value from an environment variable defined by
// key. If key isn't set in the environment, returns defaultValue
func GetenvWithDefault(key string, defaultValue string) string {
	vars := os.Environ()

	for _, v := range vars {
		pieces := strings.SplitN(v, "=", 2)
		if pieces[0] == key {
			return pieces[1]
		}
	}
	return defaultValue
}

// GetenvBoolWithDefault returns a value from an environment variable defined by
// key. If key isn't set in the environment, returns defaultValue
func GetenvBoolWithDefault(key string, defaultValue bool) bool {
	vars := os.Environ()

	for _, v := range vars {
		pieces := strings.SplitN(v, "=", 2)
		if pieces[0] == key {
			if strings.ToLower(pieces[1]) == "true" {
				return true
			}
			return false
		}
	}
	return defaultValue
}

// GetenvIntWithDefault returns a value from an environment variable defined by
// key. If key isn't set in the environment, or doesn't parse as an int,
// returns defaultValue
func GetenvIntWithDefault(key string, defaultValue int) int {
	vars := os.Environ()

	for _, v := range vars {
		pieces := strings.SplitN(v, "=", 2)
		if pieces[0] == key {
			val, err := strconv.Atoi(pieces[1])
			if err != nil {
				return defaultValue
			}
			return val
		}
	}
	return defaultValue
}

// GetenvDurationWithDefault returns a value from an environment variable defined by
// key. If key isn't set in the environment, or doesn't parse as a time.Duration,
// returns defaultValue
func GetenvDurationWithDefault(key string, defaultValue time.Duration) time.Duration {
	vars := os.Environ()

	for _, v := range vars {
		pieces := strings.SplitN(v, "=", 2)
		if pieces[0] == key {
			val, err := time.ParseDuration(pieces[1])
			if err != nil {
				return defaultValue
			}
			return val
		}
	}
	return defaultValue
}

// GetenvFloat64WithDefault returns a value from an environment variable defined by
// key. If key isn't set in the environment, or doesn't parse as a float64,
// returns defaultValue
func GetenvFloat64WithDefault(key string, defaultValue float64) float64 {
	vars := os.Environ()

	for _, v := range vars {
		pieces := strings.SplitN(v, "=", 2)
		if pieces[0] == key {
			val, err := strconv.ParseFloat(pieces[1], 64)
			if err != nil {
				return defaultValue
			}
			return val
		}
	}
	return defaultValue
}
