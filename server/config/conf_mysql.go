package config
type Mysql struct{
	Host string `yaml:"host"` 
	Port int`yaml:"port"` 
	User string `yaml:"user"` 
	Password string `yaml:"password"`
	Dbname string `yaml:"dbname"`
	LogLevel string `yaml:"log_level"`
}