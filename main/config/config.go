package config

type Config struct {
	App struct {
		Name string `yaml:"name"`
	} `yaml:"app"`
	Log struct {
		Level         string `yaml:"level"`
		FileMaxSizeMB int    `yaml:"fileMaxSizeMB"`
		//if count log files > countFilesArchive then program make files to zip archive {nawDate.zip}
		CountFilesArchive int `yaml:"countFilesArchive"`
	} `yaml:"log"`
	Server struct {
		Name string `yaml:"name"`
		Host string `yaml:"host"`
		Port string `yaml:"port"`
	} `yaml:"server"`
	HealthServer struct {
		Host string `yaml:"host"`
		Port string `yaml:"port"`
	} `yaml:"healthServer"`
	SettingsServer struct {
		Host string `yaml:"host"`
		Port string `yaml:"port"`
	} `yaml:"settingsServer"`
	Services []struct {
		Name string `yaml:"name"`
		Host string `yaml:"host"`
		Port string `yaml:"port"`
	} `yaml:"services"`
}
