package main

func initRoutes() {
	//Get Index
	router.GET("/", indexHanlder)
	router.GET("/flower", getFlowers)
}
