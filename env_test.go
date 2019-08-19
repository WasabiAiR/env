package env

import (
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestGetenvWithDefault(t *testing.T) {
	val := os.Getenv("env_test_key")
    require.Equal(t, "", val)

	val = GetenvWithDefault("env_test_key", "1234")
	require.Equal(t, "1234", "1234", val)

	os.Setenv("env_test_key", "3456")
	val = GetenvWithDefault("env_test_key", "1234")
	require.Equal(t, "3456", val)

	// explicitly set to empty string
	os.Setenv("env_test_key", "")
	val = GetenvWithDefault("env_test_key", "1234")
	require.Equal(t, "", val)
}

func TestGetenvBoolWithDefault(t *testing.T) {
	key := "env_bool_test_key"

	check := os.Getenv(key)
	require.Equal(t, "", check)

	require.True(t, GetenvBoolWithDefault(key, true))

	os.Setenv(key, "3456")
	require.True(t, GetenvBoolWithDefault(key, true))

	// explicitly set to empty string
	os.Setenv(key, "")
	require.True(t, GetenvBoolWithDefault(key, true))

	os.Setenv(key, "1")
	require.True(t, GetenvBoolWithDefault(key, false))

	os.Setenv(key, "T")
	require.True(t, GetenvBoolWithDefault(key, false))

	os.Setenv(key, "TRUE")
	require.True(t, GetenvBoolWithDefault(key, false))

	os.Setenv(key, "TrUe")
	require.True(t, GetenvBoolWithDefault(key, false))
}

func TestGetenvIntWithDefault(t *testing.T) {
	key := "env_int_test_key"

	check := os.Getenv(key)
	require.Equal(t, "", check)

	val := GetenvIntWithDefault(key, 1234)
	require.Equal(t, 1234, val)

	os.Setenv(key, "3456")
	val = GetenvIntWithDefault(key, 1234)
	require.Equal(t, 3456, val)

	// explicitly set to empty string
	os.Setenv(key, "")
	val = GetenvIntWithDefault(key, 1234)
	require.Equal(t, 1234, val)
}

func TestGetenvDurationWithDefault(t *testing.T) {
	key := "env_duration_test_key"
	defaultValue := 300 * time.Second

	check := os.Getenv(key)
	require.Equal(t, "", check)

	val := GetenvDurationWithDefault(key, defaultValue)
	require.Equal(t, 300.0, val.Seconds())

	os.Setenv(key, "600s")
	val = GetenvDurationWithDefault(key, defaultValue)
	require.Equal(t, 600.0, val.Seconds())

	// explicitly set to empty string
	os.Setenv(key, "")
	val = GetenvDurationWithDefault(key, defaultValue)
	require.Equal(t, 300.0, val.Seconds())
}

func TestGetenvFloat64WithDefault(t *testing.T) {
	key := "env_float64_test_key"

	check := os.Getenv(key)
	require.Equal(t, "", check)

	val := GetenvFloat64WithDefault(key, 123.4)
	require.Equal(t, 123.4, val)

	os.Setenv(key, "345.6")
	val = GetenvFloat64WithDefault(key, 123.4)
	require.Equal(t, 345.6, val)

	// explicitly set to empty string
	os.Setenv(key, "")
	val = GetenvFloat64WithDefault(key, 123.4)
	require.Equal(t, 123.4, val)
}
