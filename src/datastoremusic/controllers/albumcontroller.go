// albumcontroller
package controllers

import (
	//	"fmt"
	"appengine"
	"appengine/datastore"
	"datastoremusic/config"
	"datastoremusic/model"
	"html/template"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
)

/*
func main() {
	fmt.Println("Hello World!")
}
*/

func AlbumIndexController(rw http.ResponseWriter, rq *http.Request) {
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
	bandId := datastore.NewKey(c, config.BAND_TYPE, "", id_int, nil)
	band, err := model.GetBand(bandId, rq)
	if err != nil {
		log.Println("Band get error:", err)
		os.Exit(1)
	}
	title := "Albums by " + band.Name
	t, err := template.ParseFiles("src/datastoremusic/views/album/index.html")
	if err != nil {
		rw.WriteHeader(http.StatusNotFound)
		rw.Write([]byte("File not found"))
		log.Println("Template parse error:", err)
		return
	}
	t.Execute(rw, struct {
		Id      *datastore.Key
		Title   string
		Band    model.Band
		Request *http.Request
	}{Id: bandId, Title: title, Band: band, Request: rq})
}

func AlbumAddController(rw http.ResponseWriter, rq *http.Request) {
	url := rq.URL.Path
	parms := strings.Split(url, "/")
	id := parms[3]
	t, err := template.ParseFiles("src/datastoremusic/views/album/add.html")
	if err != nil {
		rw.WriteHeader(http.StatusNotFound)
		rw.Write([]byte("File not found"))
		log.Println("Template parse error:", err)
		return
	}
	genres, e := model.GetAllDocs(rq, config.GENRE_TYPE)
	if e != nil {
		log.Println(e)
		os.Exit(1)
	}
	t.Execute(rw, struct {
		Genres []model.Doc
		Id     string
		Title  string
	}{Genres: genres, Id: id, Title: "Adding Album"})
}

func AlbumVerifyCountroller(rw http.ResponseWriter, rq *http.Request) {
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
	bandId := datastore.NewKey(c, config.BAND_TYPE, "", id_int, nil)
	message := "no errors"
	var genreId *datastore.Key

	genreType := rq.FormValue("genretype")
	switch genreType {
	case "existing":
		rawId := rq.FormValue("genre_id")
		q := strings.Split(rawId, ",")
		log.Println("rawId =", rawId)
		x := q[1]
		id_int, e := strconv.ParseInt(x, 10, 64)
		if e != nil {
			message = e.Error()
		}
		genreId = datastore.NewKey(c, config.GENRE_TYPE, "", id_int, nil)
		log.Printf("genreId =", genreId)
		//	message = "not implemented yet"
		break
	case "new":
		var err error
		genre := model.Genre{rq.FormValue("genre_name")}
		genreId, err = model.AddGenre(genre, rq)
		log.Printf("genreId =", genreId)
		if err != nil {
			message = err.Error()
		}
		break
	}
	if message == "no errors" {

		year, _ := strconv.Atoi(rq.FormValue("year"))
		album := model.Album{Name: rq.FormValue("name"),
			Year: year, GenreId: genreId}
		err := model.AddAlbum(album, bandId, rq)
		if err != nil {
			message = err.Error()
		}

	}
	t, err := template.ParseFiles("src/datastoremusic/views/album/verify.html")
	if err != nil {
		rw.WriteHeader(http.StatusNotFound)
		rw.Write([]byte("File not found"))
		log.Println("Template parse error:", err)
		return
	}
	t.Execute(rw, struct {
		Id      *datastore.Key
		Title   string
		Message string
	}{Id: bandId, Title: "Verifying Album", Message: message})

}
