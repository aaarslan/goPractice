package main

import (
	"html/template"
	"log"
	"os"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("tpl.gohtml"))
}

func main() {
	fe := []string{"Camilla", "Tharja", "Linde", "Celica", "Kagero"}

	err := tpl.Execute(os.Stdout, fe)
	if err != nil {
		log.Fatalln(err)
	}
}
