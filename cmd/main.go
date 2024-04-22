package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"os"

	"github.com/balazskvancz/image-gallery-hw/internal"
	"github.com/balazskvancz/image-gallery-hw/internal/repository"

	_ "github.com/go-sql-driver/mysql"
)

type dbConfig struct {
	Host     string `json:"host"`
	Port     int    `json:"port"`
	User     string `json:"user"`
	Password string `json:"password"`
	Database string `json:"database"`
}

const configPath string = "./config.json"

func main() {
	config := getConfig()
	if config == nil {
		fmt.Printf("Nincs adatbázis config!")
		os.Exit(1)
	}

	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", config.User, config.Password, config.Host, config.Port, config.Database))
	if err != nil {
		fmt.Printf("Nincs adatbázis kapcsolat!")
		os.Exit(1)
	}
	defer db.Close()

	s := internal.NewServer(repository.New(db))

	s.Listen()
}

func getConfig() *dbConfig {
	b, err := os.ReadFile(configPath)
	if err != nil {
		return nil
	}

	var c dbConfig

	if err := json.Unmarshal(b, &c); err != nil {
		return nil
	}

	return &c
}
