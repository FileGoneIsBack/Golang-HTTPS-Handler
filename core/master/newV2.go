package master

import (
	"api/core/master/dashboard"
	"api/core/master/landing"
	"api/core/models"
	"api/core/models/server"
	"net/http"
)

var (
	Service *server.Server = server.NewServer(&server.Config{
		Addr:   "0.0.0.0:80",
		Secure: models.Config.Secure,
		Cert:   models.Config.Cert,
		Key:    models.Config.Cert,
	})
	Route  *server.Route   = server.NewSubRouter("")
	Assets *server.Handler = server.NewHandler("/_assets/", http.StripPrefix("/_assets", http.FileServer(http.Dir("assets/branding"))))
)

func NewV2() {
	Route.NewSubs(
		dashboard.Route,
		landing.Route,
	)
	Service.AddRoute(Route)
	Service.AddHandler(Assets)

	Service.AddRoute(server.NewRoute("", nil))
	if err := Service.ListenAndServe(); err != nil {
		panic(err)
	}
}