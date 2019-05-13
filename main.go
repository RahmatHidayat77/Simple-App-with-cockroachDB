package main

import (
	"database/sql"
	"log"
	"net/http"
	"text/template"

	_ "github.com/lib/pq"
)

type kontak struct {
	Id     int
	Nama   string
	Alamat string
	Telp   string
}

func dbConn() (db *sql.DB) {
	// dbDriver := "postgres"
	dbUser := "root"
	dbName := "kontak_db"
	db, err := sql.Open("postgres", "postgresql://"+dbUser+"@localhost:26257/"+dbName+"?sslmode=disable")
	if err != nil {
		panic(err.Error())
	}
	return db
}

var html = template.Must(template.ParseGlob("form/*"))

func Index(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	selDB, err := db.Query("SELECT * FROM kontak ORDER BY id DESC")
	if err != nil {
		panic(err.Error())
	}
	emp := kontak{}
	res := []kontak{}
	for selDB.Next() {
		var id int
		var nama, alamat, telp string
		err = selDB.Scan(&id, &nama, &alamat, &telp)
		if err != nil {
			panic(err.Error())
		}
		emp.Id = id
		emp.Nama = nama
		emp.Alamat = alamat
		emp.Telp = telp
		res = append(res, emp)
	}
	html.ExecuteTemplate(w, "Index", res)
	defer db.Close()
}

func Show(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	nId := r.URL.Query().Get("id")
	selDB, err := db.Query("SELECT * FROM kontak WHERE id=$1", nId)
	if err != nil {
		panic(err.Error())
	}
	emp := kontak{}
	for selDB.Next() {
		var id int
		var nama, alamat, telp string
		err = selDB.Scan(&id, &nama, &alamat, &telp)
		if err != nil {
			panic(err.Error())
		}
		emp.Id = id
		emp.Nama = nama
		emp.Alamat = alamat
		emp.Telp = telp
	}
	html.ExecuteTemplate(w, "Show", emp)
	defer db.Close()
}

func New(w http.ResponseWriter, r *http.Request) {
	html.ExecuteTemplate(w, "New", nil)
}

func Edit(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	nId := r.URL.Query().Get("id")
	selDB, err := db.Query("SELECT * FROM kontak WHERE id=$1", nId)
	if err != nil {
		panic(err.Error())
	}
	emp := kontak{}
	for selDB.Next() {
		var id int
		var nama, alamat, telp string
		err = selDB.Scan(&id, &nama, &alamat, &telp)
		if err != nil {
			panic(err.Error())
		}
		emp.Id = id
		emp.Nama = nama
		emp.Alamat = alamat
		emp.Telp = telp
	}
	html.ExecuteTemplate(w, "Edit", emp)
	defer db.Close()
}

func Insert(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	if r.Method == "POST" {
		nama := r.FormValue("nama")
		alamat := r.FormValue("alamat")
		telp := r.FormValue("telp")
		insForm, err := db.Prepare("INSERT INTO kontak (nama, alamat, telp) VALUES ($1,$2,$3)")
		log.Println("==========")
		log.Println(err)
		log.Println(telp)
		if err != nil {
			panic(err.Error())
		}
		insForm.Exec(nama, alamat, telp)
		log.Println("INSERT: nama: " + nama + " | alamat: " + alamat + " | telp " + telp)
	}
	defer db.Close()
	http.Redirect(w, r, "/", 301)
}

func Update(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	if r.Method == "POST" {
		nama := r.FormValue("nama")
		alamat := r.FormValue("alamat")
		telp := r.FormValue("telp")
		id := r.FormValue("uid")
		insForm, err := db.Prepare("UPDATE kontak SET nama=$1, alamat=$2, telp=$3 WHERE id=$4")
		if err != nil {
			panic(err.Error())
		}
		insForm.Exec(nama, alamat, telp, id)
		log.Println("UPDATE: Nama: " + nama + " | Alamat " + alamat + " | No. Telepon " + telp)
	}
	defer db.Close()
	http.Redirect(w, r, "/", 301)
}

func Delete(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	emp := r.URL.Query().Get("id")
	delForm, err := db.Prepare("DELETE FROM kontak WHERE id=$1")
	if err != nil {
		panic(err.Error())
	}
	delForm.Exec(emp)
	log.Println("DELETE")
	defer db.Close()
	http.Redirect(w, r, "/", 301)
}

func main() {
	log.Println("Server started on: http://localhost:8010")
	http.HandleFunc("/", Index)
	http.HandleFunc("/show", Show)
	http.HandleFunc("/new", New)
	http.HandleFunc("/edit", Edit)
	http.HandleFunc("/insert", Insert)
	http.HandleFunc("/update", Update)
	http.HandleFunc("/delete", Delete)
	http.ListenAndServe(":8010", nil)
}
