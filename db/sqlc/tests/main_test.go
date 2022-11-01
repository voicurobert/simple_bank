package db

import (
	"database/sql"
	_ "github.com/lib/pq"
	"github.com/voicurobert/simple_bank/db/sqlc"
	"github.com/voicurobert/simple_bank/util"
	"log"
	"os"
	"testing"
)

var testQueries *db.Queries
var testDB *sql.DB

func TestMain(m *testing.M) {

	config, err := util.LoadConfig("../../..")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}

	testDB, err = sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatalln("cannot connect to db: ", err)
	}
	testQueries = db.New(testDB)

	os.Exit(m.Run())
}
