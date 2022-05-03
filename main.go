package main

import (
	"fmt"

	"github.com/olusamimaths/AnonymousBoard/config"
	"github.com/olusamimaths/AnonymousBoard/controllers"
	"github.com/olusamimaths/AnonymousBoard/database"
	"github.com/olusamimaths/AnonymousBoard/routes"
	"github.com/olusamimaths/AnonymousBoard/services"
)

func main() {
	c := config.NewConfig()
	r := routes.NewRouter(c)

	conn := database.NewDatabaseConnection(c)

	ts := services.NewThreadService(conn)
	rs := services.NewReplyService(conn)

	tc := controllers.NewThreadController(ts)
	rc := controllers.NewReplyController(rs)

	r.RegisterThreadRoutes(tc)
	r.RegisterReplyRoutes(rc)

	fmt.Println(conn)

	r.Serve()
}
