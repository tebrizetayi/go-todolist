package config

// DatabaseConfig stores database connection options
type DatabaseConfig struct {
	Type         string
	Host         string
	Port         int
	User         string
	Password     string
	DatabaseName string
	MaxIdleConns int
	MaxOpenConns int
}

// Config stores all configuration options
type Config struct {
	Database      DatabaseConfig
	IsDevelopment bool
}
