package config

import "fmt"

type Config struct {
	Common       Common   `mapstructure:"common"`
	DevDatabase  Database `mapstructure:"dev_database"`
	ProdDatabase Database `mapstructure:"prod_database"`
}

type Common struct {
	Host       string `mapstructure:"host"`
	Port       int    `mapstructure:"port"`
	DeployMode string `mapstructure:"deploy_mode"`
	Swagger    string `mapstructure:"swagger"`
}

type Database struct {
	Username string `mapstructure:"username"`
	Password string `mapstructure:"password"`
	Host     string `mapstructure:"host"`
	Port     string `mapstructure:"port"`
	Database string `mapstructure:"database"`
	SSLMode  string `mapstructure:"ssl_mode"`
}

func (d *Database) Dsn() string {
	return fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s", d.Host, d.Port, d.Username, d.Database, d.Password, d.SSLMode)
}
