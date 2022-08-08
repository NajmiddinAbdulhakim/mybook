package api

import (
	"github.com/NajmiddinAbdulhakim/mybook/service"
	"github.com/gin-gonic/gin"
)


func NewRouter(repo service.Service) {
	r := gin.Default()
	s := service.Service{repo: repo}
}