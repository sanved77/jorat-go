package main

import (
	"log"
	"net/http"

	_ "github.com/lib/pq"
	"github.com/sanved77/jorat-go/config"
	"github.com/sanved77/jorat-go/internal/handler"
	"github.com/sanved77/jorat-go/internal/storage"
)

func main() {

	config.Init()

	PORT := config.GetString("server.PORT")

	db := storage.DBConnect()
	defer db.Close()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/" {
			http.ServeFile(w, r, "/var/www/jor.at/index.html")
			return
		}
		handler.Redirect(w, r, db)
	})

	http.HandleFunc("/create", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Method not allowed, use POST", http.StatusBadRequest)
		}
		handler.Create(w, r, db)
	})

	log.Printf("Server starting on port %s", PORT)
	log.Fatal(http.ListenAndServe(PORT, nil))
}
