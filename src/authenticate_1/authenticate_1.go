package main

import (
	"database/sql"
	"github.com/go-martini/martini"
	_ "github.com/go-sql-driver/mysql"
	"net/http"
)

func SetupDB() *sql.DB {
	db, err := sql.Open("mysql", "root:cacca1971@/lesson4")
	PanicIf(err)
	return db
}

func PanicIf(err error) {
	if err != nil {
		panic(err)
	}
}
func main() {
	m := martini.Classic()
	m.Map(SetupDB())
	m.Post("/login", PostLogin)
	m.Run()
}

func PostLogin(req *http.Request, db *sql.DB) (int, string) {
	var id string
	email, password := req.FormValue("email"), req.FormValue("password")
	stmtOut, err := db.Prepare("SELECT id FROM lesson4.users where email=? and password=?")
	if err != nil {
		panic(err)
	}
//The statement err := db.QueryRow("select id from users where email=$1 and password=$2", email, password).Scan(&id)
//does not seem to work with MySQL. So I used db.Prepare first.	
	err2 := stmtOut.QueryRow(email, password).Scan(&id)
	if err2 != nil {
		return 401, "Unauthorized"
	}
	return 200, "User id is " + id
}
