//go:generate specgen service-go --spec-file ./../spec.yaml --module-name test-service --generate-path ./spec --swagger-path docs/swagger.yaml

package main

import (
	"flag"
	"github.com/husobee/vestigo"
	log "github.com/sirupsen/logrus"
	"net/http"
	"test-service/services"
	"test-service/services/v2"
	"test-service/spec"
)

func main() {
	port := flag.String("port", "8081", "port number")
	flag.Parse()

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
