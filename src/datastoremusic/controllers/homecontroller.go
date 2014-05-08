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

func GenreListController(rw http.ResponseWriter, rq *http.Request) {
	genres, e := model.GetAllDocs(rq, config.GENRE_TYPE)
	if e != nil {
		log.Println(e)
		os.Exit(1)
	}

	t, err := template.ParseFiles("src/datastoremusic/views/home/genrelist.html")
	if err != nil {
		rw.WriteHeader(http.StatusNotFound)
		rw.Write([]byte("File not found"))
		fmt.Println("Template parse error:", err)
		return
	}
	t.Execute(rw, struct{ Model []model.Doc }{Model: genres})
}

func ByGenreController(rw http.ResponseWriter, rq *http.Request) {
	c := appengine.NewContext(rq)
	url := rq.URL.Path
	parms := strings.Split(url, "/")
	rawId := parms[3]
	log.Println("rawId =", rawId)
	q := strings.Split(rawId, ",")
	x := q[1]
	id_int, e := strconv.ParseInt(x, 10, 64)
	if e != nil {
		log.Println("Parse error:", e)
		os.Exit(1)

	}
	genreId := datastore.NewKey(c, config.GENRE_TYPE, "", id_int, nil)
	log.Println("Sending key", genreId)
	genreName := model.GetGenreName(rq, genreId)
	title := genreName + " Albums"
	bands, e := model.GetBandsByGenre(rq, genreId)
	if e != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		rw.Write([]byte("Bands not found"))
		log.Println("Band retrieval:", e)
		return
	}
	if len(bands) == 0 {
		log.Println("No match was found")
	}
	t, err := template.ParseFiles("src/datastoremusic/views/home/bygenre.html")
	if err != nil {
		rw.WriteHeader(http.StatusNotFound)
		rw.Write([]byte("File not found"))
		log.Println("Template parse error:", err)
		return
	}
	t.Execute(rw, struct {
		Model   []*model.Doc
		Title   string
		Request *http.Request
	}{Model: bands, Title: title, Request: rq})

}
