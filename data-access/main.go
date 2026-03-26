package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

var db *sql.DB

func main() {
	// Optional convenience: load env vars from a local .env file (ignored if missing).
	_ = godotenv.Load()

	cfg := mysql.NewConfig()
	cfg.User = mustGetenvWithFallbacks("DBUSER")
	cfg.Passwd = mustGetenvWithFallbacks("DBPASS")
	cfg.Net = "tcp"
	cfg.Addr = "127.0.0.1:3306"
	cfg.DBName = "new_schema"
   
	// Get a database handle.
	var err error
	db, err = sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		log.Fatal(err)
	}

	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected!")
}

// mustGetenvWithFallbacks returns the env var or exits with a clear setup hint.
func mustGetenvWithFallbacks(key string) string {
	// Common fallbacks (e.g., when using MySQL/MariaDB Docker images).
	fallbackKeys := map[string][]string{
		"DBUSER": {"MYSQL_USER", "MYSQL_USERNAME"},
		"DBPASS": {"MYSQL_PASSWORD", "MYSQL_PWD"},
	}

	if val, ok := os.LookupEnv(key); ok && strings.TrimSpace(val) != "" {
		return val
	}
	for _, fk := range fallbackKeys[key] {
		if val, ok := os.LookupEnv(fk); ok && strings.TrimSpace(val) != "" {
			return val
		}
	}

	log.Fatalf("%s environment variable is not set. Set it first (PowerShell: $Env:%s=\"value\"  |  CMD: set %s=value). You can also put it in a local .env file.", key, key, key)
	return ""
}
