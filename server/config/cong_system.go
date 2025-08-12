package config
type System struct{
	Env string `yaml:"env"`
	Port int `yaml:"port"`
	Host string `yaml:"host"`
}