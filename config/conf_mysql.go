package config

type Mysql struct {
	Host     string `yaml:"host"`
	POort    string `yaml:"port"`
	DB       string `yaml:"db"`
	Use      string `yaml:"use"`
	Password string `yaml:"password"`
	LogLevel string `yaml:"log_level"`
}
