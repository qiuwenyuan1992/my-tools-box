package conf

const ConfFileName = "box.yaml"

var (
	IsTesting bool
)

type ToolsConf struct {
	Basic struct {
		Debug          bool   `yaml:"debug"`
		ConsoleLog     bool   `yaml:"consoleLog"`
		FileLog        bool   `yaml:"fileLog"`
		RotateTime     int    `yaml:"rotateTime"`
		MaxAge         int    `yaml:"maxAge"`
		Ip             string `yaml:"ip"`
		Port           int    `yaml:"port"`
		RestIp         string `yaml:"restIp"`
		RestPort       int    `yaml:"restPort"`
		Prometheus     bool   `yaml:"prometheus"`
		PrometheusPort int    `yaml:"prometheusPort"`
		PluginHosts    string `yaml:"pluginHosts"`
		Authentication bool   `yaml:"authentication"`
		IgnoreCase     bool   `yaml:"ignoreCase"`
	}
}

func init() {
	InitLogger()
	//InitClock()
}
