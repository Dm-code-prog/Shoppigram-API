package main

// Environment is the configuration struct for the application
type Environment struct {
	Zap struct {
		LogLevel string `env:"LOG_LEVEL,required"`
	}

	Postgres struct {
		DSN string `env:"POSTGRES_DSN,required"`
	}

	HTTP struct {
		Port string `env:"HTTP_PORT,required"`
	}

	Encryption struct {
		Key string `env:"ENCRYPTION_KEY,required"`
	}

	NewOrderNotifications struct {
		IsEnabled bool `env:"NEW_ORDER_NOTIFICATIONS_IS_ENABLED"`
		BatchSize int  `env:"NEW_ORDER_NOTIFICATIONS_BATCH_SIZE,required"`
		Timeout   int  `env:"NEW_ORDER_NOTIFICATIONS_TIMEOUT,required"`
	}

	NewMarketplaceNotifications struct {
		IsEnabled bool `env:"NEW_MARKETPLACE_NOTIFICATIONS_IS_ENABLED"`
		BatchSize int  `env:"NEW_MARKETPLACE_NOTIFICATIONS_BATCH_SIZE,required"`
		Timeout   int  `env:"NEW_MARKETPLACE_NOTIFICATIONS_TIMEOUT,required"`
	}

	VerifiedMarketplaceNotifications struct {
		IsEnabled bool `env:"VERIFIED_MARKETPLACE_NOTIFICATIONS_IS_ENABLED"`
		BatchSize int  `env:"VERIFIED_MARKETPLACE_NOTIFICATIONS_BATCH_SIZE,required"`
		Timeout   int  `env:"VERIFIED_MARKETPLACE_NOTIFICATIONS_TIMEOUT,required"`
	}

	Bot struct {
		Token string `env:"BOT_TOKEN,required"`
		ID    int64  `env:"BOT_ID,required"`
		Name  string `env:"BOT_NAME,required"`
	}

	DigitalOcean struct {
		Spaces struct {
			Endpoint string `env:"DIGITALOCEAN_SPACES_ENDPOINT,required"`
			Bucket   string `env:"DIGITALOCEAN_SPACES_BUCKET,required"`
			Key      string `env:"DIGITALOCEAN_SPACES_KEY,required"`
			Secret   string `env:"DIGITALOCEAN_SPACES_SECRET,required"`
		}
	}

	TelegramWebhooks struct {
		SecretToken string `env:"TELEGRAM_WEBHOOK_SECRET,required"`
	}
}
