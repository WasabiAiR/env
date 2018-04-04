package env

import (
	"os"
	"testing"
	"time"

	"github.com/cheekybits/is"
)

func TestGetenvWithDefault(t *testing.T) {
	is := is.New(t)

	val := os.Getenv("env_test_key")
	is.Equal("", val)

	val = GetenvWithDefault("env_test_key", "1234")
	is.Equal("1234", val)

	os.Setenv("env_test_key", "3456")
	val = GetenvWithDefault("env_test_key", "1234")
	is.Equal("3456", val)

	// explicitly set to empty string
	os.Setenv("env_test_key", "")
	val = GetenvWithDefault("env_test_key", "1234")
	is.Equal("", val)
}

func TestGetenvBoolWithDefault(t *testing.T) {
	is := is.New(t)
	key := "env_bool_test_key"

	check := os.Getenv(key)
	is.Equal("", check)

	val := GetenvBoolWithDefault(key, true)
	is.Equal(true, val)

	os.Setenv(key, "3456")
	val = GetenvBoolWithDefault(key, true)
	is.Equal(false, val)

	// explicitly set to empty string
	os.Setenv(key, "")
	val = GetenvBoolWithDefault(key, true)
	is.Equal(false, val)

}

func TestGetenvIntWithDefault(t *testing.T) {
	is := is.New(t)
	key := "env_int_test_key"

	check := os.Getenv(key)
	is.Equal("", check)

	val := GetenvIntWithDefault(key, 1234)
	is.Equal(1234, val)

	os.Setenv(key, "3456")
	val = GetenvIntWithDefault(key, 1234)
	is.Equal(3456, val)

	// explicitly set to empty string
	os.Setenv(key, "")
	val = GetenvIntWithDefault(key, 1234)
	is.Equal(1234, val)
}

func TestGetenvDurationWithDefault(t *testing.T) {
	is := is.New(t)
	key := "env_duration_test_key"
	defaultValue := 300 * time.Second

	check := os.Getenv(key)
	is.Equal("", check)

	val := GetenvDurationWithDefault(key, defaultValue)
	is.Equal(300, val.Seconds())

	os.Setenv(key, "600s")
	val = GetenvDurationWithDefault(key, defaultValue)
	is.Equal(600, val.Seconds())

	// explicitly set to empty string
	os.Setenv(key, "")
	val = GetenvDurationWithDefault(key, defaultValue)
	is.Equal(300, val.Seconds())
}

func TestGetenvFloat64WithDefault(t *testing.T) {
	is := is.New(t)
	key := "env_float64_test_key"

	check := os.Getenv(key)
	is.Equal("", check)

	val := GetenvFloat64WithDefault(key, 123.4)
	is.Equal(123.4, val)

	os.Setenv(key, "345.6")
	val = GetenvFloat64WithDefault(key, 123.4)
	is.Equal(345.6, val)

	// explicitly set to empty string
	os.Setenv(key, "")
	val = GetenvFloat64WithDefault(key, 123.4)
	is.Equal(123.4, val)
}
