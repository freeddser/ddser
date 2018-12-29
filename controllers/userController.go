package controllers

import (
	"fmt"
	MUser "github.com/freeddser/ddser/models"
	"github.com/freeddser/ddser/util"
	"net/http"
)

func UserPrivilegesManager(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		r.ParseForm()
		switch r.FormValue("type") {
		case "delete":
			w.Header().Set("Content-Type", "application/json")
			fmt.Fprintln(w, MUser.DelOneUser(util.StringToInt(r.FormValue("id"))))
		case "showoneuser":
			w.Header().Set("Content-Type", "application/json")
			fmt.Fprintln(w, MUser.GetOneUserInfo(util.StringToInt(r.FormValue("id"))))

		default:
			locals := make(map[string]interface{})
			locals["username"] = MUser.GetAllUserInfos()
			util.RenderHtml(w, "userprivileges", locals)
		}

	}
	if r.Method == "POST" {
		r.ParseForm()
		switch r.PostFormValue("type") {
		case "add":
			w.Header().Set("Content-Type", "application/json")
			fmt.Fprintln(w, MUser.AddOneUser(r.PostFormValue("username"), r.PostFormValue("password")).RowsAffected)
		case "update":
			w.Header().Set("Content-Type", "application/json")
			fmt.Fprintln(w, MUser.UpdateOneUser(r.PostFormValue("modifyusername"), r.PostFormValue("modifypassword"), util.StringToInt(r.PostFormValue("modifyid"))).RowsAffected)
		default:
			//response json
			w.Header().Set("Content-Type", "application/json")
			fmt.Fprintln(w, MUser.GetAllUserInfos())

		}

	}

}
