package controllers

import (
	"fmt"
	MTask "github.com/freeddser/ddser/models"
	"github.com/freeddser/ddser/util"
	"net/http"
)

func TaskShowLogHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		r.ParseForm()
		task_id := r.FormValue("task_id")
		switch r.FormValue("type") {
		case "showlog":
			w.Header().Set("Content-Type", "application/json")
			fmt.Fprintln(w, MTask.ShowTaskLog(util.StringToInt(task_id)))
		default:
			locals := make(map[string]interface{})
			locals["task_id"] = task_id
			util.RenderHtml(w, "showlog", locals)
		}

	}

}

func TaskAddHandler(w http.ResponseWriter, r *http.Request) {
	sess := GlobalSessions.SessionStart(w, r)
	if r.Method == "GET" {
		r.ParseForm()
		switch r.FormValue("type") {
		case "gettasktype":
			w.Header().Set("Content-Type", "application/json")
			fmt.Fprintln(w, MTask.GetOneTaskType(util.GetTypeInt(sess.Get("user_id"))))
		default:
			locals := make(map[string]interface{})
			locals["rolename"] = MTask.GetOneTaskType(util.GetTypeInt(sess.Get("user_id")))
			util.RenderHtml(w, "taskadd", locals)
		}
	}

	if r.Method == "POST" {
		r.ParseForm()
		switch r.PostFormValue("type") {
		case "add":
			w.Header().Set("Content-Type", "application/json")
			fmt.Fprintln(w, MTask.AddOneTask(util.GetTypeString(sess.Get("username")), r.PostFormValue("task_name"), r.PostFormValue("task_shell"), r.PostFormValue("shell_params"), r.PostFormValue("task_type_desc"), util.StringToInt(r.PostFormValue("task_type")), util.StringToInt(r.PostFormValue("token_id"))).RowsAffected)
		default:
			w.Header().Set("Content-Type", "application/json")

		}

	}

}

func TaskShowHandler(w http.ResponseWriter, r *http.Request) {

	sess := GlobalSessions.SessionStart(w, r)

	if r.Method == "GET" {
		r.ParseForm()
		switch r.FormValue("type") {
		case "showtask":
			w.Header().Set("Content-Type", "application/json")
			fmt.Fprintln(w, MTask.ShowCurrutTokenTask(util.GetTypeInt(sess.Get("user_id"))))
		case "pushtask":
			w.Header().Set("Content-Type", "application/json")
			fmt.Fprintln(w, MTask.UpdateOneTaskStatus(util.StringToInt(r.FormValue("task_id")), 502).RowsAffected)

		case "canceltask":
			w.Header().Set("Content-Type", "application/json")
			fmt.Fprintln(w, MTask.CancelOneTask(util.StringToInt(r.FormValue("task_id"))).RowsAffected)
		case "repushtask":
			w.Header().Set("Content-Type", "application/json")
			fmt.Fprintln(w, MTask.RepushOneTask(util.GetTypeString(sess.Get("username")), util.StringToInt(r.FormValue("task_id"))).RowsAffected)

		default:
			locals := make(map[string]interface{})
			locals["rolename"] = MTask.ShowCurrutTokenTask(util.GetTypeInt(sess.Get("user_id")))
			util.RenderHtml(w, "showtask", locals)
		}

	}

}
