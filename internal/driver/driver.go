package driver

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/jackc/pgconn"
	_ "github.com/jackc/pgx/v4"
	_ "github.com/jackc/pgx/v4/stdlib"
)

type DB struct {
	SQL *sql.DB
}

var dbConn = new(DB)

const maxOpenDbConn = 5
const maxIdleDbConn = 5
const maxDbLifeTime = 5 * time.Minute

func ConnectPostgres(dsn string) (*DB, error) {
	db, err := sql.Open("pgx", dsn)
	if err != nil {
		return nil, err
	}

	db.SetMaxOpenConns(maxOpenDbConn)
	db.SetMaxIdleConns(maxIdleDbConn)
	db.SetConnMaxLifetime(maxDbLifeTime)

	err = testDb(db)
	if err != nil {
		return nil, err
	}

	dbConn.SQL = db

	return dbConn, nil
}

func testDb(db *sql.DB) error {
	err := db.Ping()
	if err != nil {
		fmt.Println("Error:", err)
		return err
	}

	fmt.Println("*** Ping Successfully ***")

	return nil
}
