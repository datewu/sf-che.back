package main

import (
	"context"
	"flag"

	"github.com/datewu/sf-che/cmd/api"

	"github.com/datewu/gtea"
	"github.com/datewu/gtea/jsonlog"
)

var (
	version   = "1.0.0"
	buildTime string
)

func main() {
	var (
		port int
		env  string
	)
	flag.IntVar(&port, "port", 4000, "API server port")
	flag.StringVar(&env, "env", "development", "Environment (development|staging|production)")

	flag.Parse()
	cfg := &gtea.Config{
		Port:     port,
		Env:      env,
		Metrics:  true,
		LogLevel: jsonlog.LevelInfo,
	}
	app := gtea.NewApp(cfg)
	app.Logger.Info("APP Starting",
		map[string]string{
			"version":   version,
			"gitCommit": buildTime,
			"mode":      env,
		})
	app.AddMetaData("version", version)

	ctx := context.Background()
	app.Serve(ctx, api.New(app))
}
