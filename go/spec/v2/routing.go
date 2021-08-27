package v2

import (
	"encoding/json"
	"github.com/husobee/vestigo"
	log "github.com/sirupsen/logrus"
	"net/http"
)

func checkErrors(params *ParamsParser, w http.ResponseWriter, operationId string, url string) bool {
	if len(params.Errors) > 0 {
		w.WriteHeader(400)
		log.WithFields(log.Fields{
			"operationId": operationId,
			"url":         url,
		}).Warn(params.Errors)
		return false
	}
	return true
}

func checkOperationErrors(err error, w http.ResponseWriter, operationId string, url string) bool {
	if err != nil {
		w.WriteHeader(500)
		log.WithFields(log.Fields{
			"operationId": operationId,
			"url":         url,
		}).Error(err.Error())
		return false
	}
	return true
}

func AddEchoRoutes(router *vestigo.Router, echoService IEchoService) {
	router.Post("/v2/echo/body", func(w http.ResponseWriter, r *http.Request) {
		operationId := "v2.http.echo.echo_body"
		url := "/v2/echo/body"
		log.WithFields(log.Fields{
			"operationId": operationId,
			"url":         url,
		}).Info("Running the Post Method")

		var body Message
		json.NewDecoder(r.Body).Decode(&body)
		response, err := echoService.EchoBody(&body)
		if !checkOperationErrors(err, w, operationId, url) {
			return
		}
		if response.Ok != nil {
			w.WriteHeader(200)
			json.NewEncoder(w).Encode(response.Ok)
			log.WithFields(log.Fields{
				"json": *response.Ok,
			}).Info("Status code: 200")
			log.Info("Completing the Post Method")
			return
		}
		log.Info("Completing the Post Method")
	})
}
