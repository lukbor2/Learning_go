package main 
//Note I had to save the template folder in the root of the GOPATH variable
import (
        "database/sql"
        "github.com/go-martini/martini"
        "github.com/martini-contrib/render"
        _ "github.com/go-sql-driver/mysql"
        "net/http"
)
type Book struct {
        Title string
        Author string
        Description string
}

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
        m.Use(render.Renderer(render.Options{Layout: "layout",}))
        m.Get("/", func(ren render.Render, r *http.Request, db *sql.DB) {
                rows, err := db.Query("SELECT title, author, description FROM books")
                PanicIf(err)
                defer rows.Close()
                books := []Book{}
                for rows.Next() {
                        book := Book{}
                        err := rows.Scan(&book.Title, &book.Author, &book.Description)
                        PanicIf(err)
                        books = append(books, book)
                }
        ren.HTML(200, "books", books)
        })
        m.Run()
}

