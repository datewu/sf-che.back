package api

import (
	"net/http"

	"github.com/datewu/gtea"
	"github.com/datewu/gtea/handler/sse"
	"github.com/datewu/gtea/router"
)

func New(app *gtea.App) http.Handler {
	conf := router.DefaultConf()
	conf.CORS.TrustedOrigins = []string{"https://sf-che.com", "https://www.sf-che.com"}
	r, _ := router.NewRoutesGroup(conf)
	r.Get("/", hw)
	g := r.Group("/v1")
	setRoutes(app, g)
	return g
}

func setRoutes(_ *gtea.App, g *router.RoutesGroup) {
	g.Get("/sse_stream", sse.Demo)
}
