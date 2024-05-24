package handler

import (
	"database/sql"
	"log"
	"net/http"

	_ "github.com/lib/pq"
)

func incrementVisits(short_code string, db *sql.DB) error {
	_, err := db.Exec("UPDATE short SET visits = visits + 1 WHERE short_code = $1", short_code)
	if err != nil {
		return err
	}
	return nil
}

func Redirect(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	path := r.URL.Path[1:]

	if url, ok := FindUrl(path, db); ok {
		go func(short_code string) {
			err := incrementVisits(short_code, db)
			if err != nil {
				log.Printf("Error incrementing visits: %v", err)
			}
		}(path)
		http.Redirect(w, r, url, http.StatusFound)
	} else {
		http.Error(w, "URL Doesn't exist", http.StatusNotFound)
	}
}
