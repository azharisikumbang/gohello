package app

type Config struct {
	App AppConfig
	DB  DBConfig
}

type AppConfig struct {
	Port string
}

type DBConfig struct {
	Host     string
	Username string
	Password string
	Name     string
	Port     string
	Driver   string
}
