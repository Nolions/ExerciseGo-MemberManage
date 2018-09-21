package main

import (
	"fmt"
	"html/template"
	"log"
)

const templateLayout = "views/layout.html"
const templateHeader = "views/head.html"
const templateFooter = "views/head.html"
const templateJS = "views/jsfile.html"
const templateCSS = "views/cssfile.html"

func setLayout(pagePath string) {
	t, err := template.ParseFiles(templateLayout, templateHeader, templateCSS, templateJS, templateFooter, pagePath)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(t)
}
