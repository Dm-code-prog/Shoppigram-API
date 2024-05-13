package main

// Environment is the configuration struct for the application
type Environment struct {
	Postgres struct {
		DSN string `env:"POSTGRES_DSN,required"`
	}

	HTTP struct {
		Port string `env:"HTTP_PORT,required"`
	}

	Encryption struct {
		Key string `env:"ENCRYPTION_KEY,required"`
	}

	OrderNotifications struct {
		IsEnabled bool `env:"ORDER_NOTIFICATIONS_IS_ENABLED"`
		BatchSize int  `env:"ORDER_NOTIFICATIONS_BATCH_SIZE,required"`
		Timeout   int  `env:"ORDER_NOTIFICATIONS_TIMEOUT,required"`
	}

	Bot struct {
		Token string `env:"BOT_TOKEN,required"`
	}
}
