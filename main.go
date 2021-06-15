package main

import (
	picshot "github.com/Aakanksha-jais/picshot-golang-backend/pkg/app"

	handlerAccount "github.com/Aakanksha-jais/picshot-golang-backend/handlers/account"
	handlerBlog "github.com/Aakanksha-jais/picshot-golang-backend/handlers/blog"

	serviceAccount "github.com/Aakanksha-jais/picshot-golang-backend/services/account"
	serviceBlog "github.com/Aakanksha-jais/picshot-golang-backend/services/blog"
	serviceTag "github.com/Aakanksha-jais/picshot-golang-backend/services/tag"

	storeAccount "github.com/Aakanksha-jais/picshot-golang-backend/stores/account"
	storeBlog "github.com/Aakanksha-jais/picshot-golang-backend/stores/blog"
	storeImage "github.com/Aakanksha-jais/picshot-golang-backend/stores/image"
	storeTag "github.com/Aakanksha-jais/picshot-golang-backend/stores/tag"
)

func main() {
	app := picshot.New()

	// Dependency Injection
	blogStore := storeBlog.New()
	tagStore := storeTag.New()
	accountStore := storeAccount.New()
	imageStore := storeImage.New()

	tagService := serviceTag.New(tagStore)
	blogService := serviceBlog.New(blogStore, tagService, imageStore)
	accountService := serviceAccount.New(accountStore, blogService)

	blogHandler := handlerBlog.New(blogService)
	accountHandler := handlerAccount.New(accountService)

	// JWKS Endpoint
	app.POST("/.well-known/jwks.json", accountHandler.JWKSEndpoint)

	// Routes for Accounts
	app.POST("/login", accountHandler.Login)
	app.POST("/signup", accountHandler.Signup)
	app.POST("/logout", accountHandler.Logout)
	app.GET("/myaccount", accountHandler.Get)
	app.PUT("/myaccount", accountHandler.Update)
	app.GET("/user/{username}", accountHandler.GetUser)
	app.GET("/available", accountHandler.CheckAvailability)
	app.PUT("/password", accountHandler.UpdatePassword)
	app.DELETE("/myaccount", accountHandler.Delete)

	// Routes for Blogs
	app.GET("/blogs", blogHandler.GetAll)
	app.GET("/browse", blogHandler.Browse)
	app.POST("/blog", blogHandler.Create)
	app.GET("/blogs/{blogid}", blogHandler.Get)
	app.GET("/tags/{tag}", blogHandler.GetAllByTag)
	app.PUT("/blogs/{blogid}", blogHandler.Update)
	app.DELETE("/blogs/{blogid}", blogHandler.Delete)
	app.GET("/{accountid}/blogs", blogHandler.GetBlogsByUser)

	app.Start()
}
