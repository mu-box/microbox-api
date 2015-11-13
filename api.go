package api

import (
	"github.com/gorilla/pat"
	"github.com/jcelliott/lumber"
	"net/http"
	"reflect"
	"runtime"
	"strings"
)

var (
	Router = pat.New()
	Name   = "UNKNOWN"
	Logger lumber.Logger

	// can be set by the user of this module
	User interface{}
)

func init() {
	Router.Get("/ping", TraceRequest(pongRoute))
}

// pong to a ping.
func pongRoute(res http.ResponseWriter, req *http.Request) {
	res.WriteHeader(200)
	res.Write([]byte("pong"))
}

// Start up the api and begin responding to requests. Blocking.
func Start(address string) error {
	return http.ListenAndServe(address, Router)
}

// Traces all routes going through the api.
func TraceRequest(fn http.HandlerFunc) http.HandlerFunc {
	return func(res http.ResponseWriter, req *http.Request) {

		v := reflect.ValueOf(fn)
		if rf := runtime.FuncForPC(v.Pointer()); rf != nil {
			names := strings.Split(rf.Name(), "/")
			Logger.Info("[NA-%v] %v %v %v", Name, names[len(names)-1], req.URL.Path, req.RemoteAddr)
		}

		fn(res, req)

	}
}
