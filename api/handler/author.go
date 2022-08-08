package handler

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/NajmiddinAbdulhakim/mybook/models"
	"github.com/gin-gonic/gin"
)

// CreateUser creates author
// @Summary Create author summary
// @Description This api is using create new author
// @Tags Author
// @Accept json
// @Produce json
// @Success 200 {string} Success
// @Param author body models.User true "author body"
// @Router /author/create [post]
func (h *Handler) CreateAuthor(c *gin.Context) {
	var body models.Author

	err := c.ShouldBindJSON(&body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		log.Printf(`failed to bind json for author: %v`, err)
	}
	fmt.Println(body)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*7)
	defer cancel()

	response, err := h.serviceManager.CreateAuthor(ctx, &body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			`error`: err.Error(),
		})
		log.Printf(`failed to Create Author: %v`, err)
		return
	}

	c.JSON(http.StatusCreated, response)
}
