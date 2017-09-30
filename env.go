package env

import (
	"io/ioutil"
	"log"
	"strconv"

	yaml "gopkg.in/yaml.v2"
)

const (
	Development = "development"
	Production  = "production"
	Test        = "test"
)

// DefaultMode - this to use outside of package when use SetPAL
var DefaultMode = Development

// DefaultPath ...
var DefaultPath = "config.yaml"

// Config to store all configs
var Config *Configs

// Configs ...
type Configs struct {
	Path    string
	Mode    string
	Configs map[string]map[string]interface{}
}

// Env ...
type Env struct {
	Env string
}

// Get return eenvironment value
func (e *Env) Get() string {
	return e.Env
}

// New return Config instance
func New() {
	Config = &Configs{
		Path:    DefaultPath, // this is the default config path
		Mode:    DefaultMode,
		Configs: make(map[string]map[string]interface{}, 3),
	}
	Config.Load()
}

// SetPAL - Set Path And Load a custom file path in the same time
func SetPAL(path string) {
	Config = &Configs{
		Path:    path, // this is the default config path
		Mode:    DefaultMode,
		Configs: make(map[string]map[string]interface{}, 3),
	}
	Config.Load()
}

// Load config file data
func (c *Configs) Load() {
	// load toml file
	content, err := ioutil.ReadFile(c.Path)
	// check for error when open the file
	if err != nil {
		log.Fatal(err)
	}

	var env Env
	var data map[string]map[string]interface{}

	// get environment value
	yaml.Unmarshal(content, &env)
	// set environment to the struct
	c.Mode = env.Env

	// Unmarshal yaml data
	yaml.Unmarshal(content, &data)

	// set the three envirement mode's
	Config.Configs[Development] = data["development"]
	Config.Configs[Production] = data["production"]
	Config.Configs[Test] = data["test"]
}

// SetPath to set a custom path for config file
func (c *Configs) SetPath(path string) {
	c.Path = path
}

// Get ...
func (c *Configs) Get(key string) interface{} {
	return c.Configs[c.Mode][key]
}

// GetEnv return a specific environment values
func (c *Configs) GetEnv(env string) map[string]interface{} {
	return Config.Configs[env]
}

// GetString ...
func (c *Configs) GetString(key string) string {
	switch c.Configs[c.Mode][key].(type) {
	case int:
		return strconv.Itoa(c.Configs[c.Mode][key].(int))
	default:
		return c.Configs[c.Mode][key].(string)
	}
}

// GetInt ...
func (c *Configs) GetInt(key string) int {
	switch c.Configs[c.Mode][key].(type) {
	case string:
		integer, err := strconv.Atoi(c.Configs[c.Mode][key].(string))
		if err != nil {
			panic(err)
		}

		return integer
	default:
		return c.Configs[c.Mode][key].(int)
	}
}

// Get Config
func Get(key string) interface{} {
	return Config.Get(key)
}

// GetEnv return a specific environment values
func GetEnv(env string) map[string]interface{} {
	return Config.GetEnv(env)
}

// GetString retun a string config
func GetString(key string) string {
	return Config.GetString(key)
}

// GetInt return an integer config
func GetInt(key string) int {
	return Config.GetInt(key)
}
