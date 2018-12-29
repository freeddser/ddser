package controllers

import (
	MResource "github.com/freeddser/ddser/models"
	"github.com/freeddser/ddser/util"

	"fmt"
	"net/http"
)

func ResourceManager(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		r.ParseForm()
		switch r.FormValue("type") {
		case "delete":
			fmt.Fprintln(w, MResource.DelOneResource(util.StringToInt(r.FormValue("id"))))
		default:
			locals := make(map[string]interface{})
			locals["rolename"] = MResource.GetAllResourceInfos()
			util.RenderHtml(w, "resource", locals)
		}
	}

	if r.Method == "POST" {
		w.Header().Set("Content-Type", "application/json")
		r.ParseForm()
		fmt.Println(r.PostFormValue("type"))
		switch r.PostFormValue("type") {
		case "add":
			fmt.Println(r.PostFormValue("id"), r.PostFormValue("name"), r.PostFormValue("desc"), r.PostFormValue("url"))
			fmt.Fprintln(w, MResource.AddOneResource(r.PostFormValue("name"), r.PostFormValue("desc"), r.PostFormValue("url"), util.StringToInt(r.PostFormValue("id"))).RowsAffected)
		case "update":
			fmt.Fprintln(w, MResource.UpdateOneResource(r.PostFormValue("name"), r.PostFormValue("desc"), r.PostFormValue("url"), util.StringToInt(r.PostFormValue("id"))).RowsAffected)
		default:
			fmt.Fprintln(w, MResource.GetAllResourceInfos())
		}
	}
}
