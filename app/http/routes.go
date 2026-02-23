package http

import (
	"config-server/app/commands"
	"config-server/app/utils"
	"net/http"

	"github.com/gorilla/mux"
	loggergo "github.com/nextmillenniummedia/logger-go"
	"github.com/samber/do/v2"
)

func ProvideRoutes(i do.Injector) (routes *Routes, err error) {
	commands := do.MustInvoke[*commands.Commands](i)
	logger := do.MustInvokeAs[loggergo.ILogger](i)
	return &Routes{
		commands: commands,
		logger:   logger,
	}, nil
}

type Routes struct {
	commands *commands.Commands
	logger   loggergo.ILogger
}

func (r *Routes) GetRoutes() *mux.Router {
	route := mux.NewRouter()

	// r.Use(log.SetToRequest)
	// r.Use(log.Http)

	// r.HandleFunc("/status", status(app.Status)).Methods("GET")
	route.HandleFunc("/command/{name}", func(w http.ResponseWriter, req *http.Request) {
		vars := mux.Vars(req)
		name := vars["name"]
		params, err := utils.ReadPayload(req)
		if err != nil {
			utils.SendHttpError(w, err, r.logger)
			return
		}
		response, err := r.commands.Execute(name, params)
		if err != nil {
			utils.SendHttpError(w, err, r.logger)
			return
		}
		utils.SendHttpResponse(w, response, r.logger)
	}).Methods("POST")

	return route
}
