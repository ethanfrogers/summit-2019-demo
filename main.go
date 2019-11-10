package main

import (
	"fmt"
	"github.com/ethanfrogers/summit-2019-demo/pkg/api"
	"github.com/gorilla/mux"
	"net/http"
	"os"
)

type defaultMessageProvider struct {}

func (d *defaultMessageProvider) Message() string {
	message := os.Getenv("HELLO_MESSAGE")

	if message != "" {
		return message
	}

	return "message environment variable unset"
}

type defaultVerisonProvider struct {}

func (d *defaultVerisonProvider) Version() string {
	version := os.Getenv("VERSION")

	if version != "" {
		return version
	}

	return "development"
}

func main() {

	router := mux.NewRouter()
	webApi := api.WebAPI{
		MessageProvider: &defaultMessageProvider{},
		VersionProvider: &defaultVerisonProvider{},
	}
	webApi.Router(router)

	server := http.Server{
		Addr: ":3000",
		Handler: router,
	}

	fmt.Println("server starting on port 3000")
	if err := server.ListenAndServe(); err != http.ErrServerClosed {
		fmt.Errorf("server closed without intervention")
		os.Exit(1)
	}
	fmt.Println("server quitting gracefully")

}