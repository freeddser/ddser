package controllers

import (
	"github.com/freeddser/ddser/session"
	_ "github.com/freeddser/ddser/session/providers/memory"

	"github.com/freeddser/ddser/util"
	"net/http"
	"os"
)

var GlobalSessions *session.Manager

func InitSession() {
	//session 3600s 1h
	GlobalSessions, _ = session.NewManager("memory", "ddsersessionid", 36000)
	go GlobalSessions.GC()

}

func StaticDirHandler(mux *http.ServeMux, prefix string, staticDir string, flags int) {
	mux.HandleFunc(prefix, func(w http.ResponseWriter, r *http.Request) {
		file := staticDir + r.URL.Path[len(prefix)-1:]
		if (flags & 0x0001) == 0 {
			fi, err := os.Stat(file)
			if err != nil || fi.IsDir() {
				http.NotFound(w, r)
				return
			}
		}
		http.ServeFile(w, r, file)
	})
}

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/login", 302)

}

func DefaultHandler(w http.ResponseWriter, r *http.Request) {
	return
}

func Err405(w http.ResponseWriter, r *http.Request) {
	locals := make(map[string]interface{})
	locals["error"] = "Url not allow to access,Please contact to Admin"
	util.RenderHtml(w, "err", locals)

}

func Err404(w http.ResponseWriter, r *http.Request) {
	locals := make(map[string]interface{})
	locals["error"] = "url 404!"
	util.RenderHtml(w, "err", locals)

}
