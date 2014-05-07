// albumcontroller
package controllers

import (
	//	"fmt"
	"log"
	"net/http"
	"strings"
)

/*
func main() {
	fmt.Println("Hello World!")
}
*/

func AlbumIndexController(rw http.ResponseWriter, rq *http.Request) {
	url := rq.URL.Path
	parms := strings.Split(url, "/")
	rawId := parms[2]
	log.Println("rawId =", rawId)
	t, err := template.ParseFiles("src/datastoremusic/views/album/index.html")
	if err != nil {
		rw.WriteHeader(http.StatusNotFound)
		rw.Write([]byte("File not found"))
		fmt.Println("Template parse error:", err)
		return
	}
	t.Execute(rw, nil)
}
