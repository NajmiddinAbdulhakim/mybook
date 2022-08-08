package handler

import (
	"github.com/NajmiddinAbdulhakim/mybook/config"
	"github.com/NajmiddinAbdulhakim/mybook/service"
)

type Handler struct {
	serviceManager service.Repository
	cfg            config.Config
}

type HandlerConfig struct {
	ServiceManager service.Repository
	Cfg            config.Config
}

func New(c *HandlerConfig) *Handler {
	return &Handler{
		cfg:            c.Cfg,
		serviceManager: c.ServiceManager,
	}
}
