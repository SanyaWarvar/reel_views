package container

import (
	"rv/config"
	"rv/internal/endpoint/controller/http"
	v1 "rv/internal/endpoint/controller/http/api/v1"
	"rv/pkg/applogger"
	"rv/pkg/database/dragonfly"
	"rv/pkg/database/postgres"
	"rv/pkg/httpserver"
	"rv/pkg/response"
	"rv/pkg/restclient"
	"rv/pkg/trx"
)

type Container struct {
	cfg                *config.Config
	cacheClient        *dragonfly.Client
	pool               *postgres.Pool
	poolContextManager *postgres.ContextManager
	logger             applogger.Logger
	builder            *response.Builder
	trxManager         trx.TransactionManager
	httpServer         *httpserver.Server
	httpDispatcher     *v1.Dispatcher
	httpKernel         *http.Kernel
	restClient         restclient.RestClient

	repositories *repositories
	applications *applications
	services     *services
	workers      *workers
	caches       *cache
}

func New(cfg *config.Config) *Container {
	return &Container{cfg: cfg}
}

func (c *Container) getConfig() *config.Config {
	return c.cfg
}

func (c *Container) Start() error {
	//migrate
	if err := c.Migrate(); err != nil {
		return err
	}

	//servers
	c.getHTTPServer().Start()

	//workers
	if err := c.getWorkers().start(); err != nil {
		return err
	}
	return nil
}

func (c *Container) Stop() error {
	if err := c.getHTTPServer().Shutdown(); err != nil {
		return err
	}
	return nil
}
