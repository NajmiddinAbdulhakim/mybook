package handler

import (
	"github.com/NajmiddinAbdulhakim/mybook/config"
	"github.com/NajmiddinAbdulhakim/mybook/service"
)

type Handler struct {
	serviceManager service.Service
	cfg            config.Config
}

type HandlerConfig struct {
	ServiceManager service.Service
	Cfg            config.Config
}

func New(c *HandlerConfig) *Handler {
	return &Handler{
		cfg:            c.Cfg,
		serviceManager: c.ServiceManager,
	}
}
