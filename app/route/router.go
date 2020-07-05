package route

import (
	"flashlight/app/controller"
	"flashlight/app/middleware"
	"github.com/gorilla/mux"
	"net/http"
)

//GetRouter is a router that routes requests to their controller
func GetRouter() *mux.Router {
	r := mux.NewRouter()

	//controllers
	r.HandleFunc("/", controller.Index).Methods("GET")
	r.HandleFunc("/gallery", controller.Gallery).Methods("GET")
	r.HandleFunc("/uploads", controller.Upload).Methods("GET")

	r.HandleFunc("/login", controller.Login)
	r.HandleFunc("/logout", middleware.DeAuth(controller.Logout))
	r.HandleFunc("/register", controller.RegisterPage)

	//"REST" "Controller for Sending, Liking, Deleting Posts,Uploading
	r.HandleFunc("/sendComment", middleware.Auth(controller.SendComment)).Methods("POST")
	r.HandleFunc("/like", nil).Methods("POST")
	r.HandleFunc("/deletePicture", middleware.Auth(controller.DeletePicture)).Methods("POST")
	r.HandleFunc("/uploading", middleware.Auth(controller.Uploading)).Methods("POST")

	r.HandleFunc("/registerProcess", controller.RegisterProcess).Methods("POST")
	r.HandleFunc("/authenticate", controller.AuthenticateUser).Methods("POST")

	/*File Server & CND */

	//statics
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static", http.FileServer(http.Dir("./static/"))))
	//See all files on server
	r.PathPrefix("/files/").Handler(http.StripPrefix("/files", http.FileServer(http.Dir("."))))
	//Content-Files
	r.PathPrefix("/content/").Handler(http.StripPrefix("/content", http.FileServer(http.Dir("./content/"))))

	return r
}
