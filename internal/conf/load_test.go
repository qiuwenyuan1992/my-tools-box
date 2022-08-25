package conf

import (
	"os"
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
