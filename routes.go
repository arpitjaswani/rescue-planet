package main

import (
	admin "admin/controllers"
	api "api/controllers"

	"github.com/julienschmidt/httprouter"
)

func addRoutes() *httprouter.Router {
	router := httprouter.New()

	// API Routes
	//==================================================

	// base
	router.GET("/", api.Index)
	// users
	router.GET("/user", api.GetUsers)
	router.GET("/user/:userid", api.GetUsers)
	router.POST("/user/add", api.AddUser)
	router.PATCH("/user/update", api.UpdateUser)
	router.PATCH("/user/deactivate/:userid", api.DeactivateUser)

	// Admin Routes
	//==================================================

	// base
	router.GET("/admin", admin.Index)
	// users
	// router.GET("/admin/login", controllers.Login)

	return router
}
