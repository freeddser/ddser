package controllers

import (
	"fmt"
	MTokenTaskType "github.com/freeddser/ddser/models"
	"github.com/freeddser/ddser/util"
	"net/http"
)

func RoleTaskTypeManager(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		r.ParseForm()
		switch r.FormValue("type") {
		case "gettokentype":
			w.Header().Set("Content-Type", "application/json")
			fmt.Fprintln(w, MTokenTaskType.GetAllTokenInfos())

		case "getonetokentype":
			w.Header().Set("Content-Type", "application/json")
			fmt.Fprintln(w, MTokenTaskType.GetOneTokenTaskTypelist(r.FormValue("token_id")))

		case "deletetask":
			w.Header().Set("Content-Type", "application/json")
			fmt.Fprintln(w, MTokenTaskType.DelOneTask(util.StringToInt(r.FormValue("id"))).RowsAffected)

		default:
			locals := make(map[string]interface{})
			util.RenderHtml(w, "roletasktype", locals)
		}

	}
	if r.Method == "POST" {
		w.Header().Set("Content-Type", "application/json")
		r.ParseForm()
		switch r.PostFormValue("type") {
		case "addtask":
			fmt.Fprintln(w, MTokenTaskType.AddOneTokenTask(r.PostFormValue("tk_type"), r.PostFormValue("task_shell"), r.PostFormValue("shell_params"), r.PostFormValue("task_type_desc"), util.StringToInt(r.PostFormValue("role_id")), util.StringToInt(r.PostFormValue("token_id"))).RowsAffected)

		default:
			return

		}

	}

}
