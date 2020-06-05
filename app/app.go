package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
)

//Get all flowers in the line
func getFlowers(c *gin.Context) {
	//	server, _ := c.GetPostForm("server")
	server, _ := c.GetQuery("server")
	flower, _ := c.GetQuery("flower")

	response, err := http.Get("https://api.jx3box.com/api/flower/price/query?server=" + server + "&flower=" + flower)
	if err != nil {
		c.HTML(
			http.StatusNoContent,
			"home/index.html",
			gin.H{
				"title": "Not Found"},
		)
	}
	data, _ := ioutil.ReadAll(response.Body)

	var content Content
	json.Unmarshal(data, &content)

	c.HTML(
		http.StatusOK,
		"flowers/flowers_list.html",
		gin.H{
			"title":   server + " : " + flower,
			"serverf": server + " : " + flower,
		})

	for _, flower := range content.Data {
		c.HTML(http.StatusOK, "flowers/flowers.html",
			gin.H{
				"Map":   flower.Map,
				"Unit":  flower.Unit,
				"Price": flower.Price,
			})
	}

	//	fmt.Fprint(w, string(data))
}

func indexHanlder(c *gin.Context) {
	c.HTML(
		http.StatusOK,
		"home/index.html",
		gin.H{
			"title": "花花价",
		},
	)
}

//Router
var router *gin.Engine

func main() {
	//Set the default router
	router = gin.Default()

	//Views
	router.LoadHTMLGlob("../templates/**/*")

	//Route
	initRoutes()

	//static file serve
	router.Static("/static", "../static")

	router.Run()
}
