package main

import (
	"bytes"
	"fmt"
	"github.com/Masterminds/sprig"
	"os"
	"strconv"
	"text/template"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Printf("%s [TEMPLATE] [INPUT]...", os.Args[0])
		os.Exit(1)
	}

	templateText := os.Args[1]
	texts := os.Args[2:]

	maxLength := 0
	for _, text := range texts {
		if len(text) > maxLength {
			maxLength = len(text)
		}
	}

	t, err := template.New("template").Funcs(sprig.TxtFuncMap()).Parse(templateText)
	if err != nil {
		panic(err)
	}

	for _, text := range texts {
		var buffer bytes.Buffer
		err = t.Execute(&buffer, text)
		if err != nil {
			panic(err)
		}
		fmt.Printf("%-"+strconv.Itoa(maxLength)+"s: %s\n", text, buffer.String())
	}
}
