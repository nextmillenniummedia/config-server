package app

import (
	"config-server/app/http"

	"github.com/samber/do/v2"
)

type App struct {
	http *http.Http
}

func ProvideApp(i do.Injector) (app *App, err error) {
	http := do.MustInvoke[*http.Http](i)
	return &App{
		http: http,
	}, nil
}
