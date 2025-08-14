package main

import (
	"html/template"
	"os"
)

type Curso struct {
	Nome         string
	CargaHoraria int
}

type Cursos []Curso

func main() {
	templates := []string{
		"header.html",
		"content.html",
		"footer.html",
	}
	t := template.Must(template.New("content.html").ParseFiles(templates...))
	err := t.Execute(os.Stdout, Cursos{
		{"Go", 40},
		{"Java", 20},
		{"Python", 10},
	})
	if err != nil {
		// w.WriteHeader(http.StatusInternalServerError)
		panic(err)
	}
	// http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

	// })
	// http.ListenAndServe(":8080", nil)
}
