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
	"github.com/gin-gonic/gin"
	"github.com/nextthingco/gonnman"
	"net/http"
)

// Endpoint handler for GET /list
func renderAccessPointList(context *gin.Context) {
	render(context, gin.H{
		"title":    "Wifi Onboarding - List of Access Points",
		"pipeline": connectionServices},
		"list.html")
}

// Endpoint handler for GET /connect?id=<id>
// Wifi path is required!
func renderAccessPointAuthentication(context *gin.Context) {
	wifiPath := context.Query("id")
	Debug.Println("renderAccessPointAuthentication: Wifi path", wifiPath)

	if len(wifiPath) > 0 {
		if accessPoint, err := getAccessPointByPath(wifiPath); err == nil {
			render(context, gin.H{
				"title":    "Wifi Onboarding - Connect to " + accessPoint.Name,
				"pipeline": accessPoint},
				"connect.html")
		} else {
			Error.Println("renderAccessPointAuthentication: Access point not found")
			context.AbortWithError(http.StatusNotFound, err)
		}
	} else {
		Warning.Println("renderAccessPointAuthentication: Missing id query argument in url")
		context.Redirect(http.StatusFound, "/ap/status")
	}
}

// Helper function use to connect to wifi
func connectToAccessPoint(accessPointPath string, accessPointPassKey string) {
	var err error
	var serviceToConnect *connman.Service
	serviceToConnect = nil

	Debug.Printf("connectToAccessPoint: Connecting with key: %s, path: %s", accessPointPassKey, accessPointPath)
	// Search for access point and save credentials
	if serviceToConnect, err = getAccessPointByPath(accessPointPath); err != nil {
		Error.Println("connectToAccessPoint", err)
	} else {
		setCredentialSSID(serviceToConnect.Name)
	}

	if serviceToConnect != nil {
		if err := serviceToConnect.Connect(accessPointPassKey); err != nil {
			Warning.Println("connectToAccessPoint: Failed to Connect", err)
			setStatusError(err)
			setStatusConnecting(false)
			setStatusConnected(false)
			setStatusHasCredentials(false)
		} else {
			Debug.Println("connectToAccessPoint: Connection Successful")
			setStatusError(nil)
			setStatusConnecting(false)
			setStatusConnected(true)
		}
	}
}

// Endpoint handler for POST /connect
func renderConnectionStatus(context *gin.Context) {
	accessPointPassKey := context.PostForm("accessPointPassKey")
	accessPointPath := context.PostForm("accessPointPath")

	setCredentialPath(accessPointPath)
	setCredentialPSK(accessPointPassKey)
	setStatusHasCredentials(true)
	setStatusConnecting(true)

	status := getApplicationStatus()
	render(context, gin.H{
		"title":    "Wifi Onboarding - Attempting to Connect",
		"pipeline": status},
		"status.html")

	Debug.Println("Sending stop signal to credential channel")
	creds := getCredentials()
	// Push to the credentialsChannel
	credentialsChannel <- creds
	return
}
