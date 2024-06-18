package config

var Config *config

type Env string

const (
	EnvProduction  Env = "production"
	EnvDevelopment Env = "development"
	ENVLocal       Env = "local"
	EnvTesting     Env = "testing"
)

type config struct {
	Database struct {
		Host     string `env:"DATABASE_HOST" envDefault:"mysql"`
		Port     int    `env:"DATABASE_PORT" envDefault:"3306"`
		User     string `env:"DATABASE_USER" envDefault:"root"`
		Password string `env:"DATABASE_PASSWORD" envDefault:"admin"`
		Name     string `env:"DATABASE_NAME" envDefault:"ai-phone"`
	}
}
