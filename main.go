package main

import (
	"fmt"
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
	"net/http"
)

func main() {

	logger := log.NewLogger()
	config := configs.NewConfigLoader("./configs", logger)

	mongoDB, err := driver.NewMongoConfigs(config).ConnectToMongo()
	if err != nil {
		logger.Fatalf("cannot connect to  mongo %v", err)
		return
	}

	sqlDB, err := driver.NewSQLConfigs(config).ConnectToSQL()
	if err != nil {
		logger.Fatalf("cannot connect to sql %v", err)
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

	// Routes
	r.HandleFunc("/login", accountHandler.LogIn)
	r.HandleFunc("/signup", accountHandler.SignUp)
	r.HandleFunc("/blogs", blogHandler.GetAll)

	// Middlewares
	r.Use(middlewares.Authentication(config, logger))

	server := &http.Server{
		Handler: r,
		Addr:    fmt.Sprintf("localhost:%s", config.Get("HTTP_PORT")),
	}

	logger.Fatalf("error in starting the server: ", server.ListenAndServe())
}
