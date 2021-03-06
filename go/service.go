//go:generate specgen-go service-go --spec-file ./../spec.yaml --module-name test-service --generate-path ./spec --services-path ./services --swagger-path docs/swagger.yaml

package main

import (
	"flag"
	"net/http"
	"test-service/services"
	"test-service/services/v2"
	"test-service/spec"

	"github.com/husobee/vestigo"
	"github.com/shopspring/decimal"
	log "github.com/sirupsen/logrus"
)

func main() {
	port := flag.String("port", "8081", "port number")
	flag.Parse()

	decimal.MarshalJSONWithoutQuotes = true

	router := vestigo.NewRouter()

	router.SetGlobalCors(&vestigo.CorsAccessControl{
		AllowOrigin: []string{"*", "*"},
	})

	echoServiceV2 := &v2.EchoService{}
	echoService := &services.EchoService{}
	checkService := &services.CheckService{}

	spec.AddRoutes(router, echoServiceV2, echoService, checkService)

	router.Get("/docs/*", http.StripPrefix("/docs/", http.FileServer(http.Dir("docs"))).ServeHTTP)

	log.Infof("Starting service on port: %s", *port)
	log.Fatal(http.ListenAndServe(":"+*port, router))
}
