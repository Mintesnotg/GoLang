package main

import (
    "database/sql"
    "fmt"
    "log"

    "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func main() {
    // Capture connection properties.
    cfg := mysql.NewConfig()
    cfg.User = "root"
    cfg.Passwd = "Minte@1997##"
    cfg.Net = "tcp"
    cfg.Addr = "127.0.0.1:3306"
    cfg.DBName = "new_schema"

    // Get a database handle.
    var err error
    db, err = sql.Open("mysql", cfg.FormatDSN())
    if err != nil {
        log.Fatal(err)
    }

    pingErr := db.Ping()
    if pingErr != nil {
        log.Fatal(pingErr)
    }
    fmt.Println("Connected!")
}
// mustGetenv returns the env value for key, exiting with a clear hint if missing.
