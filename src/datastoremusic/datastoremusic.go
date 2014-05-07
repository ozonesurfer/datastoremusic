package datastoremusic

import (
	"datastoremusic/controllers"
	"net/http"
)

func init() {
	http.HandleFunc("/", controllers.HomeController)
	http.HandleFunc("/home/index", controllers.HomeController)
	http.HandleFunc("/band/add", controllers.BandAddController)
	http.HandleFunc("/band/verify", controllers.BandVerifyController)
	http.HandleFunc("/album/index/", controllers.AlbumIndexController)
	http.HandleFunc("/album/add/", controllers.AlbumAddController)
	http.HandleFunc("/album/verify/", controllers.AlbumVerifyCountroller)
	http.Handle("/resources/", http.StripPrefix("/resources/", http.FileServer(http.Dir("resources"))))
}
