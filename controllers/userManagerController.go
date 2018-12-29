package controllers

import (
	"fmt"
	UserM "github.com/freeddser/ddser/models"
	"github.com/freeddser/ddser/util"
	"net/http"
)

func UserManager(w http.ResponseWriter, r *http.Request) {
	sess := GlobalSessions.SessionStart(w, r)
	if r.Method == "GET" {
		locals := make(map[string]interface{})
		locals["username"] = sess.Get("username")
		util.RenderHtml(w, "usermanager", locals)
	}
	if r.Method == "POST" {
		//get user privileges to session
		userprivileges := UserM.GetUserRoleInfos(util.GetTypeString(sess.Get("username")))
		sess.Set("user_privileges", userprivileges)
		//set user url privileges
		sess.Set("url_access", UserM.GetUrlAccess(userprivileges))
		//response json
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintf(w, util.GetTypeString(sess.Get("user_privileges")))
	}
}
