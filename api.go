// -*- mode: go; tab-width: 2; indent-tabs-mode: 1; st-rulers: [70] -*-
// vim: ts=4 sw=4 ft=lua noet
//--------------------------------------------------------------------
// @author Daniel Barney <daniel@nanobox.io>
// @copyright 2015, Pagoda Box Inc.
// @doc
//
// @end
// Created :   12 August 2015 by Daniel Barney <daniel@nanobox.io>
//--------------------------------------------------------------------
package api

import (
	"net/http"
	"reflect"
	"runtime"
	"strings"

	"github.com/gorilla/pat"
	"github.com/pagodabox/golang-hatchet"
)

var (
	Router = pat.New()
	Name   = "UNKNOWN"
	Logger = hatchet.Logger(hatchet.DevNullLogger{})

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
