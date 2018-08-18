package util

type Config struct {
	DB *DBConfig
}

type DBConfig struct {
	Dialect  string
	Username string
	Password string
	DBName   string
	Charset  string
}

func GetConfig() *Config {
	return &Config{
		&DBConfig{
			Dialect:  "mysql",
			Username: "root",
			Password: "12345678",
			DBName:   "user",
			Charset:  "utf8",
		},
	}
}
