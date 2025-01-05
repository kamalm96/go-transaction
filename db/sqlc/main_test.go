package db

import (
	"database/sql"
	"log"
	"os"
	"testing"

	"example.com/m/v2/util"
	_ "github.com/lib/pq"
)

var testQueries *Queries
var testDB *sql.DB

func TestMain(m *testing.M) {
	var err error

	config, err := util.LoadConfig("../..")
	if err != nil {
		log.Fatal("Cannot load test config file:", err)
	}

	testDB, err = sql.Open(config.DBDriver, config.DBSource)

	if err != nil {
		log.Fatal("Cannot connect to DB:", err)
	}
	testQueries = New(testDB)

	os.Exit(m.Run())
}
