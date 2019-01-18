package main

import (
	"bytes"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/sky-uk/support-bot/localdb"
	"github.com/sky-uk/support-bot/rota"
	"github.com/sky-uk/support-bot/rota/slackhandler"
	"log"
	"net/http"
)

var myTeam = rota.NewTeam("core-infrastructure")

func serve() {
	router := httprouter.New()

	router.GET("/members", func(writer http.ResponseWriter, request *http.Request, _ httprouter.Params) {
		_, _ = fmt.Fprint(writer, myTeam.SupportHistoryForTeam())
	})

	router.POST("/members", func(writer http.ResponseWriter, request *http.Request, _ httprouter.Params) {
		_, _ = fmt.Fprint(writer, myTeam.SupportHistoryForTeam())
	})

	router.DELETE("/members/:name", func(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
		if err := myTeam.Remove(params.ByName("name")); err != nil {
			_, _ = fmt.Fprint(writer, )
		}
	})

	router.GET("/support/next", func(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
		_, _ = fmt.Fprint(writer, slackhandler.MemberOptions(myTeam.SupportHistoryForTeam()))

	})

	http.Handle("/metrics", promhttp.Handler())

	http.HandleFunc("/support/confirm", func(writer http.ResponseWriter, request *http.Request) {
		buf := new(bytes.Buffer)
		_, _ = buf.ReadFrom(request.Body)
	})

	log.Fatal(http.ListenAndServe(":9090", router))
}

func main() {
	myTeam.Reset()
	defer localdb.Close()
	serve()
}
