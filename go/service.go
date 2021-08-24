//go:generate specgen service-go --spec-file ./../spec.yaml --module-name test-service  --generate-path .

package main

import (
	"flag"
	"fmt"
	"github.com/husobee/vestigo"
	"log"
	"net/http"
	"test-service/spec"
	"test-service/spec/v2"
)

func main() {
	port := flag.String("port", "8081", "port number")
	flag.Parse()

	router := vestigo.NewRouter()

	router.SetGlobalCors(&vestigo.CorsAccessControl{
		AllowOrigin: []string{"*", "*"},
	})

	echoServiceV2 := &v2.EchoService{}
	echoService := &spec.EchoService{}
	checkService := &spec.CheckService{}

	spec.AddRoutes(router, echoServiceV2, echoService, checkService)

	fmt.Println("Starting service on port: " + *port)
	log.Fatal(http.ListenAndServe(":"+*port, router))
}
