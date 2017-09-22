// Copyright (C) 2017 Next Thing Co. <software@nextthing.co>
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 2 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with this program.  If not, see <http://www.gnu.org/licenses/>

package main

import (
	"context"
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

var serverInitialized = false
var serverRunning = false
var router *gin.Engine
var server *http.Server
var credentialsChannel chan Credentials

var viewLocation = "./view/*"
var staticLocation = "./static"
var defaultPort = ":8080"

func initializeServer() error {
	if serverInitialized {
		return errors.New("Server already initialized")
	}
	credentialsChannel = make(chan Credentials)

	// Set to production mode if not debugging
	if !debugMode {
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
		Addr:    defaultPort,
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
	credentials = <-credentialsChannel
	return credentials, nil
}

func stopServer() {
	// Give the server 5 seconds to shutdown, then kill it
	Debug.Println("Received prompt to shut down server!")
	Info.Println("Shutting down server in 5 seconds")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
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
