package main

import (
	"flag"
	"html/template"
	"io/ioutil"
	"os"
)

type Page struct {
	Contents string
}

func generatePage(dir string, filePath string) {
	lastDirChar := dir[len(dir)-1:]

	if lastDirChar != "/" {
		dir += "/"
	}

	fileContents, fileContentsErr := ioutil.ReadFile(dir + filePath)

	if fileContentsErr != nil {
		panic(fileContentsErr)
	}

	page := Page{
		Contents: string(fileContents),
	}

	t := template.Must(template.ParseFiles("template.tmpl"))

	name := string(filePath)[:len(string(filePath))-4]
	newFile, err := os.Create(dir + name + ".html")

	if err != nil {
		panic(err)
	}

	t.Execute(newFile, page)
}

func main() {
	filePath := flag.String("file", "", "Get file from command-line input")
	dir := flag.String("dir", "", "Find all .txt files and generate separate HTML files for each")
	flag.Parse()

	if *dir == "" {
		generatePage("./", *filePath)
	} else {
		files, filesErr := ioutil.ReadDir(*dir)

		if filesErr != nil {
			panic(filesErr)
		}

		for _, file := range files {
			lastChar := file.Name()[len(file.Name())-4:]

			if lastChar == ".txt" {
				generatePage(*dir, file.Name())
			}
		}
	}

}
