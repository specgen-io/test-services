package spec

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
	logEchoBody := log.Fields{"operationId": "echo.echo_body", "method": "POST", "url": "/echo/body"}
	router.Post("/echo/body", func(w http.ResponseWriter, r *http.Request) {
		log.WithFields(logEchoBody).Info("Received request")
		var body Message
		err := json.NewDecoder(r.Body).Decode(&body)
		if err != nil {
			log.Warnf("Decoding body JSON failed: %s", err.Error())
			w.WriteHeader(400)
			log.WithFields(logEchoBody).WithField("status", 400).Info("Completed request")
			return
		}
		response, err := echoService.EchoBody(&body)
		if err != nil {
			log.Errorf("Error returned from service implementation: %s", err.Error())
			w.WriteHeader(500)
			log.WithFields(logEchoBody).WithField("status", 500).Info("Completed request")
			return
		}
		if response.Ok != nil {
			w.WriteHeader(200)
			json.NewEncoder(w).Encode(response.Ok)
			log.WithFields(logEchoBody).WithField("status", 200).Info("Completed request")
			return
		}
		log.Error("Failed request - operation implementation did not provide result")
		w.WriteHeader(500)
		log.WithFields(logEchoBody).WithField("status", 500).Info("Completed request")
	})

	logEchoQuery := log.Fields{"operationId": "echo.echo_query", "method": "GET", "url": "/echo/query"}
	router.Get("/echo/query", func(w http.ResponseWriter, r *http.Request) {
		log.WithFields(logEchoQuery).Info("Received request")
		queryParams := NewParamsParser(r.URL.Query())
		intQuery := queryParams.Int("int_query")
		stringQuery := queryParams.String("string_query")
		if len(queryParams.Errors) > 0 {
			// TODO: Convert Errors collection to single string message
			log.Warnf("Can't parse query params: %s", queryParams.Errors)
			w.WriteHeader(400)
			log.WithFields(logEchoQuery).WithField("status", 400).Info("Completed request")
			return
		}
		response, err := echoService.EchoQuery(intQuery, stringQuery)
		if err != nil {
			log.Errorf("Error returned from service implementation: %s", err.Error())
			w.WriteHeader(500)
			log.WithFields(logEchoQuery).WithField("status", 500).Info("Completed request")
			return
		}
		if response.Ok != nil {
			w.WriteHeader(200)
			json.NewEncoder(w).Encode(response.Ok)
			log.WithFields(logEchoQuery).WithField("status", 200).Info("Completed request")
			return
		}
		// TODO: Move this to response errors checks
		log.Error("Failed request - operation implementation did not provide result")
		w.WriteHeader(500)
		log.WithFields(logEchoQuery).WithField("status", 500).Info("Completed request")
	})

	logEchoHeader := log.Fields{"operationId": "echo.echo_header", "method": "GET", "url": "/echo/header"}
	router.Get("/echo/header", func(w http.ResponseWriter, r *http.Request) {
		log.WithFields(logEchoHeader).Info("Received request")
		headerParams := NewParamsParser(r.Header)
		intHeader := headerParams.Int("Int-Header")
		stringHeader := headerParams.String("String-Header")
		if len(headerParams.Errors) > 0 {
			// TODO: Convert Errors collection to single string message
			log.Warnf("Can't parse header params: %s", headerParams.Errors)
			w.WriteHeader(400)
			log.WithFields(logEchoHeader).WithField("status", 400).Warn(headerParams.Errors)
			return
		}
		response, err := echoService.EchoHeader(intHeader, stringHeader)
		if err != nil {
			log.Errorf("Error returned from service implementation: %s", err.Error())
			w.WriteHeader(500)
			log.WithFields(logEchoHeader).WithField("status", 500).Error(err.Error())
			return
		}
		if response.Ok != nil {
			json.NewEncoder(w).Encode(response.Ok)
			w.WriteHeader(200)
			log.WithFields(logEchoHeader).WithField("status", 200).Info("Completed request")
			return
		}
		log.Error("Failed request - operation implementation did not provide result")
		w.WriteHeader(500)
		log.WithFields(logEchoHeader).WithField("status", 500).Info("Completed request")
	})
	router.SetCors("/echo/header", &vestigo.CorsAccessControl{
		AllowHeaders: []string{"Int-Header", "String-Header"},
	})

	router.Get("/echo/url_params/:int_url/:string_url", func(w http.ResponseWriter, r *http.Request) {
		operationId := "http.echo.echo_url_params"
		url := "/echo/url_params/:int_url/:string_url"
		log.WithFields(log.Fields{
			"operationId": operationId,
			"url":         url,
		}).Info("Running the Get Method")

		urlParams := NewParamsParser(r.URL.Query())
		intUrl := urlParams.Int(":int_url")
		stringUrl := urlParams.String(":string_url")
		if !checkErrors(urlParams, w, operationId, url) {
			return
		}
		response, err := echoService.EchoUrlParams(intUrl, stringUrl)
		if !checkOperationErrors(err, w, operationId, url) {
			return
		}
		if response.Ok != nil {
			w.WriteHeader(200)
			json.NewEncoder(w).Encode(response.Ok)
			log.WithFields(log.Fields{
				"json": *response.Ok,
			}).Info("Status code: 200")
			log.Info("Completing the Get Method")
			return
		}
		log.Info("Completing the Get Method")
	})
}

