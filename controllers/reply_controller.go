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

type ReplyController interface {
	ListReplies(c *gin.Context)
	GetReply(c *gin.Context)
	CreateReply(c *gin.Context)
	ReportReply(c *gin.Context)
	DeleteReply(c *gin.Context)
}

type replyController struct {
	rs services.ReplyService
}

func NewReplyController(rs services.ReplyService) ReplyController {
	return &replyController{rs}
}

func (rc *replyController) ListReplies(c *gin.Context) {
	tid := c.Param("tid")
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
	err, replys := rc.rs.ListByThreadID(tid, page)
	if err != nil {
		c.JSON(utils.ErrorFromDatabase(err))
		return
	}

	c.JSON(http.StatusOK, replys)
	return
}

func (rc *replyController) GetReply(c *gin.Context) {
	id := c.Param("id")
	err, reply := rc.rs.GetByID(id)
	if err != nil {
		c.JSON(utils.ErrorFromDatabase(err))
		return
	}

	c.JSON(http.StatusOK, reply)
	return
}

func (rc *replyController) ReportReply(c *gin.Context) {
	id := c.Param("id")
	err := rc.rs.Report(id)
	if err != nil {
		c.JSON(utils.ErrorFromDatabase(err))
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
	})
	return
}

func (rc *replyController) DeleteReply(c *gin.Context) {
	id := c.Param("id")
	password := c.Query("password")
	if password == "" {
		c.JSON(utils.CreateApiError(http.StatusBadRequest, errors.New("you must supply password query")))
		return
	}
	err := rc.rs.DeleteWithPassword(id, password)
	if err != nil {
		c.JSON(utils.ErrorFromDatabase(err))
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
	})
	return
}

func (rc *replyController) CreateReply(c *gin.Context) {
	tid := c.Param("tid")
	var r models.Reply

	if err := c.ShouldBindJSON(&r); err != nil {
		c.JSON(utils.CreateApiError(http.StatusBadRequest, errors.New("invalid request body")))
		return
	}
	err, reply := rc.rs.Create(tid, r)
	if err != nil {
		c.JSON(utils.ErrorFromDatabase(err))
		return
	}

	c.JSON(http.StatusCreated, reply)
	return
}
