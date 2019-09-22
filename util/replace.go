package util

import (
	"bytes"
	"log"
	"text/template"

	t "github.com/usagikeri/daimoku/template"
)

func ReplaceTemplate(text string) string {
	s := map[string]string{"Table": text}
	var doc bytes.Buffer

	templ, err := template.New("new").Parse(t.TexTemplate)
	if err != nil {
		log.Fatal(err)
	}

	if err := templ.Execute(&doc, s); err != nil {
		panic(err)
	}

	return doc.String()
}

func ReplaceParts(key string, text string) string {
	var doc bytes.Buffer

	if key == "B3" {
		s := map[string]string{"B3": text}
		templ, err := template.New("new").Parse(t.B3Template)
		if err != nil {
			log.Fatal(err)
		}
		if err := templ.Execute(&doc, s); err != nil {
			log.Fatal(err)
		}
	} else if key == "B4" {
		s := map[string]string{"B4": text}
		templ, err := template.New("new").Parse(t.B4Template)
		if err != nil {
			log.Fatal(err)
		}
		if err := templ.Execute(&doc, s); err != nil {
			log.Fatal(err)
		}
	} else if key == "M1" {
		s := map[string]string{"M1": text}
		templ, err := template.New("new").Parse(t.M1Template)
		if err != nil {
			log.Fatal(err)
		}
		if err := templ.Execute(&doc, s); err != nil {
			log.Fatal(err)
		}
	}

	return doc.String()
}
