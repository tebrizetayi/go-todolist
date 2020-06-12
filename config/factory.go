package config

var Cnf = &Config{
	Database: DatabaseConfig{
		Type:         "postgres",
		Host:         "localhost",
		Port:         5432,
		User:         "postgres",
		Password:     "",
		DatabaseName: "go_todolist_server",
		MaxIdleConns: 5,
		MaxOpenConns: 5,
	},
	IsDevelopment: true,
}

var testCnf = &Config{
	Database: DatabaseConfig{
		Type:         "postgres",
		Host:         "localhost",
		Port:         5432,
		User:         "postgres",
		Password:     "",
		DatabaseName: "go_todolist_server_test",
		MaxIdleConns: 5,
		MaxOpenConns: 5,
	},
	IsDevelopment: true,
}

func NewConfig() *Config {
	return Cnf
}

func NewTestConfig() *Config {
	return testCnf
}
