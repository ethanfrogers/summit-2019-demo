package api

import (
	"github.com/gorilla/mux"
	"net/http"
)

type MessageProvider interface {
	Message() string
}

type VersionProvider interface {
	Version() string
}

type WebAPI struct {
	MessageProvider MessageProvider
	VersionProvider VersionProvider
}

func (wapi *WebAPI) Router(r *mux.Router) {
	r.HandleFunc("/hello", wapi.HelloSpinSumHandler).Methods(http.MethodGet)
	r.HandleFunc("/version", wapi.VersionHandler).Methods(http.MethodGet)
}

func (wapi *WebAPI) HelloSpinSumHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(wapi.MessageProvider.Message()))
}

func (wapi *WebAPI) VersionHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(wapi.VersionProvider.Version()))
}