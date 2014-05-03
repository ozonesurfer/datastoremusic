package controllers

import (
	"datastoremusic/config"
	"datastoremusic/model"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
)

func HomeController(rw http.ResponseWriter, rq *http.Request) {
	bands, e := model.GetAllDocs(rq, config.BAND_TYPE)
	if e != nil {
		log.Println(e)
		os.Exit(1)
	}

	t, err := template.ParseFiles("src/datastoremusic/views/home/index.html")
	if err != nil {
		rw.WriteHeader(http.StatusNotFound)
		rw.Write([]byte("File not found"))
		fmt.Println("Template parse error:", err)
		return
	}
	t.Execute(rw, struct{ Bands []model.Doc }{Bands: bands})

}
