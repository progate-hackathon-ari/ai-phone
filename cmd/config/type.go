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
	App struct {
		Addr        string `env:"SERVER_ADDR" envDefault:":8080"`
		Env         Env    `env:"ENV"`
		AllowOrigin string `env:"ALLOW_ORIGINS"`
	}

	Database struct {
		Host     string `env:"DATABASE_HOST"`
		Port     int    `env:"DATABASE_PORT"`
		User     string `env:"DATABASE_USER"`
		Password string `env:"DATABASE_PASSWORD"`
		Name     string `env:"DATABASE_NAME"`
	}

	Otel struct {
		Addr      string `env:"OTEL_ADDR"`
		ProjectID string `env:"OTEL_PROJECT_ID"`
		IsUse     bool   `env:"OTEL_USE"`
	}

	Cloudflare struct {
		AccountID       string `env:"CLOUDFLARE_ACCOUNT_ID"`
		Endpoint        string `env:"CLOUDFLARE_ENDPOINT"`
		AccessKeyID     string `env:"CLOUDFLARE_ACCESS_KEY_ID"`
		AccessKeySecret string `env:"CLOUDFLARE_ACCESS_KEY_SECRET"`
		BucketName      string `env:"CLOUDFLARE_BUCKET_NAME"`
		Region          string `env:"CLOUDFLARE_REGION"`
	}

	Google struct {
		ClientID     string `env:"GOOGLE_CLIENT_ID"`
		ClientSecret string `env:"GOOGLE_CLIENT_SECRET"`
		RedirectURI  string `env:"GOOGLE_REDIRECT_URI"`
	}

	Azure struct {
		TenantID                 string `env:"AZURE_TENANT_ID"`
		ClientID                 string `env:"AZURE_CLIENT_ID"`
		ClientSecret             string `env:"AZURE_CLIENT_SECRET"`
		BlobServiceURL           string `env:"AZURE_BLOB_SERVICE_URL"`
		BlobServiceContainerName string `env:"AZURE_BLOB_SERVICE_CONTAINER_NAME"`
	}
}
