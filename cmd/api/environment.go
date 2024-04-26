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
}
