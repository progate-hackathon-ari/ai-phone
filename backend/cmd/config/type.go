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
		CloudFrontURI string `json:"CLOUD_FRONT_URI" envDefault:"https://dh93p4xay7grb.cloudfront.net"`
		S3BucketName  string `json:"S3_BUCKET_NAME" envDefault:"ai-phone-img"`
	}
	Database struct {
		Host     string `env:"DATABASE_HOST" envDefault:"127.0.0.1"`
		Port     int    `env:"DATABASE_PORT" envDefault:"3306"`
		User     string `env:"DATABASE_USER" envDefault:"root"`
		Password string `env:"DATABASE_PASSWORD" envDefault:"admin"`
		Name     string `env:"DATABASE_NAME" envDefault:"ai-phone"`
	}
}
