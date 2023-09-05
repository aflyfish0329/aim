package main

import (
	"test/context/im/adapter/in/web/router"
	"test/context/im/infra/db"
	"test/context/im/infra/web"
)

func main() {
	startWeb()
}

func startWeb() {
	db := db.NewDB()
	c := web.NewWeb()

	router.Setup(c, db)

	c.Run(":8080")
}
