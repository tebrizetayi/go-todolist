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

/*host     = "localhost"
port     = 5555
user     = "postgres"
password = "postgres"
dbname   = "testdb"*/

func NewConfig() *Config {
	return Cnf
}
