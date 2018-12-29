package controllers

import (
	"fmt"
	MRole "github.com/freeddser/ddser/models"
	"github.com/freeddser/ddser/util"
	"net/http"
	"strings"
)

func RoleResourceManager(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		r.ParseForm()
		switch r.FormValue("type") {
		case "getallroleresource":
			w.Header().Set("Content-Type", "application/json")
			fmt.Fprintln(w, MRole.GetAllRoleResource())

		case "getoneroleresource":
			w.Header().Set("Content-Type", "application/json")
			fmt.Fprintln(w, MRole.GetCurrentRoleResource(util.StringToInt(r.FormValue("id"))))
		default:
			locals := make(map[string]interface{})
			locals["rolename"] = MRole.GetAllRoleResourceInfos()
			util.RenderHtml(w, "roleresource", locals)
		}
	}

	if r.Method == "POST" {
		r.ParseForm()
		switch r.PostFormValue("type") {
		case "update":
			roleIdCX := r.PostFormValue("roleidCX")
			roleId := strings.Replace(roleIdCX, "cb1=", "", -1)
			id := strings.Split(roleId, "&")
			MRole.DelOneRoleResource(util.StringToInt(r.PostFormValue("role_id")))
			for _, v := range id {
				MRole.AddOneRoleResource(util.StringToInt(r.PostFormValue("role_id")), util.StringToInt(v))
			}
		default:
			w.Header().Set("Content-Type", "application/json")
			fmt.Fprintln(w, MRole.GetAllRoleResourceInfos())

		}

	}

}
