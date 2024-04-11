package main

// Environment is the configuration struct for the application
type Environment struct {
	Postgres struct {
		DSN string `env:"POSTGRES_DSN"`
	}

	HTTP struct {
		Port string `env:"HTTP_PORT"`
	}
}
