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
	Aws struct {
		CloudFrontURI string `env:"CLOUD_FRONT_URI" envDefault:"https://d14ubbdtfe7bkh.cloudfront.net"`
		S3BucketName  string `env:"S3_BUCKET_NAME" envDefault:"ai-phone-us-west-2"`
	}
	Database struct {
		Host     string `env:"DATABASE_HOST" envDefault:"127.0.0.1"`
		Port     int    `env:"DATABASE_PORT" envDefault:"3306"`
		User     string `env:"DATABASE_USER" envDefault:"root"`
		Password string `env:"DATABASE_PASSWORD" envDefault:"admin"`
		Name     string `env:"DATABASE_NAME" envDefault:"ai-phone"`
	}
}
