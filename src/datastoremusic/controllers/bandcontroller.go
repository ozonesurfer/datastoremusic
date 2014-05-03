package controllers

import (
	//	"appengine/datastore"
	"datastoremusic/config"
	"datastoremusic/model"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
)

func BandAddController(rw http.ResponseWriter, rq *http.Request) {
	t, err := template.ParseFiles("src/datastoremusic/views/band/add.html")
	if err != nil {
		rw.WriteHeader(http.StatusNotFound)
		rw.Write([]byte("File not found"))
		fmt.Println("Template parse error:", err)
		return
	}
	locations, e := model.GetAllDocs(rq, config.LOCATION_TYPE)
	if e != nil {
		log.Println(e)
		os.Exit(1)
	}
	t.Execute(rw, struct{ Locations []model.Doc }{Locations: locations})
}

func BandVerifyController(rw http.ResponseWriter, rq *http.Request) {
	/*	name := rq.FormValue("name")
		var locationId *datastore.Key
		var albums []model.Album
		locationType := rq.FormValue("loctype")
	*/
	http.Redirect(rw, rq, "/home/index", http.StatusFound)
}
