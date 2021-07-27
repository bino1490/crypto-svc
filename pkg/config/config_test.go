package config

import (
	"testing"

	"github.com/stretchr/testify/require"
)

// TestGetString checks for success scenario of
// string config retrieval
func TestGetString(t *testing.T) {
	profile := SrvConfig.GetString("profile")
	require.NotNil(t, profile)
}

// TestGetInt checks for success scenario of
// int config retrieval
// func TestGetInt(t *testing.T) {
// 	port := SrvConfig.GetInt("http.port")
// 	require.NotEmpty(t, port)
// }

// TestCastToIntFailure checks for failure scenario of
// int config cast
func TestCastToIntFailure(t *testing.T) {
	unknown := castToInt("")
	require.Empty(t, unknown)
}

// TestCastToInt64Failure checks for failure scenario of
// int64 config cast
func TestCastToInt64Failure(t *testing.T) {
	unknown := castToInt64("")
	require.Empty(t, unknown)
}

// TestCastToInt32Failure checks for failure scenario of
// int32 config cast
func TestCastToInt32Failure(t *testing.T) {
	unknown := castToInt32("")
	require.Empty(t, unknown)
}

// TestCastToFloat64Failure checks for failure scenario of
// float64 config cast
func TestCastToFloat64Failure(t *testing.T) {
	unknown := castToFloat64("")
	require.Empty(t, unknown)
}

// TestCastToFloat64Success checks for success scenario of
// float64 config cast
func TestCastToFloat64Success(t *testing.T) {
	float64Val := castToFloat64(float64(0.5))
	require.Equal(t, float64(0.5), float64Val)
}

// TestCastToFloat32Failure checks for failure scenario of
// float32 config cast
func TestCastToFloat32Failure(t *testing.T) {
	unknown := castToFloat32("")
	require.Empty(t, unknown)
}

// TestCastToFloat32Success checks for success scenario of
// float32 config cast
func TestCastToFloat32Success(t *testing.T) {
	float32Val := castToFloat32(float32(0.5))
	require.Equal(t, float32(0.5), float32Val)
}

// TestCastToBoolFailure checks for failure scenario of
// bool config cast
func TestCastToBoolFailure(t *testing.T) {
	unknown := castToBoolean("")
	require.Empty(t, unknown)
}

// TestCastToBoolSuccess checks for success scenario of
// bool config cast
func TestCastToBoolSuccess(t *testing.T) {
	unknown := castToBoolean(bool(true))
	require.True(t, unknown)
}

// TestGetBoolFailure checks for failure scenario of
// bool config retrieval
func TestGetBoolFailure(t *testing.T) {
	unknown := SrvConfig.GetBool("invalidBool")
	require.Empty(t, unknown)
}

// TestGetFloat32Failure checks for failure scenario of
// float32 config retrieval
func TestGetFloat32Failure(t *testing.T) {
	unknown := SrvConfig.GetFloat32("invalidFloat32")
	require.Empty(t, unknown)
}

// TestGetFloat64Failure checks for failure scenario of
// bool config retrieval
func TestGetFloat64Failure(t *testing.T) {
	unknown := SrvConfig.GetFloat64("invalidFloat64")
	require.Empty(t, unknown)
}

// TestFindEmptyKeyList checks for success scenario of
// retrieving all configs when list of keys sent is empty
func TestFindEmptyKeyList(t *testing.T) {
	inputMap := make(map[string]interface{})
	output := SrvConfig.find(inputMap, []string{})
	require.Equal(t, inputMap, output)
}

// TestFindStringInterface checks for success scenario of
// retrieving string interface configs
func TestFindStringInterface(t *testing.T) {
	inputMap := make(map[string]interface{})
	inputMap["https"] = map[string]interface{}{
		"port": 8080,
	}
	output := SrvConfig.find(inputMap, []string{"https", "port"})
	require.Equal(t, 8080, output)
}

// TestFindIntInterface checks for failure scenario of
// retrieving form an invalid config
func TestFindIntInterface(t *testing.T) {
	inputMap := make(map[string]interface{})
	inputMap["appId"] = map[int]interface{}{
		1: 1000,
	}
	output := SrvConfig.find(inputMap, []string{"appId", "1"})
	require.Empty(t, output)
}
