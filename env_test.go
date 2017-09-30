package env

import (
	"reflect"
	"testing"
)

func TestGet(t *testing.T) {
	SetPAL("testdata/config.yaml")

	if Config.Get("SITE_NAME") != "Micro" {
		t.Fatal("error")
	}
}

func TestGetString(t *testing.T) {
	SetPAL("testdata/config.yaml")

	if GetString("DB_PORT") != "3306" {
		t.Fatal("value is not a string")
	}
}

func TestGetInt(t *testing.T) {
	SetPAL("testdata/config.yaml")

	if GetInt("DB_PORT") != 3306 {
		t.Fatal("value is not an integer")
	}
}
func TestGetBool(t *testing.T) {
	SetPAL("testdata/config.yaml")

	v := reflect.TypeOf(GetBool("SERVER_HTTPS"))

	if v.Kind() != reflect.Bool {
		t.Fatal("value is not an boolean")
	}
}

func TestMode(t *testing.T) {
	SetPAL("testdata/config.yaml")

	if Config.Mode != "development" {
		t.Fatalf("env mode is not development, is: %v", Config.Mode)
	}
}
func TestNew(t *testing.T) {
	DefaultPath = "testdata/config.yaml"
	New()

	if GetInt("DB_PORT") != 3306 {
		t.Fatal("the New func work as expected")
	}
}
func TestGetSpecificEnvironmentConfigs(t *testing.T) {
	SetPAL("testdata/config.yaml")
	GetEnv(Development)
	// TODO
}
