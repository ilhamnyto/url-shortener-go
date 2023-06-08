package database

import (
	"database/sql"
	"fmt"

	"github.com/ilhamnyto/url-shortener-go/config"
	_ "github.com/lib/pq"
)

func ConnectDB() (*sql.DB) {
	var (
		db_host = config.GetString(config.POSTGRES_HOST)
		db_port = config.GetString(config.POSTGRES_PORT)
		db_user = config.GetString(config.POSTGRES_USER)
		db_pass = config.GetString(config.POSTGRES_PASSWORD)
		db_name = config.GetString(config.POSTGRES_DB)
	)

	dsn := fmt.Sprintf(
		`host=%s port=%s user=%s password=%s dbname=%s sslmode=disable`,
		db_host, db_port, db_user, db_pass, db_name,
	)

	dbsql, err := sql.Open("postgres", dsn)

	if err != nil {
		panic(err)
	}

	if err := dbsql.Ping(); err != nil {
		panic(err)
	}

	return dbsql
}

func MigrateDB (db *sql.DB) {
	fmt.Println("Start Migrating DB.")
	var (
		queryCreateUserTable = `
		CREATE TABLE IF NOT EXISTS users (
            id serial primary key,
            username varchar(100) NOT NULL,
            email varchar(100) NOT NULL,
            password varchar(100) NOT NULL,
            salt varchar(100) NOT NULL,
            created_at timestamp,
            updated_at timestamp
        )
		`

		queryCreateURLTable = `
			CREATE TABLE IF NOT EXISTS urls (
				id serial primary key,
				ulid varchar(100) NOT NULL,
				user_id int NOT NULL,
				long_url text NOT NULL,
				short_url text,
				visits int default 0,
				created_at timestamp,
				CONSTRAINT fk_urls
				FOREIGN KEY(user_id)
				REFERENCES users(id)
			)
		`
	)

	stmt, err := db.Prepare(queryCreateUserTable)

	if err != nil {
		panic(err)
	}

	if _, err := stmt.Exec(); err != nil {
		panic(err)
	}
	
	stmt, err = db.Prepare(queryCreateURLTable)
	
	if err != nil {
		panic(err)
	}

	if _, err := stmt.Exec(); err != nil {
		panic(err)
	}

	fmt.Println("Migrate Success.")
}