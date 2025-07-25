package config

type Config struct {
	App struct {
		name string `yaml:"name"`
	} `yaml:"app"`
	Server struct {
		Host string `yaml:"host"`
		Port string `yaml:"port"`
	} `yaml:"server"`
}
