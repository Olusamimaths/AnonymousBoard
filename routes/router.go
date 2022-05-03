package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/olusamimaths/AnonymousBoard/config"
	"github.com/olusamimaths/AnonymousBoard/controllers"
)

type Router interface {
	gin.IRouter
	Serve() error
	RegisterThreadRoutes(c controllers.ThreadController)
	RegisterReplyRoutes(c controllers.ReplyController)
}

type router struct {
	*gin.Engine
	c *config.Config
}

func NewRouter(c *config.Config) Router {
	config := c.Get()
	r := gin.New()

	if config.GetString("ENVIRONMENT") == "prod" {
		gin.SetMode(gin.ReleaseMode)
	}
	if config.GetBool("app.log") {
		r.Use(gin.Logger())
	}

	setupDefaults(r)
	return &router{Engine: r, c: c}
}

func (r *router) Serve() error {
	port := r.c.Get().GetString("app.port")
	return r.Run(":" + port)
}
