package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
)

//Get all flowers in the line
func getFlowers(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	server := r.PostForm.Get("server")
	flower := r.PostForm.Get("flower")

	response, err := http.Get("https://api.jx3box.com/api/flower/price/query?server=" + server + "&flower=" + flower)
	if err != nil {
		fmt.Fprint(w, "Getting wrong data!")
	}
	data, _ := ioutil.ReadAll(response.Body)

	var content Content
	parserErr := json.Unmarshal(data, &content)
	if parserErr != nil {
		fmt.Fprint(w, parserErr.Error())
	}

	t, _ := template.ParseFiles("views/flowers.html")

	for _, flower := range content.Data {
		//fmt.Fprint(w, "地图: "+flower.Map)
		t.Execute(w, flower.Map)
	}

	//	fmt.Fprint(w, string(data))
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("views/index.html")
	t.Execute(w, nil)
}

func main() {
	//Router
	router := mux.NewRouter()
	router.HandleFunc("/", indexHandler).Methods("GET")
	//	router.HandleFunc("/flowers", getFlowers)
	router.HandleFunc("/", getFlowers).Methods("POST")
	http.ListenAndServe(":3000", router)
}
