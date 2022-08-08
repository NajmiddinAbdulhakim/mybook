package api

import (
	"github.com/NajmiddinAbdulhakim/mybook/api/handler"
	"github.com/NajmiddinAbdulhakim/mybook/service"

	"github.com/NajmiddinAbdulhakim/mybook/config"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"     // swagger embed files
	ginSwagger "github.com/swaggo/gin-swagger" // gin-swagger middleware
)

type Option struct {
	ServiceManager service.Service
	Conf           config.Config
}

func NewRouter(option Option) *gin.Engine {
	router := gin.Default()

	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	handler := handler.New(&handler.HandlerConfig{
		ServiceManager: option.ServiceManager,
		Cfg:            option.Conf,
	})

	author := router.Group("author")
	{
		author.POST(`/create`, handler.CreateAuthor)
	}

	book := router.Group("book")
	{
		book.POST(`/create`, handler.CreateBook)
	}

	user := router.Group("user")
	{
		user.POST(`/create`, handler.CreateUser)
	}

	url := ginSwagger.URL("seagger/doc.json")
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))
	router.Run()

	return router
}
