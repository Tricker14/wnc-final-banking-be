package database

import (
	"fmt"
	"log"
	"os"
	"path"
	"path/filepath"

	_ "github.com/go-sql-driver/mysql"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/jmoiron/sqlx"
	_ "github.com/joho/godotenv/autoload"
)

var (
	dbname   = os.Getenv("DB_DATABASE")
	password = os.Getenv("DB_PASSWORD")
	username = os.Getenv("DB_USERNAME")
	port     = os.Getenv("DB_PORT")
	host     = os.Getenv("DB_HOST")
)

type Db *sqlx.DB

func Open() *sqlx.DB {
	// Opening a driver typically will not attempt to connect to the database.
	db, err := sqlx.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?multiStatements=true&parseTime=true", username, password, host, port, dbname))
	driver, err := mysql.WithInstance(db.DB, &mysql.Config{})
	if err != nil {
		// This will not be a connection error, but a DSN parse error or
		// another initialization error.
		log.Fatal(err)
	}
	db.SetConnMaxLifetime(0)
	db.SetMaxIdleConns(50)
	db.SetMaxOpenConns(50)

	p, err := filepath.Abs("F://Documents//Code//Golang//internet-banking")
	p = filepath.ToSlash(p)
	p = path.Join(p, "migrations")

	m, err := migrate.NewWithDatabaseInstance(
		fmt.Sprintf("file://%s", p),
		"mysql", driver)

	if err != nil {
		log.Fatal("err create migration instance ", err)
	}

	err = m.Up()

	return db
}
