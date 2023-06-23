package main

import (
	"flag"
	"html/template"
	"io/ioutil"
	"os"

	"github.com/tevino/log"
)

type Page struct {
	Contents string
}

func generatePage(dir string, filePath string) {
	log.SetOutputLevel(log.INFO)

	lastDirChar := dir[len(dir)-1:]

	if lastDirChar != "/" {
		dir += "/"
	}

	fileContents, fileContentsErr := ioutil.ReadFile(dir + filePath)

	if fileContentsErr != nil {
		log.Info("Error reading file: ", fileContentsErr)
	}

	page := Page{
		Contents: string(fileContents),
	}

	t := template.Must(template.ParseFiles("template.tmpl"))

	name := string(filePath)[:len(string(filePath))-4]
	newFile, err := os.Create(dir + name + ".html")

	if err != nil {
		log.Info("Error creating file: ", err)
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
			log.Info("Error reading directory: ", filesErr)
		}

		for _, file := range files {
			lastChar := file.Name()[len(file.Name())-4:]

			if lastChar == ".txt" {
				generatePage(*dir, file.Name())
			}
		}
	}

}
