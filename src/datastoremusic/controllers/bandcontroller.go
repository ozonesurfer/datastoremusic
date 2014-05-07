package controllers

import (
	"appengine"
	"appengine/datastore"
	"datastoremusic/config"
	"datastoremusic/model"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
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
	c := appengine.NewContext(rq)
	name := rq.FormValue("name")
	message := "no errors"
	var locationId *datastore.Key
	var albums []model.Album
	locationType := rq.FormValue("loctype")
	switch locationType {
	case "existing":
		rawId := rq.FormValue("location_id")
		q := strings.Split(rawId, ",")
		log.Println("rawId =", rawId)
		x := q[1]
		id_int, e := strconv.ParseInt(x, 10, 64)
		if e != nil {
			message = e.Error()
		}
		locationId = datastore.NewKey(c, config.LOCATION_TYPE, "", id_int, nil)
		log.Printf("locationId =", locationId)
		//	message = "not implemented yet"
		break
	case "new":
		var err error
		location := model.Location{rq.FormValue("city"), rq.FormValue("state"), rq.FormValue("country")}
		locationId, err = model.AddLocation(location, rq)
		log.Printf("locationId =", locationId)
		if err != nil {
			message = err.Error()
		}
		break
	}
	if message == "no errors" {
		band := model.Band{Name: name, LocationId: locationId, Albums: albums}
		_, err := model.AddBand(band, rq)
		if err != nil {
			message = "Band add: " + err.Error()
		}
	}
	t, err := template.ParseFiles("src/datastoremusic/views/band/verify.html")
	if err != nil {
		rw.WriteHeader(http.StatusNotFound)
		rw.Write([]byte("File not found"))
		fmt.Println("Template parse error:", err)
		return
	}
	//http.Redirect(rw, rq, "/home/index", http.StatusFound)
	t.Execute(rw, message)
}
