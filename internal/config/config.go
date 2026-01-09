package config

type Config struct {
	Server   ServerConfig
	Database DbConfig
	Security SecurityConfig
}

type ServerConfig struct {
	Name string
	Env  string
	Port string
}

type DbConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	Name     string
	SSLMode  string
}

type SecurityConfig struct {
	JWTSecret string
}
