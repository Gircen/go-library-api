package config

type Config struct {
	App struct {
		Name string `yaml:"name"`
	} `yaml:"app"`
	Log struct {
		Level string `yaml:"level"`
	} `yaml:"log"`
	Server struct {
		Host string `yaml:"host"`
		Port string `yaml:"port"`
	} `yaml:"server"`
}
