package controllers

import (
	UserM "github.com/freeddser/ddser/models"
	"github.com/freeddser/ddser/util"
	"net/http"
	"time"
)

func UserSignout(w http.ResponseWriter, r *http.Request) {
	sess := GlobalSessions.SessionStart(w, r)
	sess.Delete("username")
	http.Redirect(w, r, "/", 302)

}

func UserLogin(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		sess := GlobalSessions.SessionStart(w, r)
		sess.Get("yzm")
		locals := make(map[string]interface{})
		util.RenderHtml(w, "userlogin", locals)
	}

	if r.Method == "POST" {
		r.ParseForm()
		locals := make(map[string]interface{})
		sess := GlobalSessions.SessionStart(w, r)
		if r.FormValue("yzm") != sess.Get("yzm") {
			locals["loginstatus"] = "Auth Code Error"
			util.RenderHtml(w, "userlogin", locals)
		} else {
			if UserM.CheckUser(r.FormValue("username"), r.FormValue("password")) {
				createtime := sess.Get("createtime")
				if createtime == nil {
					sess.Set("createtime", time.Now().Unix())
				} else if (createtime.(int64) + 360) < (time.Now().Unix()) {
					GlobalSessions.SessionDestroy(w, r)
					sess = GlobalSessions.SessionStart(w, r)
				}
				ct := sess.Get("username")
				if ct == nil {
					sess.Set("username", r.FormValue("username"))
					sess.Set("password", util.GetMd5Str(r.FormValue("password")))
					sess.Set("user_id", UserM.GetUserId(util.GetTypeString(sess.Get("username"))))
					http.Redirect(w, r, "/admin/manager", 302)
				}
				http.Redirect(w, r, "/admin/manager", 302)
			} else {
				locals["loginstatus"] = "username or password wrong"
				util.RenderHtml(w, "userlogin", locals)
			}
		}

	}

}
