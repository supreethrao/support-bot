package main

import (
	"fmt"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/sky-uk/support-bot/rota"
	"log"
	"net/http"
	"path/filepath"
	"runtime"
)

var myTeam = rota.NewTeam(teamMembersFilePath())

func serve() {
	handleMembersList()
	http.Handle("/metrics", promhttp.Handler())
	log.Fatal(http.ListenAndServe("localhost:9090", nil))
}

func handleMembersList() {
	http.HandleFunc("/members", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			_, err := fmt.Fprint(w, myTeam.List())
			if err != nil {
				log.Fatal(err)
			}
		}
	})
}

func main() {
	serve()
}

func teamMembersFilePath() string {
	_, filename, _, _ := runtime.Caller(1)
	basePath := filepath.Dir(filename)
	return basePath + "/core-team-members.yml"
}
