package handler

import "database/sql"

func FindUrl(short_code string, db *sql.DB) (string, bool) {
	var url string
	err := db.QueryRow("SELECT long_url FROM short WHERE short_code = $1", short_code).Scan(&url)

	if err != nil {
		return "", false
	}
	return url, true
}
