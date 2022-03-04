package config

type PostgresConfig struct {
	User     string
	Password string
	Address  string
	Db       string
	Poolsize int
	Table    struct {
		ChatHistory       string
		ChatUsers    string
	}
}