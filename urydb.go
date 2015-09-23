package urydb

import (
	"bufio"
	"errors"
	"os"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

var (
	// ErrNoURYDB is the error thrown when URYDB is not present in the environment.
	ErrNoURYDB = errors.New("URYDB not in environment")
	// ErrNoConnFile is the error thrown when there is no urydb connection file.
	ErrNoConnFile = errors.New("couldn't find any connection file")
)

// ConnFiles is the list of possible places to search for a urydb file.
var ConnFiles = []string{
	".urydb",
	"${HOME}/.urydb",
	"/etc/urydb",
	"/usr/local/etc/urydb",
}

func getConnString() (connString string, err error) {
	connString, err = getConnStringEnv()
	if err != nil {
		connString, err = getConnStringFile()
	}
	return
}

func getConnStringEnv() (connString string, err error) {
	connString, err = os.Getenv("URYDB"), nil
	if connString == "" {
		err = ErrNoURYDB
	}
	return
}

func getConnStringFile() (connString string, err error) {
	for _, rawPath := range ConnFiles {
		path := os.ExpandEnv(rawPath)
		file, ferr := os.Open(path)
		if ferr != nil {
			connString = ""
			continue
		}

		bufrd := bufio.NewReader(file)
		connString, ferr = bufrd.ReadString('\n')

		if ferr != nil {
			connString = ""
			continue
		}

		return
	}

	if connString == "" {
		err = ErrNoConnFile
	}
	return
}

// GetDB returns a pointer to an open URY database connection
func GetDB() (*sqlx.DB, error) {
	connString, err := getConnString()
	if err != nil {
		return nil, err
	}

	return sqlx.Open("postgres", connString)
}
