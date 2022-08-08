package handler

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/NajmiddinAbdulhakim/mybook/models"
	"github.com/gin-gonic/gin"
)

// CreateUser creates user
// @Summary Create user summary
// @Description This api is using create new user
// @Tags User
// @Accept json
// @Produce json
// @Success 200 {string} Success
// @Param user body models.User true "user body"
// @Router /user/create [post]
func (h *Handler) CreateUser(c *gin.Context) {
	var body models.User

	err := c.ShouldBindJSON(&body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		log.Fatalf(`failed to bind json for user: %v`, err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*7)
	defer cancel()

	response, err := h.serviceManager.CreateUser(ctx,&body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			`error`: err.Error(),
		})
		log.Fatalf(`failed to create user: %v`, err)
	}

	c.JSON(http.StatusCreated,response)
}
