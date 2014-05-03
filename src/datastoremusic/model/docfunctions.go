package model

import (
	"appengine"
	"appengine/datastore"
	"datastoremusic/config"
	"log"
	"net/http"
)

type Doc struct {
	Id    *datastore.Key
	Value interface{}
}

type Band struct {
	Name       string
	LocationId *datastore.Key
	Albums     []Album
}

type Album struct {
	Name    string
	GenreId *datastore.Key
	Year    int
}

type Location struct {
	City, State, Country string
}

func GetAllDocs(rq *http.Request, docType string) ([]Doc, error) {
	c := appengine.NewContext(rq)
	var docs []Doc
	q := datastore.NewQuery(docType).KeysOnly()
	keys, err := q.GetAll(c, nil)
	if err != nil {
		log.Println("GetAllDocs error: " + err.Error())
		return nil, err
	}
	var obj interface{}
	for _, key := range keys {
		switch docType {
		case config.BAND_TYPE:
			obj = Band{}
			break
		case config.LOCATION_TYPE:
			obj = Location{}
		}
		err2 := datastore.Get(c, key, &obj)
		if err2 != nil {
			log.Println("GetAllDocs error: " + err2.Error())
			return nil, err2
		}
		doc := Doc{Id: key, Value: obj}
		docs = append(docs, doc)
	}
	return docs, nil
}

func (this Doc) LocToString() string {
	var city, state, country string
	location := this.Value.(Location)
	if location.City != "" {
		city = location.City
	} else {
		city = "(city)"
	}
	if location.State != "" {
		state = location.State
	} else {
		state = "(state/province)"
	}
	country = location.Country
	locString := city + ", " + state + " " + country
	return locString
}
