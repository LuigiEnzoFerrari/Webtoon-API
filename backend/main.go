package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

// Taking the credentials to log into database
var dbuser string = os.Getenv("DBUSER");
var dbname string = os.Getenv("DBNAME");
var dbpass string = os.Getenv("DBPASS");
var dbhost string = os.Getenv("DBHOST");
var connStr = "user=" + dbuser +
" dbname=" + dbname +
" password=" + dbpass +
" host=" + dbhost +
" sslmode=disable"

var fields string = `webtoon_id, title, genre, summary,
episodes, 'Create by', view,
subscribe, grade, released_date,
url, likes, 'Written by'`

// Webtoons data structure
type Webtoon struct {
	Title sql.NullString
	Info  WebtoonInfo
}

type WebtoonInfo struct {
	Id            int
	Description   sql.NullString
	Genre         sql.NullString
	Summary       sql.NullString
	Episodes      int
	Created_by    sql.NullString
	View          sql.NullString
	Subscribers   sql.NullString
	Grade         float32
	Released_date sql.NullString
	Url           sql.NullString
	Likes         sql.NullString
	Written_by    sql.NullString
}

func displayWebtoons(w http.ResponseWriter, r *http.Request) {
	var (
		webtoons     Webtoon
		webtoonInfo  *WebtoonInfo
		webtoonsList []Webtoon
	)

	query := "SELECT " + fields + " FROM dataset;"

	webtoonInfo = &webtoons.Info
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	row, err := db.Query(query)
	if err != nil {
		log.Fatal(err)
	}
	defer row.Close()
	for row.Next() {
		err := row.Scan(
			&webtoonInfo.Id, &webtoons.Title,
			&webtoonInfo.Genre,
			&webtoonInfo.Summary, &webtoonInfo.Episodes,
			&webtoonInfo.Created_by, &webtoonInfo.View,
			&webtoonInfo.Subscribers, &webtoonInfo.Grade,
			&webtoonInfo.Released_date,
			&webtoonInfo.Url,
			&webtoonInfo.Likes, &webtoonInfo.Written_by)
		if err != nil {
			log.Println(err)
		}
		webtoonsList = append(webtoonsList, webtoons)
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(webtoonsList)
}

func handleTitle(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	title := vars["title"]

	var (
		webtoons     Webtoon
		webtoonInfo  *WebtoonInfo
		webtoonsList []Webtoon
	)
	
	query := "SELECT " + fields + " FROM dataset WHERE title='" + title + "';"
	webtoonInfo = &webtoons.Info
	db, err := sql.Open("postgres",
		connStr)
	if err != nil {
		log.Fatal(err)
	}
	row, err := db.Query(query)
	if err != nil {
		log.Fatal(err)
	}
	defer row.Close()
	if row.Next() {
		err := row.Scan(
			&webtoonInfo.Id, &webtoons.Title,
			&webtoonInfo.Genre,
			&webtoonInfo.Summary, &webtoonInfo.Episodes,
			&webtoonInfo.Created_by, &webtoonInfo.View,
			&webtoonInfo.Subscribers, &webtoonInfo.Grade,
			&webtoonInfo.Released_date,
			&webtoonInfo.Url,
			&webtoonInfo.Likes, &webtoonInfo.Written_by)
		if err != nil {
			log.Fatal(err)
		}
		webtoonsList = append(webtoonsList, webtoons)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(webtoonsList)
	}
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Welcome to webtoon database")
	}).Methods("GET")
	router.HandleFunc("/webtoons", displayWebtoons).Methods("GET", "POST")
	router.HandleFunc("/search", handleTitle).Queries("title", "{title}")
	log.Println("Server is running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
