package db

import (
	"database/sql"
	"fmt"
	"io"
	"os"
	"strings"
	"sub/utils/log"
)

var STATEMENTS = make(map[string]*sql.Stmt)

func LoadStatements(path string) error {
	if db == nil {
		log.Fatal("Database not initialized")
	}

	if _, err := os.Stat(path); os.IsNotExist(err) {
		return fmt.Errorf("SQL file does not exist: %v", err)
	}

	file, err := os.Open(path)
	if err != nil {
		return fmt.Errorf("error opening SQL file: %v", err)
	}
	defer file.Close()

	data, err := io.ReadAll(file)
	if err != nil {
		return fmt.Errorf("error reading SQL file: %v", err)
	}

	stmts := strings.Split(string(data), "\n-- ")
	if strings.HasPrefix(stmts[0], "-- ") {
		stmts[0] = stmts[0][3:]
	} else {
		stmts = stmts[1:]
	}

	for _, stmt := range stmts {
		tmp := strings.SplitN(stmt, "\n", 2)
		name := strings.TrimSpace(tmp[0])
		statementStr := strings.TrimSpace(tmp[1])
		statement, err := db.Prepare(statementStr)
		if err != nil {
			log.Error("Error preparing statement", "name", name, "err", err)
			continue
		}
		STATEMENTS[name] = statement
	}

	log.Info("Statements loaded successfully")
	return nil
}

func GetStatement(name string) *sql.Stmt {
	stmt, ok := STATEMENTS[name]
	if !ok {
		log.Fatal("statement not found", "stmt", name)
	}
	return stmt
}

var (
	FetchFlagsStmt           *sql.Stmt
	UpdateExpiredFlagsStmt   *sql.Stmt
	UpdateFlagStmt           *sql.Stmt
	InsertFlagStmt           *sql.Stmt
	FetchAllExploitDataStmt  *sql.Stmt
	FetchExploitDataStmt     *sql.Stmt
	FetchAllTimelineDataStmt *sql.Stmt
	FetchTimelineDataStmt    *sql.Stmt
	FetchAllTeamsDataStmt    *sql.Stmt
	FetchTeamsDataStmt       *sql.Stmt
)

func InitStatements(path string) {
	err := LoadStatements(path)
	if err != nil {
		log.Fatal("Error loading statements", "err", err)
	}

	FetchFlagsStmt = GetStatement("FetchFlags")
	UpdateExpiredFlagsStmt = GetStatement("UpdateExpiredFlags")
	UpdateFlagStmt = GetStatement("UpdateFlag")
	InsertFlagStmt = GetStatement("InsertFlag")
	FetchAllExploitDataStmt = GetStatement("FetchAllExploitData")
	FetchExploitDataStmt = GetStatement("FetchExploitData")
	FetchAllTimelineDataStmt = GetStatement("FetchAllTimelineData")
	FetchTimelineDataStmt = GetStatement("FetchTimelineData")
	FetchAllTeamsDataStmt = GetStatement("FetchAllTeamsData")
	FetchTeamsDataStmt = GetStatement("FetchTeamsData")
}
