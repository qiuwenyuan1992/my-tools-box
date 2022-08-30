package conf

import (
	"os"
	"reflect"
	"testing"
)

func TestEnv(t *testing.T) {
	key := "KUIPER__BASIC__CONSOLELOG"
	value := "true"

	err := os.Setenv(key, value)
	if err != nil {
		t.Error(err)
	}

	c := ToolsConf{}
	err = LoadConfig(&c)
	if err != nil {
		t.Error(err)
	}
	t.Log(c.Basic.Port)
	//if c.Basic.ConsoleLog != true {
	//	t.Errorf("env variable should set it to true")
	//}
}
func TestJsonCamelCase(t *testing.T) {
	key := "HTTPPULL__DEFAULT__BODYTYPE"
	value := "event"

	err := os.Setenv(key, value)
	if err != nil {
		t.Error(err)
	}

	const ConfigName = "redis/httppull.yaml"
	c := make(map[string]interface{})
	err = LoadConfigByName(ConfigName, &c)
	if err != nil {
		t.Error(err)
	}

	if casted, success := c["default"].(map[string]interface{}); success {
		if casted["bodyType"] != "event" {
			t.Errorf("env variable should set it to event")
		}
	} else {
		t.Errorf("returned value does not contains map under 'Basic' key")
	}
}

func TestKeyReplacement(t *testing.T) {
	m := createRandomConfigMap()
	expected := createExpectedRandomConfigMap()

	applyKey(m, "Seed")
	applyKey(m, "interval")

	if !reflect.DeepEqual(m, expected) {
		t.Errorf("key names within list should be applied \nexpected - %s\nmap      - %s", expected, m)
	}
}
func createExpectedRandomConfigMap() map[string]interface{} {
	input := createRandomConfigMap()
	def := input["default"]
	if defMap, ok := def.(map[string]interface{}); ok {
		tmp := defMap["seed"]
		delete(defMap, "seed")
		defMap["Seed"] = tmp
	}
	return input
}

func createRandomConfigMap() map[string]interface{} {
	pattern := make(map[string]interface{})
	pattern["count"] = 50
	defaultM := make(map[string]interface{})
	defaultM["interval"] = 1000
	defaultM["seed"] = 1
	defaultM["pattern"] = pattern
	defaultM["deduplicate"] = 0
	ext := make(map[string]interface{})
	ext["interval"] = 100
	dedup := make(map[string]interface{})
	dedup["interval"] = 100
	dedup["deduplicated"] = 50
	input := make(map[string]interface{})
	input["default"] = defaultM
	input["ext"] = ext
	input["dedup"] = dedup
	return input
}

