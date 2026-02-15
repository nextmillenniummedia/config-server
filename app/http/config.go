package http

import (
	"time"

	configgo "github.com/nextmillenniummedia/config-go"
	"github.com/samber/do/v2"
)

type Config struct {
	Port                      int           `config:"doc='http port that the server uses',field=port,default=3000"`
	ShutdownLiveCheckInterval time.Duration `config:"doc='',field=SHUTDOWN_LIVE_CHECK_INTERVAL_MS,format=ms"`
	ShutdownTimeout           time.Duration `config:"doc='',field=SHUTDOWN_TIMEOUT_MS,format=ms"`
	ShutdownAfterDelay        time.Duration `config:"doc='',field=SHUTDOWN_AFTER_DELAY_MS,format=ms"`
	AuthTestToken             string        `config:"field=AUTH_TEST_TOKEN,default=kjasomkfgds"`
}

func ProvideConfig(i do.Injector) (config Config, err error) {
	err = configgo.InitConfig(&config, configgo.Setting{
		Prefix: "HTTP",
		Title:  "Http config",
	}).Process()
	return
}
