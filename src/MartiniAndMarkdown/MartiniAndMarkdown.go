package main 

import (
        "github.com/go-martini/martini"
        "github.com/russross/blackfriday"
        "net/http"
)

//Before running this, move index.html to the public folder


func main() {
        m := martini.Classic()
        m.Post("/generate", func(r *http.Request) []byte {
                body := r.FormValue("body")
                return blackfriday.MarkdownBasic([]byte(body))
        })

        m.Run()
}

