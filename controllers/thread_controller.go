package controllers

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/olusamimaths/AnonymousBoard/models"
	"github.com/olusamimaths/AnonymousBoard/services"
	"github.com/olusamimaths/AnonymousBoard/utils"
)

type ThreadController interface {
	ListThreads(c *gin.Context)
	GetThread(c *gin.Context)
	CreateThread(c *gin.Context)
	Report(c *gin.Context)
	Delete(c *gin.Context)
}

type threadController struct {
	ts services.ThreadService
}

func NewThreadController(ts services.ThreadService) ThreadController {
	return &threadController{ts}
}

func (tc *threadController) ListThreads(c *gin.Context) {
	page := 0
	pageQuery := c.Query("page")

	if pageQuery != "" {
		p, err := strconv.Atoi(pageQuery)
		if err != nil {
			c.JSON(utils.CreateApiError(http.StatusBadRequest, errors.New("invalid page query parameter")))
			return
		}
		page = p
	}
	err, threads := tc.ts.List(page)
	if err != nil {
		c.JSON(utils.ErrorFromDatabase(err))
		return
	}

	c.JSON(http.StatusOK, threads)
	return
}

func (tc *threadController) GetThread(c *gin.Context) {
	id := c.Param("id")
	err, thread := tc.ts.GetByID(id)
	if err != nil {
		c.JSON(utils.ErrorFromDatabase(err))
		return
	}

	c.JSON(http.StatusOK, thread)
	return
}

func (tc *threadController) Report(c *gin.Context) {
	id := c.Param("id")
	err := tc.ts.Report(id)
	if err != nil {
		c.JSON(utils.ErrorFromDatabase(err))
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
	})
	return
}

func (tc *threadController) Delete(c *gin.Context) {
	id := c.Param("id")
	password := c.Query("password")
	if password == "" {
		c.JSON(utils.CreateApiError(http.StatusBadRequest, errors.New("you must supply password query")))
		return
	}
	err := tc.ts.DeleteWithPassword(id, password)
	if err != nil {
		c.JSON(utils.ErrorFromDatabase(err))
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
	})
	return
}

func (tc *threadController) CreateThread(c *gin.Context) {
	var t models.Thread
	if err := c.ShouldBindJSON(&t); err != nil {
		c.JSON(utils.CreateApiError(http.StatusBadRequest, errors.New("invalid request body")))
		return
	}
	err, thread := tc.ts.Create(t)
	if err != nil {
		c.JSON(utils.ErrorFromDatabase(err))
		return
	}

	c.JSON(http.StatusCreated, thread)
	return
}