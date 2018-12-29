package middleware

import (
	"fmt"
	"github.com/freeddser/ddser/controllers"
	"github.com/freeddser/ddser/util"
	"log"
	"net/http"
	"runtime/debug"
	"time"
)

type Middleware func(http.HandlerFunc, ...interface{}) http.HandlerFunc

//1 2 3--> 3 2 1
func Chain(f http.HandlerFunc, middlewares ...Middleware) http.HandlerFunc {
	for _, m := range middlewares {
		f = m(f)
	}
	return f
}

func RecoverHandler() Middleware {
	middleware := func(next http.HandlerFunc, args ...interface{}) http.HandlerFunc {
		handler := func(w http.ResponseWriter, r *http.Request) {
			e, ok := recover().(error)
			if ok {
				http.Error(w, e.Error(), http.StatusInternalServerError)
				w.WriteHeader(http.StatusInternalServerError)
				// logging
				log.Println("WARN: panic fired in %v.panic - %v", next, e)
				log.Println(string(debug.Stack()))
			}
			next(w, r)
		}
		return handler
	}
	return middleware
}

func AuthHandler() Middleware {
	middleware := func(next http.HandlerFunc, args ...interface{}) http.HandlerFunc {
		handler := func(w http.ResponseWriter, r *http.Request) {
			sess := controllers.GlobalSessions.SessionStart(w, r)
			ct := sess.Get("username")
			if ct == nil {
				util.Loger("user not login")
				if r.RequestURI != "/login" {
					http.Redirect(w, r, "/login", 302)
				}
			} else {
				util.Loger("user---- login")
				if r.RequestURI == "/login" {
					http.Redirect(w, r, "/admin/manager", 302)
				}
			}
			next(w, r)
		}
		return handler
	}
	return middleware
}

func Timer() Middleware {
	middleware := func(next http.HandlerFunc, args ...interface{}) http.HandlerFunc {
		handler := func(w http.ResponseWriter, r *http.Request) {
			defer func(begin time.Time) {
				fmt.Println(r.Method, r.RequestURI, r.Proto, "----->", time.Since(begin))
			}(time.Now())
			next(w, r)
		}
		return handler
	}
	return middleware
}

func LogerClient() Middleware {
	middleware := func(next http.HandlerFunc, args ...interface{}) http.HandlerFunc {
		handler := func(w http.ResponseWriter, r *http.Request) {
			//fmt.Println("TLS:", r.TLS)
			//fmt.Println("server:", r.Host, "client:", r.RemoteAddr)
			//fmt.Println(r.Method, r.RequestURI, r.Proto)
			//for key, value := range r.Header {
			//	fmt.Println(key, ":", value)
			//}
			//if r.Method == "POST" {
			//	r.ParseForm()
			//	fmt.Println("POST")
			//	fmt.Println(r.PostForm)
			//
			//}
			next(w, r)
		}
		return handler
	}
	return middleware
}

func AuthURL() Middleware {
	middleware := func(next http.HandlerFunc, args ...interface{}) http.HandlerFunc {
		handler := func(w http.ResponseWriter, r *http.Request) {
			sess := controllers.GlobalSessions.SessionStart(w, r)
			if !isKeyInMap(util.GetTypeMap(sess.Get("url_access")), r.URL.Path) {
				// w.WriteHeader(http.StatusMethodNotAllowed)
				http.Redirect(w, r, "/err/405", 302)
			}
			util.Loger("can access")
			next(w, r)
		}
		return handler
	}
	return middleware
}

func isKeyInMap(data map[string]interface{}, key string) bool {
	if _, ok := data[key]; ok {
		return true
	} else {
		return false
	}
}
