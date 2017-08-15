package main

func initRoutes() {
	router.GET("/", redirectToStatus)
	router.GET("/generate_204", redirectToStatus)
	
	apRoutes := router.Group("/ap")
	{
		apRoutes.GET("/list", renderAccessPointList)
		apRoutes.GET("/connect", renderAccessPointAuthentication)
		apRoutes.POST("/connect", renderConnectionStatus)
		apRoutes.GET("/status", renderStatus)
	}
}
