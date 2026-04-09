package main

import (
	"database/sql"
	"flag"
	"log/slog"
	"net/http"
	"os"

	_ "github.com/lib/pq"
	"github.com/slikhithreddy22/snippetbox/internal/models"
)

type application struct {
	logger   *slog.Logger
	snippets *models.SnippetModel
}

func main() {
	addr := flag.String("addr", ":4000", "HTTP network address")
	dsn := flag.String("dsn", "postgres://web:pass@localhost/snippetbox?sslmode=disable", "My DB Name")
	flag.Parse()

	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	db, err := openDB(*dsn)
	if err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}

	defer db.Close()
	app := application{
		logger:   logger,
		snippets: &models.SnippetModel{DB: db},
	}

	logger.Info("Starting server", slog.String("addr", ":4000"))

	mux := app.routes()
	err = http.ListenAndServe(*addr, mux)
	logger.Error(err.Error())
	os.Exit(1)
}

func openDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		db.Close()
		return nil, err
	}
	return db, err
}
