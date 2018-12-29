package controllers

import (
	"fmt"
	MRole "github.com/freeddser/ddser/models"
	"github.com/freeddser/ddser/util"
	"net/http"
)

func RoleManager(w http.ResponseWriter, r *http.Request) {

	if r.Method == "GET" {
		r.ParseForm()
		switch r.FormValue("type") {
		case "delete":
			w.Header().Set("Content-Type", "application/json")
			fmt.Fprintln(w, MRole.DelOneRole(util.StringToInt(r.FormValue("id"))))
		case "showonerole":
			w.Header().Set("Content-Type", "application/json")
			fmt.Fprintln(w, MRole.GetOneRoleInfo(util.StringToInt(r.FormValue("id"))))

		default:
			locals := make(map[string]interface{})
			locals["rolename"] = MRole.GetAllRoleInfos()
			util.RenderHtml(w, "role", locals)
		}

	}
	if r.Method == "POST" {
		w.Header().Set("Content-Type", "application/json")
		r.ParseForm()
		switch r.PostFormValue("type") {
		case "add":
			fmt.Fprintln(w, MRole.AddOneRole(r.PostFormValue("rolename")).RowsAffected)
		case "update":
			fmt.Fprintln(w, MRole.UpdateOneRole(r.PostFormValue("modifyrolename"), util.StringToInt(r.PostFormValue("modifyid"))).RowsAffected)
		default:
			fmt.Fprintln(w, MRole.GetAllRoleInfos())

		}

	}

}
