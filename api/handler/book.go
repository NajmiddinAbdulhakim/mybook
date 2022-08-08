package handler

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/NajmiddinAbdulhakim/mybook/models"
	"github.com/gin-gonic/gin"
)
// CreateUser creates book
// @Summary Create book summary
// @Description This api is using create new book
// @Tags Book
// @Accept json
// @Produce json
// @Success 200 {string} Success
// @Param book body models.User true "book body"
// @Router /book/create [post]
func (h *Handler) CreateBook(c *gin.Context) {

	var body models.Book

	err := c.ShouldBindJSON(&body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			`error`: err.Error(),
		})
		log.Printf("failed to bind json for book")
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*7)
	defer cancel()
	
	response, err := h.serviceManager.CreateBook(ctx, &body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			`error`: err.Error(),
		})
		log.Printf("failed to create book: %v", err)

	}

	c.JSON(http.StatusCreated, response)
}
