package db

import (
	"database/sql"
	"io"
	"os"
	"sub/utils/log"

	_ "github.com/mattn/go-sqlite3"
)

var db *sql.DB

func InitDB(path string) {
	var err error
	db, err = sql.Open("sqlite3", path)
	if err != nil {
		log.Fatal("Error opening database", "err", err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal("Error connecting to the database", "err", err)
	}

	db.SetMaxOpenConns(1)

	log.Info("Successfully connected to the database")
}

func CloseDB() {
	if db == nil {
		log.Fatal("Database not initialized")
	}

	err := db.Close()
	if err != nil {
		log.Fatal("Error closing database", "err", err)
	}

	log.Info("Database connection closed")
}

func ExecSQLFile(path string) {
	if db == nil {
		log.Fatal("Database not initialized")
	}

	if _, err := os.Stat(path); os.IsNotExist(err) {
		log.Fatal("SQL file does not exist", "err", err)
	}

	file, err := os.Open(path)
	if err != nil {
		log.Fatal("Error opening SQL file", "err", err)
	}
	defer file.Close()

	data, err := io.ReadAll(file)
	if err != nil {
		log.Fatal("Error reading SQL file", "err", err)
	}

	_, err = db.Exec(string(data))
	if err != nil {
		log.Fatal("Error executing SQL", "err", err)
	}

	log.Info("SQL executed successfully")
}

func CleanDB() {
	if db == nil {
		log.Fatal("Database not initialized")
	}

	rows, err := db.Query("SELECT name FROM sqlite_master WHERE type='table'")
	if err != nil {
		log.Error("Error querying table names", "err", err)
		return
	}
	defer rows.Close()

	tables := make([]string, 0)
	for rows.Next() {
		var table string
		err = rows.Scan(&table)
		if err != nil {
			log.Error("Error scanning table", "err", err)
			return
		}
		tables = append(tables, table)
	}

	for _, table := range tables {
		_, err := db.Exec("DELETE FROM " + table)
		if err != nil {
			log.Error("Error cleaning", "table", table, "err", err)
			return
		}
	}

	log.Info("Database cleaned successfully")
}

func DropTables() {
	if db == nil {
		log.Fatal("Database not initialized")
	}

	rows, err := db.Query("SELECT name FROM sqlite_master WHERE type='table'")
	if err != nil {
		log.Error("Error querying table names", "err", err)
		return
	}
	defer rows.Close()

	tables := make([]string, 0)
	for rows.Next() {
		var table string
		err = rows.Scan(&table)
		if err != nil {
			log.Error("Error scanning table", "err", err)
			return
		}
		tables = append(tables, table)
	}

	for _, table := range tables {
		_, err := db.Exec("DROP TABLE " + table)
		if err != nil {
			log.Error("Error dropping", "table", table, "err", err)
			return
		}
	}

	log.Info("Tables Dropped successfully")
}
