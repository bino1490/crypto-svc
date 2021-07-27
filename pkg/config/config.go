package config

import (
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"

	"gopkg.in/yaml.v2"
)

var (
	rootDirectory = "/"
	delimiter     = "."

	// SrvConfig service configuration
	SrvConfig *config
)

type config struct {
	value map[string]interface{}
}

func (c *config) GetVariable(key string) interface{} {
	val := os.Getenv(key)
	if len(val) == 0 {
		return c.find(c.value, strings.Split(key, delimiter))
	}
	if variable, _ := base64.StdEncoding.DecodeString(val); len(variable) != 0 {
		return string(variable)
	}
	return nil
}

// GetString returns the value of the key as a string type.
func (c *config) GetString(key string) string {
	return castToString(c.GetVariable(key))
}

// GetBool returns the value of the key as a boolean type.
func (c *config) GetBool(key string) bool {
	return castToBoolean(c.GetVariable(key))
}

// GetInt returns the value of the key as a int type.
func (c *config) GetInt(key string) int {
	return castToInt(c.GetVariable(key))
}

// GetInt32 returns the value of the key as a int32 type.
func (c *config) GetInt32(key string) int32 {
	return castToInt32(c.GetVariable(key))
}

// GetInt64 returns the value of the key as a int64 type.
func (c *config) GetInt64(key string) int64 {
	return castToInt64(c.GetVariable(key))
}

// GetFloat32 returns the value of the key as a float32 type.
func (c *config) GetFloat32(key string) float32 {
	return castToFloat32(c.GetVariable(key))
}

// GetFloat64 returns the value of the key as a float64 type.
func (c *config) GetFloat64(key string) float64 {
	return castToFloat64(c.GetVariable(key))
}

func initRootDir() {
	_, b, _, _ := runtime.Caller(0)
	rootDirectory = filepath.Join(filepath.Dir(b), "..", "..")
}

// init loads the configuration YAML file to initialize the
// key/value pairs.
func init() {
	initRootDir()
	env := os.Getenv("APP_ENV_PROFILE")
	if len(env) == 0 {
		env = "local"
	}
	curDir, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	if curDir == "/" {
		rootDirectory = curDir
	}
	configFilePath := filepath.Join(rootDirectory, "conf", "config-"+env+".yml")
	file, err := ioutil.ReadFile(filepath.Clean(configFilePath))
	if err != nil {
		panic(err)
	}
	value := make(map[string]interface{})
	if err := yaml.Unmarshal(file, &value); err != nil {
		panic(err)
	}

	SrvConfig = &config{value: value}
}

// find recursively searches for the key and its corresponding value
//  in the nested map.
func (c *config) find(obj map[string]interface{}, key []string) interface{} {
	// Key without delimiter.
	if len(key) == 0 {
		return obj
	}
	next, ok := obj[key[0]]
	if ok {
		// Not nested key lookup
		if len(key) == 1 {
			return next
		}
		// Nested key lookup
		switch next.(type) {
		case map[interface{}]interface{}:
			m := make(map[string]interface{})
			for k, val := range next.(map[interface{}]interface{}) {
				m[fmt.Sprintf("%v", k)] = val
			}
			return c.find(m, key[1:])
		case map[string]interface{}:
			return c.find(next.(map[string]interface{}), key[1:])
		default:
			return nil
		}
	}
	return nil
}

// castToString casts the input interface to an string type.
func castToString(v interface{}) string {
	if v != nil {
		return fmt.Sprintf("%v", v)
	}
	return ""
}

// castToBoolean casts the input interface to an boolean type.
func castToBoolean(v interface{}) bool {
	val := false
	switch v.(type) {
	case string:
		val, _ = strconv.ParseBool(v.(string))
	default:
		val, _ = strconv.ParseBool(fmt.Sprintf("%t", v))
	}
	return val
}

// castToInt casts the input interface to an int type.
func castToInt(v interface{}) int {
	var val int64
	switch v.(type) {
	case string:
		val, _ = strconv.ParseInt(v.(string), 0, 0)
	default:
		val, _ = strconv.ParseInt(fmt.Sprintf("%d", v), 0, 0)
	}
	return int(val)
}

// castToInt32 casts the input interface to an int32 type.
func castToInt32(v interface{}) int32 {
	var val int64
	switch v.(type) {
	case string:
		val, _ = strconv.ParseInt(v.(string), 0, 0)
	default:
		val, _ = strconv.ParseInt(fmt.Sprintf("%d", v), 0, 0)
	}
	return int32(val)
}

// castToInt64 casts the input interface to an int64 type.
func castToInt64(v interface{}) int64 {
	var val int64
	switch v.(type) {
	case string:
		val, _ = strconv.ParseInt(v.(string), 0, 0)
	default:
		val, _ = strconv.ParseInt(fmt.Sprintf("%d", v), 0, 0)
	}
	return val
}

// castToFloat32 casts the input interface to an float32 type.
func castToFloat32(v interface{}) float32 {
	var val float64
	switch v.(type) {
	case string:
		val, _ = strconv.ParseFloat(v.(string), 32)
	default:
		val, _ = strconv.ParseFloat(fmt.Sprintf("%f", v), 32)
	}
	return float32(val)
}

// castToFloat64 casts the input interface to an float64 type.
func castToFloat64(v interface{}) float64 {
	var val float64
	switch v.(type) {
	case string:
		val, _ = strconv.ParseFloat(v.(string), 32)
	default:
		val, _ = strconv.ParseFloat(fmt.Sprintf("%f", v), 32)
	}
	return val
}
