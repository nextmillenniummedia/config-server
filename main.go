package main

import (
	"config-server/app"
	"config-server/app/commands"
	"config-server/app/http"
	"config-server/app/log"
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/samber/do/v2"
)

func main() {
	injector := do.New()

	do.Provide(injector, log.ProvideConfig)
	do.Provide(injector, log.ProvideLogger)

	do.Provide(injector, commands.ProvideCommands)

	do.Provide(injector, http.ProvideConfig)
	do.Provide(injector, http.ProvideRoutes)
	do.Provide(injector, http.ProvideHttp)

	do.Provide(injector, app.ProvideApp)

	_, err := do.Invoke[*app.App](injector)
	if err != nil {
		panic(err)
	}

	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)

	<-c
	fmt.Println() // For new line before stopping
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	injector.ShutdownWithContext(ctx)
}
