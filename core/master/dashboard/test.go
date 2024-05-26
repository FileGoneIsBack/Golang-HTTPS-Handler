package dashboard

import (
	"api/core/models"
	"api/core/models/functions"
	"api/core/models/server"
	"net/http"
)

func init() {
	Route.NewSub(server.NewRoute("/test", func(w http.ResponseWriter, r *http.Request) {
		type Page struct {
			Name, Title, Vers       string
		}
		functions.Render(Page{
			Name:         "Twilight",
			Title:        "Dashboard",
			Vers:         models.Config.Vers,
		}, w, "test", "test.html")
	}))
}
