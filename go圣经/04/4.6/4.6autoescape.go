package main

import (
	"html/template"
	"log"
	"os"
)

func main() {
	const templ = `<p>A: {{.A}}</p><p>B: {{.B}}</p>`
	t := template.Must(template.New("escape").Parse(templ))
	var data struct {
		A string	// untrusted plain text
		B template.HTML // trusted HTML
	}

	data.A = "<b>Hello! A </b>"
	data.B = "<b>Hello! B </b>"

	if err := t.Execute(os.Stdout,data) ; err != nil { log.Fatal(err)}
}
