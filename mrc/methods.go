package mrc

import (
	"net/http"

	"github.com/ecsaas/main-router-config/DEFINE_VARIABLES/ROUTERS/METHODS"
	"github.com/gorilla/mux"
)

func Delete(
	mux *mux.Router,
	PathPrefix string,
	HandlerFunc func(http.ResponseWriter, *http.Request),
	Path string,
) {
	mux.PathPrefix(PathPrefix).
		Path(Path).
		HandlerFunc(HandlerFunc).
		Methods(METHODS.DELETE)
}

func Get(
	mux *mux.Router,
	PathPrefix string,
	HandlerFunc func(http.ResponseWriter, *http.Request),
	Path string,
) {
	mux.PathPrefix(PathPrefix).
		Path(Path).
		HandlerFunc(HandlerFunc).
		Methods(METHODS.GET)
}

func Options(
	mux *mux.Router,
	PathPrefix string,
	HandlerFunc func(http.ResponseWriter, *http.Request),
	Path string,
) {
	mux.PathPrefix(PathPrefix).
		Path(Path).
		HandlerFunc(HandlerFunc).
		Methods(METHODS.OPTIONS)
}

func Patch(
	mux *mux.Router,
	PathPrefix string,
	HandlerFunc func(http.ResponseWriter, *http.Request),
	Path string,
) {
	mux.PathPrefix(PathPrefix).
		Path(Path).
		HandlerFunc(HandlerFunc).
		Methods(METHODS.PATCH)
}

func Post(
	mux *mux.Router,
	PathPrefix string,
	HandlerFunc func(http.ResponseWriter, *http.Request),
	Path string,
) {
	mux.PathPrefix(PathPrefix).
		Path(Path).
		HandlerFunc(HandlerFunc).
		Methods(METHODS.POST)
}

func Put(
	mux *mux.Router,
	PathPrefix string,
	HandlerFunc func(http.ResponseWriter, *http.Request),
	Path string,
) {
	mux.PathPrefix(PathPrefix).
		Path(Path).
		HandlerFunc(HandlerFunc).
		Methods(METHODS.PUT)
}

func Request(
	mux *mux.Router,
	PathPrefix string,
	HandlerFunc func(http.ResponseWriter, *http.Request),
	Path string,
) {
	mux.PathPrefix(PathPrefix).
		Path(Path).
		HandlerFunc(HandlerFunc)
}