func AddCheckRoutes(router *vestigo.Router, checkService ICheckService) {
	router.Get("/check/empty", func(w http.ResponseWriter, r *http.Request) {
		operationId := "http.check.check_empty"
		url := "/check/empty"
		log.WithFields(log.Fields{
			"operationId": operationId,
			"url":         url,
		}).Info("Running the Get Method")

		response, err := checkService.CheckEmpty()
		if !checkOperationErrors(err, w, operationId, url) {
			return
		}
		if response.Ok != nil {
			w.WriteHeader(200)
			json.NewEncoder(w).Encode(response.Ok)
			log.WithFields(log.Fields{
				"json": *response.Ok,
			}).Info("Status code: 200")
			log.Info("Completing the Get Method")
			return
		}
		log.Info("Completing the Get Method")
	})
	router.Get("/check/query", func(w http.ResponseWriter, r *http.Request) {
		operationId := "http.check.check_query"
		url := "/check/query"
		log.WithFields(log.Fields{
			"operationId": operationId,
			"url":         url,
		}).Info("Running the Get Method")

		queryParams := NewParamsParser(r.URL.Query())
		pString := queryParams.String("p_string")
		pStringOpt := queryParams.StringNullable("p_string_opt")
		pStringArray := queryParams.StringArray("p_string_array")
		pDate := queryParams.Date("p_date")
		pDateArray := queryParams.DateArray("p_date_array")
		pDatetime := queryParams.DateTime("p_datetime")
		pInt := queryParams.Int("p_int")
		pLong := queryParams.Int64("p_long")
		pDecimal := queryParams.Decimal("p_decimal")
		pEnum := Choice(queryParams.StringEnum("p_enum", ChoiceValuesStrings))
		pStringDefaulted := queryParams.StringDefaulted("p_string_defaulted", "the default value")
		if !checkErrors(queryParams, w, operationId, url) {
			return
		}
		response, err := checkService.CheckQuery(pString, pStringOpt, pStringArray, pDate, pDateArray, pDatetime, pInt, pLong, pDecimal, pEnum, pStringDefaulted)
		if !checkOperationErrors(err, w, operationId, url) {
			return
		}
		if response.Ok != nil {
			w.WriteHeader(200)
			json.NewEncoder(w).Encode(response.Ok)
			log.WithFields(log.Fields{
				"json": *response.Ok,
			}).Info("Status code: 200")
			log.Info("Completing the Get Method")
			return
		}
		log.Info("Completing the Get Method")
	})
	router.Get("/check/url_params/:int_url/:string_url/:float_url/:bool_url/:uuid_url/:decimal_url/:date_url", func(w http.ResponseWriter, r *http.Request) {
		operationId := "http.check.check_url_params"
		url := "/check/url_params/:int_url/:string_url/:float_url/:bool_url/:uuid_url/:decimal_url/:date_url"
		log.WithFields(log.Fields{
			"operationId": operationId,
			"url":         url,
		}).Info("Running the Get Method")

		urlParams := NewParamsParser(r.URL.Query())
		intUrl := urlParams.Int64(":int_url")
		stringUrl := urlParams.String(":string_url")
		floatUrl := urlParams.Float32(":float_url")
		boolUrl := urlParams.Bool(":bool_url")
		uuidUrl := urlParams.Uuid(":uuid_url")
		decimalUrl := urlParams.Decimal(":decimal_url")
		dateUrl := urlParams.Date(":date_url")
		if !checkErrors(urlParams, w, operationId, url) {
			return
		}
		response, err := checkService.CheckUrlParams(intUrl, stringUrl, floatUrl, boolUrl, uuidUrl, decimalUrl, dateUrl)
		if !checkOperationErrors(err, w, operationId, url) {
			return
		}
		if response.Ok != nil {
			w.WriteHeader(200)
			json.NewEncoder(w).Encode(response.Ok)
			log.WithFields(log.Fields{
				"json": *response.Ok,
			}).Info("Status code: 200")
			log.Info("Completing the Get Method")
			return
		}
		log.Info("Completing the Get Method")
	})
	router.Get("/check/forbidden", func(w http.ResponseWriter, r *http.Request) {
		operationId := "http.check.check_forbidden"
		url := "/check/forbidden"
		log.WithFields(log.Fields{
			"operationId": operationId,
			"url":         url,
		}).Info("Running the Get Method")

		response, err := checkService.CheckForbidden()
		if !checkOperationErrors(err, w, operationId, url) {
			return
		}
		if response.Ok != nil {
			w.WriteHeader(200)
			json.NewEncoder(w).Encode(response.Ok)
			log.WithFields(log.Fields{
				"json": *response.Ok,
			}).Info("Status code: 200")
			log.Info("Completing the Get Method")
			return
		}
		if response.Forbidden != nil {
			w.WriteHeader(403)
			log.Info("Status code: 403")
			log.Info("Completing the Get Method")
			return
		}
		log.Info("Completing the Get Method")
	})
}
