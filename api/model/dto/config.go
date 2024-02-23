package dto

type Config struct {
	ABCD string `mapstructure:"ABC"`
	Database  Database `mapstructure:",squash"`
	Redis Redis `mapstructure:",squash"`
	RabbitMQ RabbitMQ `mapstructure:",squash"`
	SMTP SMTP `mapstructure:",squash"`
	Twillio Twillio `mapstructure:",squash"`
}

type Database struct {
	Username string `mapstructure:"DATABASE_USERNAME"`
	Password string `mapstructure:"DATABASE_PASSWORD"`
	Port     string `mapstructure:"DATABASE_PORT"`
	Name     string `mapstructure:"DATABASE_NAME"`
	SSLMode  string `mapstructure:"DATABASE_SSLMODE"`
}

type Redis struct {
	Port     string `mapstructure:"REDIS_PORT"`
	Password string `mapstructure:"REDIS_PASSWORD"`
	DB       int    `mapstructure:"REDIS_DB"`
}

type RabbitMQ struct {
	Username string `mapstructure:"RABBITMQ_USERNAME"`
	Password string `mapstructure:"RABBITMQ_PASSWORD"`
	Port     string `mapstructure:"RABBITMQ_PORT"`
}

type SMTP struct {
	EmailFrom string `mapstructure:"SMTP_EMAIL_FROM"`
	EmailPassword string `mapstructure:"SMTP_EMAIL_PASSWORD"`
	Host string `mapstructure:"SMTP_HOST"`
	Port string `mapstructure:"SMTP_PORT"`
}

type Twillio struct {
	AccountSID string `mapstructure:"TWILIO_ACCOUNT_SID"`
	AuthToken string `mapstructure:"TWILIO_AUTH_TOKEN"`
}

type JWTSecret struct {
	SecretKey string `json:"secretkey"`
}
