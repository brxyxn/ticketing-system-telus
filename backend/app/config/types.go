package config

type Sql struct {
	Host     string
	Port     string
	User     string
	Name     string
	Password string
	Sslmode  string
}

type Cache struct {
	Host     string
	Port     string
	Name     int
	Password string
}

type Config struct {
	Port string
	Sql
	Cache
}
