package main

import (
	"fmt"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/sky-uk/support-bot/rota"
	"log"
	"net/http"
)

var myTeam = rota.NewTeam("core-infrastructure")

func serve() {

	http.HandleFunc("/members", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			// Would be useful to provide full support details of team members
			_, _ = fmt.Fprint(w, myTeam.List())
		}
	})

	http.Handle("/metrics", promhttp.Handler())

	http.HandleFunc("/support/next", func(writer http.ResponseWriter, request *http.Request) {
		if request.Method == http.MethodGet {
			_, _ = fmt.Fprint(writer, rota.Next(myTeam))
		}
	})

	http.HandleFunc("/support/override", func(writer http.ResponseWriter, request *http.Request) {
		// fetch the name to be used and fix it
	})

	http.HandleFunc("/team/member", func(writer http.ResponseWriter, request *http.Request) {
		switch request.Method {
		case http.MethodPost:
			// Add new team member
		case http.MethodDelete:
			//	Remove member from support rota
		}
	})

	log.Fatal(http.ListenAndServe("localhost:9090", nil))
}

func main() {
	serve()
}
