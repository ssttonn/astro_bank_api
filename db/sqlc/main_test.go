package astrobank_db

import (
	"database/sql"
	"log"
	"os"
	"testing"

	_ "github.com/lib/pq"
)

const (
	dbDriver = "postgres"
	dbSource = "postgresql://astrotify-admin:sample@localhost:5432/astrobank_db?sslmode=disable"
)

var testQueries *Queries
var testDB *sql.DB

func TestMain(m *testing.M) {
	conn, error := sql.Open(dbDriver, dbSource)
	if error != nil {
		log.Fatal("Can't connect to db: ", error)
	}

	testDB = conn

	testQueries = New(testDB)

	os.Exit(m.Run())
}
