package main

import (
	"flag"
	"os"
)

const version = "1.0"

// Config structure contains all the important and reusable data we will need for the API
type config struct {
	port int
	env  string
	db   struct {
		dsn          string
		maxOpenConns int
		maxIdleConns int
		maxIdleTime  string
	}

	limiter struct {
		rps     float64
		burst   int
		enabled bool
	}
}

type application struct {
	config config
}

func main() {
	// Create
	var cfg config
	flag.IntVar(&cfg.port, "port", 7000, "API server port")
	flag.StringVar(&cfg.env, "env", "development", "Environment (development|staging|production)")

	// Get the database connection string, aka data source name (DSN)
	flag.StringVar(&cfg.db.dsn, "db-dsn", os.Getenv("DSN2"), "PostgreSQL DSN")

	// Set up restrictions for the database connections
	flag.IntVar(&cfg.db.maxOpenConns, "db-max-open-conns", 25, "PostgreSQL max open connections")
	flag.IntVar(&cfg.db.maxIdleConns, "db-max-idle-conns", 25, "PostgreSQL max idle connections")
	flag.StringVar(&cfg.db.maxIdleTime, "db-max-idle-time", "15m", "PostgreSQL max idle time")

}
