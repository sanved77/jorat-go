package handler

import (
	"database/sql"
	"encoding/json"
	"net/http"

	_ "github.com/lib/pq"
)

type CreateReqPayload struct {
	Short   string `json:"short"`
	LongUrl string `json:"long_url"`
}

type ResponseMessage struct {
	StatusCode int    `json:"statusCode"`
	Message    string `json:"message"`
}

func writeHeader(w http.ResponseWriter, msg string, code int) {
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(ResponseMessage{
		StatusCode: code,
		Message:    msg,
	})
}

func Create(w http.ResponseWriter, r *http.Request, db *sql.DB) {

	var payload CreateReqPayload

	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		writeHeader(w, "Bad JSON", http.StatusNotAcceptable)
		return
	}

	if payload.Short == "" || payload.LongUrl == "" {
		writeHeader(w, "bad input", http.StatusBadRequest)
		return
	}

	var exists bool
	if err := db.QueryRow("SELECT EXISTS(SELECT 1 FROM short WHERE short_code=$1)", payload.Short).Scan(&exists); err != nil {
		writeHeader(w, "Issues retriving from database: "+err.Error(), http.StatusInternalServerError)
		return
	}
	if exists {
		writeHeader(w, "Short code already in use", http.StatusConflict)
		return
	}

	if _, err := db.Exec("INSERT INTO short (short_code, long_url) VALUES ($1, $2)", payload.Short, payload.LongUrl); err != nil {
		writeHeader(w, "issues working with the database: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(struct {
		StatusCode int    `json:"statusCode"`
		Message    string `json:"payload"`
	}{
		StatusCode: 200,
		Message:    "Short URL creation successful",
	})
}
