package main

import (
	"flag"
	"html/template"
	"io/ioutil"
	"os"
	// "fmt"
)

type Page struct {
	Contents string
}

func main() {
	filePath := flag.String("file", "", "Get file from command-line input")
	flag.Parse()

	fileContents, fileContentsErr := ioutil.ReadFile(*filePath)

	if fileContentsErr != nil {
		panic(fileContentsErr)
	}

	page := Page{
		Contents: string(fileContents),
	}

	t := template.Must(template.ParseFiles("template.tmpl"))

	newFile, err := os.Create("new.html")

	if err != nil {
		panic(err)
	}

	t.Execute(newFile, page)
	// t := template.Must(template.New("template.tmpl").ParseFiles("template.tmpl"))

	// newFile, err := os.Create("new.html")

	// if err != nil {
	// 	panic(err)
	// }

	// t.Execute(newFile, page)
}
