package http

import (
	"context"
	"fmt"
	core_http "net/http"
	"time"

	loggergo "github.com/nextmillenniummedia/logger-go"
	"github.com/samber/do/v2"
)

type Http struct {
	config Config
	logger loggergo.ILogger
	server *core_http.Server
	cw     *ConnectionWatcher
	stop   chan int
}

func ProvideHttp(i do.Injector) (inst *Http, err error) {
	config, err := do.Invoke[Config](i)
	if err != nil {
		return
	}
	logger := do.MustInvokeAs[loggergo.ILogger](i)
	inst = NewHttp(config, logger)
	err = inst.Start()
	return
}

func NewHttp(config Config, logger loggergo.ILogger) *Http {
	host := fmt.Sprintf(":%d", config.Port)
	cw := NewConnectionWatcher()
	server := &core_http.Server{
		Addr:      host,
		ConnState: cw.OnStateChange,
	}
	return &Http{
		config: config,
		server: server,
		logger: logger.Clone().From("http"),
		cw:     cw,
		stop:   make(chan int),
	}
}

func (h *Http) Start() error {
	go func() {
		h.logger.Info("Listen", "port", h.config.Port)
		err := h.server.ListenAndServe()
		if err != nil {
			time.Sleep(h.config.ShutdownAfterDelay)
			h.logger.Error("Shutdown final",
				"message", err.Error(),
				"connections", h.cw.Count(),
				"delay", h.config.ShutdownAfterDelay,
			)
			h.stop <- 1
		}
	}()
	return nil
}

func (h *Http) Shutdown(ctx context.Context) error {
	h.logger.Info("Stop")
	return nil
}
