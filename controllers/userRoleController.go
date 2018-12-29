package controllers

import (
	"fmt"
	MRole "github.com/freeddser/ddser/models"
	"github.com/freeddser/ddser/util"
	"net/http"
	"strings"
)

func UserRoleManager(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		r.ParseForm()
		switch r.FormValue("type") {
		case "getallrole":
			w.Header().Set("Content-Type", "application/json")
			fmt.Fprintln(w, MRole.GetAllRole())

		case "getonerole":
			w.Header().Set("Content-Type", "application/json")
			fmt.Fprintln(w, MRole.GetCurrentUserRole(util.StringToInt(r.FormValue("id"))))

		default:
			locals := make(map[string]interface{})
			// locals["username"] = sess.Get("username")
			locals["rolename"] = MRole.GetAllUserRoleInfos()
			util.RenderHtml(w, "userrole", locals)
		}

	}
	if r.Method == "POST" {
		r.ParseForm()
		switch r.PostFormValue("type") {
		case "update":
			a := r.PostFormValue("roleidCX")
			c := strings.Replace(a, "cb1=", "", -1)
			d := strings.Split(c, "&")
			MRole.DelOneUserRole(util.StringToInt(r.PostFormValue("userid")))
			for _, v := range d {
				MRole.AddOneUserRole(util.StringToInt(r.PostFormValue("userid")), util.StringToInt(v))
			}
		default:
			w.Header().Set("Content-Type", "application/json")
			fmt.Fprintln(w, MRole.GetAllUserRoleInfos())

		}

	}

}
