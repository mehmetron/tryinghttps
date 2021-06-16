package main

import (
	"encoding/json"
	"fmt"
	"github.com/brianvoe/gofakeit/v6"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"log"
	"net/http"
	"time"
)

type Course struct {
	Id            int
	CourseName    string `db:"course_name"`
	Description   string
	PublishedDate time.Time `db:"published_date"`
}

var Schema = `
DROP TABLE IF EXISTS course;

CREATE TABLE course(
	id serial primary key,
	course_name varchar(255) NOT NULL,
	description varchar(500),
	published_date DATE NOT NULL DEFAULT CURRENT_DATE
);

`

var db *sqlx.DB

func init() {
	// this Pings the database trying to connect
	// use sqlx.Open() for sql.Open() semantics
	d, err := sqlx.Open("postgres", "user=postgres password=turhan99 dbname=postgres host=postgresql port=5432 sslmode=disable TimeZone=Asia/Istanbul")
	if err != nil {
		log.Fatalln(err)
	}
	db = d

	// exec the schema or fail; multi-statement Exec behavior varies between
	// database drivers;  pq will exec them all, sqlite3 won't, ymmv
	fmt.Println("Create DB schema...")
	db.MustExec(Schema)

	fmt.Println("Populating DB...")
	populateDB()
}

func populateDB() {

	// batch insert courses with maps
	courseMaps := []map[string]interface{}{
		{"courseName": gofakeit.Sentence(5), "description": gofakeit.Sentence(10)},
		{"courseName": gofakeit.Sentence(5), "description": gofakeit.Sentence(10)},
		{"courseName": gofakeit.Sentence(5), "description": gofakeit.Sentence(10)},
		{"courseName": gofakeit.Sentence(5), "description": gofakeit.Sentence(10)},
		{"courseName": gofakeit.Sentence(5), "description": gofakeit.Sentence(10)},
		{"courseName": gofakeit.Sentence(5), "description": gofakeit.Sentence(10)},
	}
	_, err := db.NamedExec(`INSERT INTO course (course_name, description) VALUES (:courseName, :description)`, courseMaps)
	if err != nil {
		fmt.Println(err)
	}

}

func GetAllCourses() []Course {
	var courses []Course
	db.Select(&courses, "SELECT * FROM course ORDER BY course_name ASC")
	course1, course2 := courses[0], courses[1]

	fmt.Printf("%#v\n%#v", course1, course2)
	return courses

}

// Routes ------------------------------------------------------------------------------------------------------------

func courses(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	courses := GetAllCourses()
	json.NewEncoder(w).Encode(courses)
}

func hello(w http.ResponseWriter, req *http.Request) {

	fmt.Fprintf(w, "hopefully it'l work this time \n")
}

func headers(w http.ResponseWriter, req *http.Request) {

	for name, headers := range req.Header {
		for _, h := range headers {
			fmt.Fprintf(w, "%v: %v\n", name, h)
		}
	}
}

func main() {

	http.HandleFunc("/", hello)
	http.HandleFunc("/headers", headers)
	http.HandleFunc("/courses", courses)

	http.ListenAndServe(":8090", nil)
}
