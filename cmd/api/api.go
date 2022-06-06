package api

import (
	"net/http"

	"github.com/datewu/gtea"
	"github.com/datewu/gtea/handler"
	"github.com/datewu/gtea/handler/sse"
)

func New(app *gtea.App) http.Handler {
	conf := handler.DefaultConf()
	conf.CORS.TrustedOrigins = []string{"sf-che.com", "*.sf-che.com"}
	r, _ := handler.NewRouterGroup(conf)
	g := r.Group("")
	setRoutes(app, g)
	return g
}

func setRoutes(_ *gtea.App, g *handler.RouterGroup) {
	g.Get("/sse_stream", sse.Demo)
	g.Get("/ping", handler.HealthCheck)
}
