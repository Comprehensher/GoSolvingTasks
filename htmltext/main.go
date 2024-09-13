package main

import (
	"os"
	"text/template"
)

func Exec(t *template.Template) error {
	return t.Execute(os.Stdout, &Kayak)
}
func main() {
	allTemplates, err :=
		template.ParseGlob("templates/*.html")
	if err == nil {
		selectedTemplated :=
			allTemplates.Lookup("template.html")
		err = Exec(selectedTemplated)
	}
	if err != nil {
		Printfln("Error: %v %v", err.Error())
	}
}
