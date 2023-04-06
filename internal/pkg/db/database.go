package dbpackage

import (
	"context"
	"database/sql"
	"eden/internal/pkg/config"
	"fmt"
	"github.com/golang-migrate/migrate/v4"
	"log"

	pg "github.com/golang-migrate/migrate/v4/database/postgres"

	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
)

var (
	DB  *sql.DB
	err error
)

// SetupDB opens a database and saves the reference to `DB`.
func SetupDB() {
	var db = DB

	configuration := config.GetConfig()
	database := configuration.Database.Dbname
	username := configuration.Database.Username
	password := configuration.Database.Password
	host := configuration.Database.Host
	port := configuration.Database.Port

	db, err = sql.Open("postgres", "host="+host+" port="+port+" user="+username+" dbname="+database+" sslmode=disable password="+password)
	if err != nil {
		fmt.Println("db err: ", err)
	}

	// NOTE: do config stuff for connection here
	DB = db

	migration()
}

func migration() {
	log.Println("Starting migrations")
	driver, err := pg.WithInstance(DB, &pg.Config{})
	m, err := migrate.NewWithDatabaseInstance(
		"file://internal/pkg/db/migrations",
		"postgres", driver)
	err = m.Up()
	if err != nil {
		log.Println(err.Error())
	}
	log.Println("ending migrations")
}

func GetDB() *sql.DB {
	return DB
}

// Transaction is a wrapper function for database transactions.
func Transaction(ctx context.Context, fn func(tx *sql.Tx) (interface{}, error)) (result interface{}, err error) {
	tx, err := GetDB().BeginTx(ctx, &sql.TxOptions{})
	if err != nil {
		return nil, err
	}

	defer func() {
		if p := recover(); p != nil {
			_ = tx.Rollback()
			panic(p)
		}
		if err != nil {
			if rbErr := tx.Rollback(); rbErr != nil {
				err = fmt.Errorf("tx err: %v, rb err: %v", err, rbErr)
			}
		} else {
			err = tx.Commit()
		}
	}()

	result, err = fn(tx)

	return
}
