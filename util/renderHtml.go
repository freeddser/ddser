package util

import (
	"github.com/freeddser/ddser/config"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"path"
)

var Templates = make(map[string]*template.Template)

func RecursionDir(templatePath string) {
	fileInfoArr, err := ioutil.ReadDir(templatePath)
	CheckErr(err)
	for _, fileinfo := range fileInfoArr {
		if fileinfo.IsDir() {
			RecursionDir(templatePath + "/" + fileinfo.Name())

		}
		templateName := fileinfo.Name()
		if ext := path.Ext(templateName); ext != ".html" {
			continue
		}
		log.Println("Loading template:", templatePath+"/"+fileinfo.Name())
		t := template.Must(template.New(fileinfo.Name()).Delims("[[", "]]").ParseFiles(templatePath + "/" + fileinfo.Name()))
		Templates[templateName] = t

	}
}

func RenderHtml(w http.ResponseWriter, tmpl string, locals map[string]interface{}) {
	tmpl += config.MustGetString("server.Tpl_Suffix")
	err := Templates[tmpl].Execute(w, locals)
	CheckErr(err)
}
