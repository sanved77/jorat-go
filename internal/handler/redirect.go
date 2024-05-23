package handler

import (
	"database/sql"
	"net/http"

	_ "github.com/lib/pq"
)

func Redirect(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	path := r.URL.Path[1:]

	if url, ok := FindUrl(path, db); ok {
		http.Redirect(w, r, url, http.StatusFound)
	} else {
		http.Error(w, "URL Doesn't exist", http.StatusNotFound)
	}
}
