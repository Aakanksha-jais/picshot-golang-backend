package main

import (
	"fmt"
	"net/http"

	"github.com/Aakanksha-jais/picshot-golang-backend/driver"
	handlerAccount "github.com/Aakanksha-jais/picshot-golang-backend/handlers/account"
	handlerBlog "github.com/Aakanksha-jais/picshot-golang-backend/handlers/blog"
	"github.com/Aakanksha-jais/picshot-golang-backend/middlewares"
	"github.com/Aakanksha-jais/picshot-golang-backend/pkg/configs"
	"github.com/Aakanksha-jais/picshot-golang-backend/pkg/log"
	serviceAccount "github.com/Aakanksha-jais/picshot-golang-backend/services/account"
	serviceBlog "github.com/Aakanksha-jais/picshot-golang-backend/services/blog"
	storeAccount "github.com/Aakanksha-jais/picshot-golang-backend/stores/account"
	storeBlog "github.com/Aakanksha-jais/picshot-golang-backend/stores/blog"
	storeTag "github.com/Aakanksha-jais/picshot-golang-backend/stores/tag"
	"github.com/gorilla/mux"
)

func main() {
	config, err := configs.NewConfigLoader("./configs")
	if err != nil {
		return
	}

	logger := log.NewLogger()
	mongoDB, err := driver.NewMongoConfigs(config).ConnectToMongo(logger)
	if err != nil {
		return
	}

	sqlDB, err := driver.NewSQLConfigs(config).ConnectToSQL(logger)
	if err != nil {
		return
	}

	// Dependency Injection
	blogStore := storeBlog.New(mongoDB, logger)
	tagStore := storeTag.New(mongoDB, logger)
	accountStore := storeAccount.New(sqlDB, logger)

	blogService := serviceBlog.New(blogStore, tagStore, logger)
	accountService := serviceAccount.New(accountStore, blogService, logger)

	blogHandler := handlerBlog.New(blogService, logger)
	accountHandler := handlerAccount.New(accountService, logger, config)

	r := mux.NewRouter()

	// Routes for Accounts
	r.HandleFunc("/login", accountHandler.Login).Methods(http.MethodPost)
	r.HandleFunc("/signup", accountHandler.Signup).Methods(http.MethodPost)
	r.HandleFunc("/myaccount", accountHandler.Get).Methods(http.MethodGet)
	r.HandleFunc("/myaccount", accountHandler.Update).Methods(http.MethodPut)
	r.HandleFunc("/user/{username}", accountHandler.GetUser).Methods(http.MethodGet)
	r.HandleFunc("/available", accountHandler.CheckAvailability).Queries("username", "{username}").Methods(http.MethodGet)
	r.HandleFunc("/available", accountHandler.CheckAvailability).Queries("email", "{email}").Methods(http.MethodGet)
	r.HandleFunc("/available", accountHandler.CheckAvailability).Queries("phone", "{phone}").Methods(http.MethodGet)

	// Routes for Blogs
	r.HandleFunc("/blogs", blogHandler.GetAll).Methods(http.MethodGet)

	// Middlewares
	r.Use(middlewares.Authentication(config, logger))

	server := &http.Server{
		Handler: r,
		Addr:    fmt.Sprintf("localhost:%s", config.Get("HTTP_PORT")),
	}

	logger.Infof("starting server at PORT: %v", config.Get("HTTP_PORT"))

	logger.Fatalf("error in starting the server: %v", server.ListenAndServe())
}
