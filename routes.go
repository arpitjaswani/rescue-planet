package main

import (
	admin "admin/controllers"
	api "api/controllers"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func addRoutes() *httprouter.Router {
	router := httprouter.New()
	router.RedirectTrailingSlash = true

	// Static files server
	// ====================================================================
	router.ServeFiles("/assets/*filepath", http.Dir("vendor/assets"))

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
	router.GET("/admin/login", admin.Login)
	router.POST("/admin/credentials", admin.CheckLogin)
	router.GET("/admin/register", admin.Register)
	router.GET("/admin/adopt", IsAuthorized("admin", admin.Adopt))
	router.POST("/admin/addUser", GenerateToken("admin", admin.AddUser))

	return router
}
