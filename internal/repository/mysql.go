package repository

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func migrate() error {

	query := `
	CREATE TABLE IF NOT EXISTS urls (
		id BIGINT AUTO_INCREMENT PRIMARY KEY,
		short_code VARCHAR(10) NOT NULL UNIQUE,
		original_url TEXT NOT NULL,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
	);`

	_, err := DB.Exec(query)
	return err
}

func InitMySQL() error {

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true&tls=false",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASS"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
	)

	log.Printf("MySQL DSN host: %s port: %s\n",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
	)

	var err error

	for i := 0; i < 10; i++ {

		DB, err = sql.Open("mysql", dsn)
		if err != nil {
			log.Printf("sql.Open error: %v\n", err)
			time.Sleep(3 * time.Second)
			continue
		}

		err = DB.Ping()
		if err == nil {
			log.Println("Connected to MySQL!")
			return migrate()
		}

		log.Printf("Waiting for MySQL... attempt %d error: %v\n", i+1, err)
		time.Sleep(3 * time.Second)
	}

	return fmt.Errorf("failed to connect mysql: %v", err)
}
