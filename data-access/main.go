package main

import (
	"bufio"

	"database/sql"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/go-sql-driver/mysql"
	"golang.org/x/term"
)

var db *sql.DB

func main() {
	cfg := mysql.NewConfig()
	cfg.User = getenvOrPrompt("DBUSER", "DB user", false)
	cfg.Passwd = getenvOrPrompt("DBPASS", "DB password", true)
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

// getenvOrPrompt reads key from env, or prompts (interactive only) and sets it for this process.
func getenvOrPrompt(key, label string, secret bool) string {
	if val, ok := os.LookupEnv(key); ok && strings.TrimSpace(val) != "" {
		return val
	}

	if !term.IsTerminal(int(os.Stdin.Fd())) {
		log.Fatalf("%s environment variable is not set and cannot prompt in non-interactive mode. Set it first (PowerShell: $Env:%s=\"value\"  |  CMD: set %s=value).", key, key, key)
	}

	reader := bufio.NewReader(os.Stdin)
	for {
		if secret {
			// Go's stdlib doesn't provide a reliable cross-platform no-echo prompt without extra handling.
			// We still read from stdin normally to keep setup simple.
			fmt.Printf("%s (will be visible while typing): ", label)
		} else {
			fmt.Printf("%s: ", label)
		}
		line, err := reader.ReadString('\n')
		if err != nil {
			log.Fatalf("failed reading %s from stdin: %v", key, err)
		}
		val := strings.TrimSpace(line)
		if val == "" {
			continue
		}
		_ = os.Setenv(key, val)
		return val
	}
}
