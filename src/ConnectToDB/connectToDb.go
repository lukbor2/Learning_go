package main 

//Connecting to MySql db using the go-sql-driver. There are other alternatives but I used this one

import (
        "database/sql"
        "fmt"
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
		fmt.Println("Starting now")
        m := martini.Classic()
        m.Map(SetupDB())
        m.Get("/", func(db *sql.DB, r *http.Request, rw http.ResponseWriter) {
        //search := "%" + r.URL.Query().Get("search") + "%"
                rows, err := db.Query("SELECT title, author, description FROM books")
                PanicIf(err)
                defer rows.Close()
                var title, author, description string
                for rows.Next() {
                        err := rows.Scan(&title, &author, &description)
                        PanicIf(err)
                fmt.Fprintf(rw, "Title: %s\nAuthor: %s\nDescription: %s\n\n", title, author, description)
                }
        })
        m.Run()
}

