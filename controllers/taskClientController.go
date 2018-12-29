package controllers

import (
	"fmt"
	MTaskClient "github.com/freeddser/ddser/models"
	"github.com/freeddser/ddser/util"
	"net/http"
)

//显示tasklist to Client
func TaskClientShowHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method == "GET" {
		r.ParseForm()
		switch r.FormValue("type") {
		case "auth_token":
			w.Header().Set("Content-Type", "application/json")
			token_id := MTaskClient.GetTokenId(r.FormValue("token"))
			if token_id == 0 {
				fmt.Fprintln(w, `{"token":"invaild"}`)
			} else {
				fmt.Fprintln(w, MTaskClient.ShowClientCurrutTokenTask(token_id))
			}
		case "update_task_status":
			w.Header().Set("Content-Type", "application/json")
			token_id := MTaskClient.GetTokenId(r.FormValue("token"))
			if token_id == 0 {
				fmt.Fprintln(w, `{"token":"invaild"}`)
			} else {
				task_id := r.FormValue("task_id")
				status := r.FormValue("status")
				if util.StringToInt(status) == 1001 {
					MTaskClient.AddOneTaskLogForClient(util.StringToInt(task_id), token_id)
				}
				fmt.Fprintln(w, MTaskClient.UpdateOneTaskStatusForClient(util.StringToInt(task_id), util.StringToInt(status), token_id).RowsAffected)
			}

		case "update_task_log":
			w.Header().Set("Content-Type", "application/json")
			token_id := MTaskClient.GetTokenId(r.FormValue("token"))
			if token_id == 0 {
				fmt.Fprintln(w, `{"token":"invaild"}`)
			} else {
				task_id := r.FormValue("task_id")
				task_logs := util.HtmlToMysql(r.FormValue("task_logs"))
				client_infos := r.FormValue("client_infos")
				task_status := r.FormValue("task_status")
				pusher := r.FormValue("pusher")
				fmt.Fprintln(w, MTaskClient.UpdateTaskLogForClient(util.StringToInt(task_id), token_id, task_logs, client_infos, task_status, pusher).RowsAffected)
			}
		default:
			w.Header().Set("Content-Type", "application/json")
			fmt.Fprintln(w, `{"access":"invaild"}`)
		}
	}
}
