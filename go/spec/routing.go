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
	router.Post("/echo/body", func(w http.ResponseWriter, r *http.Request) {
		operationId := "http.echo.echo_body"
		url := "/echo/body"
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
		if response.Ok == nil {
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
	router.Get("/echo/query", func(w http.ResponseWriter, r *http.Request) {
		operationId := "http.echo.echo_query"
		url := "/echo/query"
		log.WithFields(log.Fields{
			"operationId": operationId,
			"url":         url,
		}).Info("Running the Get Method")

		queryParams := NewParamsParser(r.URL.Query())
		intQuery := queryParams.Int("int_query")
		stringQuery := queryParams.String("string_query")
		if !checkErrors(queryParams, w, operationId, url) {
			return
		}
		response, err := echoService.EchoQuery(intQuery, stringQuery)
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
	router.Get("/echo/header", func(w http.ResponseWriter, r *http.Request) {
		operationId := "http.echo.echo_header"
		url := "/echo/header"
		log.WithFields(log.Fields{
			"operationId": operationId,
			"url":         url,
		}).Info("Running the Get Method")

		headerParams := NewParamsParser(r.Header)
		intHeader := headerParams.Int("Int-Header")
		stringHeader := headerParams.String("String-Header")
		if !checkErrors(headerParams, w, operationId, url) {
			return
		}
		response, err := echoService.EchoHeader(intHeader, stringHeader)
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
