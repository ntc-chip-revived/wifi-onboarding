package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func redirectToStatus(context *gin.Context) {
	Debug.Println("redirectToStatus: redirecting to /ap/status")
	context.Redirect(http.StatusFound, "/ap/status")
}

func renderStatus(context *gin.Context) {
	status := getApplicationStatus()
	render(context, gin.H{
		"title":    "Wifi Onboarding - Status",
		"pipeline": status},
		"status.html")
}
