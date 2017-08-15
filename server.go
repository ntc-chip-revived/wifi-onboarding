package main

import (
	"context"
	"errors"
	"net/http"
	"time"
	"github.com/gin-gonic/gin"
)

var serverInitialized = false
var serverRunning = false
var router *gin.Engine
var server *http.Server
var credentialsChannel chan Credentials

var viewLocation = "./view/*"
var staticLocation = "./static"

func initializeServer() error {
	if serverInitialized {
		return errors.New("Server already initialized")
	}
	credentialsChannel = make(chan Credentials)

	// Set to production mode if not debugging
	if (!debugMode) {
		gin.SetMode(gin.ReleaseMode)
	}

	// Create router with default middleware (log and recovery)
	router = gin.Default()

	// Serve static files
	router.Static("/static", staticLocation)
	router.StaticFile("/hotspot.html", "./static/hotspot.html")

	// Load view templates and init routes
	router.LoadHTMLGlob(viewLocation)

	// Initialize router routes
	initRoutes()

	server = &http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	serverInitialized = true
	return nil
}

func startServer() (Credentials, error) {
	if !serverInitialized {
		initializeServer()
	}

	if serverRunning {
		return Credentials{}, errors.New("Server already running")
	}
	serverRunning = true

	go func() {
		Debug.Println("startServer: listening on", server.Addr)
		if err := server.ListenAndServe(); err != nil {
			Debug.Println("startServer:", err)
		}
		Debug.Println("startServer thread completed")
	}()

	Debug.Println("startServer: listening on credentials channel")
	// Wait until someone pushes data to the credentialsChannel
	credentials = <- credentialsChannel
	return credentials, nil
}

func stopServer() {
	// Give the server 5 seconds to shutdown, then kill it
	Debug.Println("Received prompt to shut down server!")
	Info.Println("Shutting down server in 5 seconds")
	ctx, cancel := context.WithTimeout(context.Background(), 5 * time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		Error.Println("Server Shutdown:", err)
	}
	serverRunning = false
	Debug.Println("Server down")
}

func render(context *gin.Context, data gin.H, templateName string) {
	switch context.Request.Header.Get("Accept") {
	case "application/json":
		context.JSON(http.StatusOK, data["pipeline"])
	default:
		context.HTML(http.StatusOK, templateName, data)
	}
}
