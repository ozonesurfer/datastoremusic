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
	http.Handle("/resources/", http.StripPrefix("/resources/", http.FileServer(http.Dir("resources"))))
}
