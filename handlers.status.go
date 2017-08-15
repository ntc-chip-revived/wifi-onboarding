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
