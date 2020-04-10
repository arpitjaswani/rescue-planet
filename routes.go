package main

import (
	"api/controllers"

	"github.com/julienschmidt/httprouter"
)

func addRoutes() *httprouter.Router {
	router := httprouter.New()

	// base routes
	router.GET("/", controllers.Index)

	// user routes
	router.GET("/user", controllers.GetUsers)
	router.GET("/user/:userid", controllers.GetUsers)
	router.POST("/user/add", controllers.AddUser)
	router.PATCH("/user/update", controllers.UpdateUser)
	router.PATCH("/user/deactivate/:userid", controllers.DeactivateUser)

	return router
}
