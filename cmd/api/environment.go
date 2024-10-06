package main

// Environment is the configuration struct for the application
type Environment struct {
	Zap struct {
		LogLevel string `env:"LOG_LEVEL,default=INFO"`
	}

	HTTP struct {
		Port string `env:"HTTP_PORT,default=8080"`
	}

	Cache struct {
		MaxSize int64 `env:"CACHE_MAX_SIZE,default=100000000"`
	}

	NewOrderNotifications struct {
		IsEnabled  bool `env:"NEW_ORDER_NOTIFICATIONS_IS_ENABLED,default=true"`
		BatchSize  int  `env:"NEW_ORDER_NOTIFICATIONS_BATCH_SIZE,default=1"`
		TimeoutSec int  `env:"NEW_ORDER_NOTIFICATIONS_TIMEOUT,default=1"`
	}

	NewMarketplaceNotifications struct {
		IsEnabled  bool `env:"NEW_MARKETPLACE_NOTIFICATIONS_IS_ENABLED,default=true"`
		BatchSize  int  `env:"NEW_MARKETPLACE_NOTIFICATIONS_BATCH_SIZE,default=1"`
		TimeoutSec int  `env:"NEW_MARKETPLACE_NOTIFICATIONS_TIMEOUT,default=1"`
	}

	VerifiedMarketplaceNotifications struct {
		IsEnabled  bool `env:"VERIFIED_MARKETPLACE_NOTIFICATIONS_IS_ENABLED,default=true"`
		BatchSize  int  `env:"VERIFIED_MARKETPLACE_NOTIFICATIONS_BATCH_SIZE,default=1"`
		TimeoutSec int  `env:"VERIFIED_MARKETPLACE_NOTIFICATIONS_TIMEOUT,default=1"`
	}

	Postgres struct {
		DSN      string `env:"POSTGRES_DSN,required=true"`
		MinConns int    `env:"POSTGRES_MIN_CONNS,default=1"`
		MaxConns int    `env:"POSTGRES_MAX_CONNS,default=10"`
	}

	Bot struct {
		Token string `env:"BOT_TOKEN,required=true"`
		ID    int64  `env:"BOT_ID,required=true"`
		Name  string `env:"BOT_NAME,required=true"`
	}

	AWS struct {
		Cloudwatch struct {
			Namespace string `env:"AWS_CLOUD_WATCH_NAMESPACE,required=true"`
		}
	}

	DigitalOcean struct {
		Spaces struct {
			Endpoint string `env:"DIGITALOCEAN_SPACES_ENDPOINT,required=true"`
			Bucket   string `env:"DIGITALOCEAN_SPACES_BUCKET,required=true"`
			Key      string `env:"DIGITALOCEAN_SPACES_KEY,required=true"`
			Secret   string `env:"DIGITALOCEAN_SPACES_SECRET,required=true"`
		}
	}

	TelegramWebhooks struct {
		SecretToken string `env:"TELEGRAM_WEBHOOK_SECRET,required=true"`
	}

	CloudPayments struct {
		Login                  string `env:"CLOUDPAYMENTS_LOGIN,required=true"`
		Password               string `env:"CLOUDPAYMENTS_PASSWORD,required=true"`
		MaxTransactionDuration string `env:"CLOUDPAYMENTS_MAX_TRANSACTION_DURATION,default=1h"`
	}
}
