package routes

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/olusamimaths/AnonymousBoard/controllers"
	"github.com/olusamimaths/AnonymousBoard/utils"
)

func setupDefaults(r *gin.Engine) {
	r.Use(gin.Recovery())
	r.GET("/api", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"health": "ok"})
		return
	})
	r.NoRoute(func(c *gin.Context) {
		c.JSON(utils.CreateApiError(http.StatusNotFound, errors.New("no route found")))
		return
	})
}

func (r *router) RegisterThreadRoutes(c controllers.ThreadController) {
	rg := r.Group("/api/threads")
	rg.GET("/", c.ListThreads)
	rg.POST("/", c.CreateThread)
	rg.GET("/:id", c.GetThread)
	rg.PUT("/:id", c.Report)
	rg.DELETE("/:id", c.Delete)
}

func (r *router) RegisterReplyRoutes(c controllers.ReplyController) {
	rg := r.Group("/api/replies/:tid")
	rg.GET("/", c.ListReplies)
	rg.POST("/", c.CreateReply)
	rg.GET("/:id", c.GetReply)
	rg.PUT("/:id", c.ReportReply)
	rg.DELETE("/:id", c.DeleteReply)
}