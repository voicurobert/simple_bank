package db

import (
	"database/sql"
	_ "github.com/lib/pq"
	"github.com/voicurobert/simple_bank/db/sqlc"
	"log"
	"os"
	"testing"
)

const (
	dbDriver = "postgres"
	dbSource = "postgresql://root:secret@localhost:5437/simple_bank?sslmode=disable"
)

var testQueries *db.Queries
var testDB *sql.DB

func TestMain(m *testing.M) {
	var err error
	testDB, err = sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatalln("cannot connect to db: ", err)
	}
	testQueries = db.New(testDB)

	os.Exit(m.Run())
}
