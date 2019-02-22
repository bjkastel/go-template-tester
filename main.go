package main

import (
	"github.com/Masterminds/sprig"
	"html/template"
	"os"
)

func main() {
	templateText := os.Args[1]
	texts := os.Args[2:]

	t, err := template.New("template").Funcs(sprig.FuncMap()).Parse(templateText + "\n")
	if err != nil {
		panic(err)
	}

	for _, text := range texts {
		err = t.Execute(os.Stdout, text)
		if err != nil {
			panic(err)
		}
	}
}
