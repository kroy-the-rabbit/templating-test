package main

import (
	"os"
	"text/template"
)

func main() {
	tmpl, err := template.ParseFiles("templates/index.go.html")
	if err != nil {
		panic(err)
	}

	data := struct {
		Name  string
		Items []string
	}{
		Name:  "MyName",
		Items: []string{"Item1", "Item2", "Item3"},
	}

	// Render the template with the context
	err = tmpl.Execute(os.Stdout, data)
	if err != nil {
		panic(err)
	}
}
