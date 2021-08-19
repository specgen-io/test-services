//go:generate specgen service-go --spec-file ./../spec.yaml --generate-path .

package main

import (
	"flag"
	"fmt"
	"github.com/husobee/vestigo"
	"github.com/specgen-io/test-service/go/spec"
	"log"
	"net/http"
)

func main() {
	port := flag.String("port", "8081", "port number")
	flag.Parse()

	router := vestigo.NewRouter()

	router.SetGlobalCors(&vestigo.CorsAccessControl{
		AllowOrigin: []string{"*", "*"},
	})

	spec.AddRoutes(router)

	fmt.Println("Starting service on port: " + *port)
	log.Fatal(http.ListenAndServe(":"+*port, router))
}
