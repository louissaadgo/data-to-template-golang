package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

type food struct {
	Name string
	ID   string
}

func init() {
	tpl = template.Must(template.ParseGlob("*.gohtml"))
}

func main() {
	slice := []string{
		"GO",
		"Golang",
	}
	lang := map[string]string{
		"Hi":           "Hola",
		"How Are You?": "Como Estas?",
	}
	burger := food{
		"Burger",
		"BR",
	}
	//passing a slice
	err := tpl.ExecuteTemplate(os.Stdout, "slice.gohtml", slice)
	if err != nil {
		log.Fatalln(err)
	}
	//passing a slice as variable
	err = tpl.ExecuteTemplate(os.Stdout, "slicevar.gohtml", slice)
	if err != nil {
		log.Fatalln(err)
	}
	//passing a map
	err = tpl.ExecuteTemplate(os.Stdout, "map.gohtml", lang)
	if err != nil {
		log.Fatalln(err)
	}
	//passing a map as variable
	err = tpl.ExecuteTemplate(os.Stdout, "mapvar.gohtml", lang)
	if err != nil {
		log.Fatalln(err)
	}
	//passing a struct
	err = tpl.ExecuteTemplate(os.Stdout, "struct.gohtml", burger)
	if err != nil {
		log.Fatalln(err)
	}
	//passing a struct as var
	err = tpl.ExecuteTemplate(os.Stdout, "structvar.gohtml", burger)
	if err != nil {
		log.Fatalln(err)
	}
}
