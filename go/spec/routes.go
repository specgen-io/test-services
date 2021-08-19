package spec

import (
	"github.com/husobee/vestigo"
	"github.com/specgen-io/test-service/go/spec/v2"
)

func AddRoutes(router *vestigo.Router) {
	v2.AddEchoRoutes(router, &v2.EchoService{})
	AddEchoRoutes(router, &EchoService{})
	AddCheckRoutes(router, &CheckService{})
}
