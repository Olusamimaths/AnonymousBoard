package main

import (
	"fmt"

	"github.com/olusamimaths/AnonymousBoard/config"
	"github.com/olusamimaths/AnonymousBoard/database"
	"github.com/olusamimaths/AnonymousBoard/routes"
)

func main() {
	c := config.NewConfig()
	r := routes.NewRouter(c)

	conn := database.NewDatabaseConnection(c)

	fmt.Println(conn)

	r.Serve()
}
