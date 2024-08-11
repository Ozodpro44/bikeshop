package handlers

import (
	"bikeshop/models"
	log "bikeshop/pkg/logger"
	"bikeshop/storage"
	"database/sql"
	"fmt"
	"html/template"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strconv"

	_ "github.com/lib/pq"
)

type handlers struct {
	storage storage.StorageI
	log     log.Log
}

type Handlers struct {
	Storage storage.StorageI
	Log     log.Log
}

var templates = template.Must(template.ParseFiles(
	"templates/layout.html",
	"templates/index.html",
	"templates/admin.html",
))

var db *sql.DB

func init() {
	var err error
	connStr := "user=postgres dbname=postgres sslmode=disable password=postgres"
	db, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	rows, err := db.Query("SELECT id, name, price, image FROM bikes")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Println(err)
	defer rows.Close()

	bikes := []models.Bike{}
	for rows.Next() {
		var bike models.Bike
		if err := rows.Scan(&bike.ID, &bike.Name, &bike.Price, &bike.Image); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		bikes = append(bikes, bike)
	}
	fmt.Println(bikes)

	data := struct {
		Title string
		Bikes []models.Bike
	}{
		Title: "Bike Shop",
		Bikes: bikes,
	}

	if err := templates.ExecuteTemplate(w, "layout.html", data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func AdminHandler(w http.ResponseWriter, r *http.Request) {
	data := struct {
		Title string
	}{
		Title: "Admin Panel",
	}
	if err := templates.ExecuteTemplate(w, "layout.html", data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func AddBikeHandler(w http.ResponseWriter, r *http.Request) {
	name := r.FormValue("name")
	price, err := strconv.ParseFloat(r.FormValue("price"), 64)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	file, _, err := r.FormFile("image")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer file.Close()

	imagePath := filepath.Join("static/images", fmt.Sprintf("%s.jpg", name))
	out, err := os.Create(imagePath)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer out.Close()

	_, err = io.Copy(out, file)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	_, err = db.Exec("INSERT INTO bikes (name, price, image) VALUES ($1, $2, $3)", name, price, fmt.Sprintf("%s.jpg", name))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/admin", http.StatusSeeOther)
}
