package controllers

import (
	"fmt"
	UserM "github.com/freeddser/ddser/models"
	"github.com/freeddser/ddser/util"
	"net/http"
)

func UserProfile(w http.ResponseWriter, r *http.Request) {

	sess := GlobalSessions.SessionStart(w, r)
	if r.Method == "GET" {
		locals := make(map[string]interface{})
		locals["username"] = sess.Get("username")
		util.RenderHtml(w, "profile", locals)
	}
	if r.Method == "POST" {

		r.ParseForm()
		switch r.FormValue("type") {
		//auth token
		case "updatepass":
			oldpass := r.PostFormValue("oldpass")
			password1 := r.PostFormValue("password1")
			password2 := r.PostFormValue("password2")
			if util.GetMd5Str(oldpass) == util.GetTypeString(sess.Get("password")) {
				fmt.Println("ok,change pass")
				if password1 == password2 {
					w.Header().Set("Content-Type", "application/json")
					fmt.Fprintln(w, UserM.MChangePassword(util.GetTypeInt(sess.Get("user_id")), password1).RowsAffected)
					sess.Set("password", util.GetMd5Str(password1))
				}
			} else {
				fmt.Println("faild, oldpass wrong")
			}

		default:
			return
		}

	}

}
