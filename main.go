package main

import (
	"github.com/Aakanksha-jais/picshot-golang-backend/pkg/app"

	handlerAccount "github.com/Aakanksha-jais/picshot-golang-backend/handlers/account"
	handlerBlog "github.com/Aakanksha-jais/picshot-golang-backend/handlers/blog"
	serviceAccount "github.com/Aakanksha-jais/picshot-golang-backend/services/account"
	serviceBlog "github.com/Aakanksha-jais/picshot-golang-backend/services/blog"
	storeAccount "github.com/Aakanksha-jais/picshot-golang-backend/stores/account"
	storeBlog "github.com/Aakanksha-jais/picshot-golang-backend/stores/blog"
	storeImage "github.com/Aakanksha-jais/picshot-golang-backend/stores/image"
	storeTag "github.com/Aakanksha-jais/picshot-golang-backend/stores/tag"
)

func main() {
	app := app.New()

	// Dependency Injection
	blogStore := storeBlog.New()
	tagStore := storeTag.New()
	accountStore := storeAccount.New()
	imageStore := storeImage.New()

	blogService := serviceBlog.New(blogStore, tagStore, imageStore)
	accountService := serviceAccount.New(accountStore, blogService)

	blogHandler := handlerBlog.New(blogService)
	accountHandler := handlerAccount.New(accountService)

	// Routes for Accounts
	app.POST("/login", accountHandler.Login)
	app.POST("/signup", accountHandler.Signup)
	app.GET("/myaccount", accountHandler.Get)
	app.PUT("/myaccount", accountHandler.Update)
	app.GET("/user/{username}", accountHandler.GetUser)
	app.GET("/available", accountHandler.CheckAvailability)
	app.POST("/changepassword", accountHandler.UpdatePassword)

	// Routes for Blogs
	app.GET("/blogs", blogHandler.GetAll)
	app.POST("/blog", blogHandler.Create)
	app.GET("/blogs/{blogid}", blogHandler.Get)
	app.GET("/{accountid}/blogs", blogHandler.GetBlogsByUser)

	app.Start()
}
