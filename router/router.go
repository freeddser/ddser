package router

import (
	c "github.com/freeddser/ddser/controllers"
	"github.com/freeddser/ddser/middleware"
	"github.com/freeddser/ddser/util/yzm"

	"net/http"
)

var recoverHandler = middleware.RecoverHandler()
var auth = middleware.AuthHandler()
var authURL = middleware.AuthURL()
var log = middleware.LogerClient()
var timer = middleware.Timer()

func NewRouter() *http.ServeMux {
	mux := http.NewServeMux()
	c.StaticDirHandler(mux, "/static/", "./static", 0)
	mux.HandleFunc("/pic", yzm.Pic)
	//user login
	mux.HandleFunc("/", middleware.Chain(auth(c.IndexHandler), log, timer, recoverHandler))
	//user login
	mux.HandleFunc("/login", middleware.Chain(auth(c.UserLogin), log, timer, recoverHandler))
	mux.HandleFunc("/profile", middleware.Chain(auth(c.UserProfile), log, timer, recoverHandler))
	mux.HandleFunc("/signout", middleware.Chain(log(c.UserSignout), log, timer, recoverHandler))
	mux.HandleFunc("/admin/default", middleware.Chain(log(c.DefaultHandler), timer, recoverHandler))
	mux.HandleFunc("/admin/manager", middleware.Chain(auth(c.UserManager), log, timer, recoverHandler))

	//user manager
	mux.HandleFunc("/admin/manager/user", middleware.Chain(authURL(c.UserPrivilegesManager), auth, log, timer, recoverHandler))
	mux.HandleFunc("/admin/manager/role", middleware.Chain(authURL(c.RoleManager), auth, log, timer, recoverHandler))

	mux.HandleFunc("/admin/manager/userrole", middleware.Chain(authURL(c.UserRoleManager), auth, log, timer, recoverHandler))

	mux.HandleFunc("/admin/manager/roleresource", middleware.Chain(authURL(c.RoleResourceManager), auth, log, timer, recoverHandler))

	mux.HandleFunc("/admin/manager/resource", middleware.Chain(authURL(c.ResourceManager), auth, log, timer, recoverHandler))

	mux.HandleFunc("/admin/manager/roletasktype", middleware.Chain(authURL(c.RoleTaskTypeManager), auth, log, timer, recoverHandler))

	//upload sql file
	mux.HandleFunc("/upload/sql", middleware.Chain(auth(c.UploadSQLFileHandler), log, timer, recoverHandler))

	//err 404
	mux.HandleFunc("/err/404", middleware.Chain(log(c.Err404), timer, recoverHandler))
	//err 405
	mux.HandleFunc("/err/405", middleware.Chain(log(c.Err405), timer, recoverHandler))

	//Task
	mux.HandleFunc("/task/add", middleware.Chain(authURL(c.TaskAddHandler), auth, log, timer, recoverHandler))
	mux.HandleFunc("/task/show", middleware.Chain(authURL(c.TaskShowHandler), auth, log, timer, recoverHandler))
	mux.HandleFunc("/task/tasklog", middleware.Chain(authURL(c.TaskShowLogHandler), auth, log, timer, recoverHandler))

	//task for client
	mux.HandleFunc("/taskclient/show", middleware.Chain(log(c.TaskClientShowHandler), timer, recoverHandler))

	return mux
}
