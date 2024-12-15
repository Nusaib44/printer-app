package config

// DefaultConfig provides default configurations
func DefaultConfig() *Config {
	return &Config{
		ServerPort:  "8080",
		DatabaseDSN: "user:password@/dbname",
	}
}
