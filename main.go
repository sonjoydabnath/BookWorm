package main

import (
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/sonjoydabnath/BookWorm/controller"
	"github.com/sonjoydabnath/BookWorm/model/configs"
	"github.com/sonjoydabnath/BookWorm/model/dbcon"
	"github.com/sonjoydabnath/BookWorm/view"
)

//html page handler
/*func HtmlHandler() {

	http.Handle("/uploads/", http.StripPrefix("/uploads/", http.FileServer(http.Dir("uploads"))))    //file server for raw file serving inside html
	http.Handle("/resource/", http.StripPrefix("/resource/", http.FileServer(http.Dir("resource")))) //file server for raw file serving inside html


	http.HandleFunc("/", controller.Home)
	http.HandleFunc("/login", controller.Login)
	http.HandleFunc("/logout", controller.Logout)
	http.HandleFunc("/signup", controller.SignUp)
	http.HandleFunc("/about", controller.About)
	http.HandleFunc("/contact", controller.Contact)
	http.HandleFunc("/user-home", controller.UserHome)
	http.HandleFunc("/my-unpublished-book", controller.MyUnpublishedBook)
	http.HandleFunc("/publish-new-book", controller.PublishNewBook)
	http.HandleFunc("/my-published-book", controller.MyPublishedBook)
	http.HandleFunc("/user-list", controller.UserList)
	http.HandleFunc("/un-published-book", controller.UnPublishedBook)
	http.HandleFunc("/publishedbook", controller.PublishedBook)
	http.HandleFunc("/admin-review-book", controller.AdminReviewBook)
	http.HandleFunc("/approve-book", controller.ApproveBook)
	http.HandleFunc("/reject", controller.RejectBook)
	http.HandleFunc("/update-book", controller.UpdateBook)
	http.HandleFunc("/view-book", controller.ViewBook)
	http.HandleFunc("/subscribe-book", controller.SubscribeBook)
	http.HandleFunc("/unsubscribe-book", controller.UnsubscribeBook)
	http.HandleFunc("/unpublish-book", controller.UnpublishBook)

	//undone yet in mvc fashion
	/*http.HandleFunc("/view-user", ViewUser)//ar lagbe na view user
	/*http.HandleFunc("/block-user", BlockUser)
	http.HandleFunc("/send-notification", SendNotification)
	http.HandleFunc("/submit-notification", SubmitNotification)

}*/

func HtmlHandlerMux() {

	http.Handle("/uploads/", http.StripPrefix("/uploads/", http.FileServer(http.Dir("uploads"))))    //file server for raw file serving inside html
	http.Handle("/resource/", http.StripPrefix("/resource/", http.FileServer(http.Dir("resource")))) //file server for raw file serving inside html

	router.HandleFunc("/", controller.Home)
	router.HandleFunc("/login", controller.Login)
	router.HandleFunc("/logout", controller.LogOut)
	router.HandleFunc("/signup", controller.SignUp)
	router.HandleFunc("/about", controller.About)
	router.HandleFunc("/contact", controller.Contact)
	router.HandleFunc("/user-home", controller.UserHome)
	router.HandleFunc("/my-un-published-book", controller.MyUnPublishedBook)
	router.HandleFunc("/publish-new-book", controller.PublishNewBook)
	router.HandleFunc("/my-published-book", controller.MyPublishedBook)
	router.HandleFunc("/user-list", controller.UserList)
	router.HandleFunc("/un-published-book", controller.UnPublishedBook)
	router.HandleFunc("/publishedbook", controller.PublishedBook)
	router.HandleFunc("/admin-review-book", controller.AdminReviewBook)
	router.HandleFunc("/approve-book", controller.ApproveBook)
	router.HandleFunc("/reject", controller.RejectBook)
	router.HandleFunc("/update-book", controller.UpdateBook)
	router.HandleFunc("/view-book", controller.ViewBook)
	router.HandleFunc("/subscribe-book", controller.SubscribeBook)
	router.HandleFunc("/unsubscribe-book", controller.UnsubscribeBook)
	router.HandleFunc("/unpublish-book", controller.UnpublishBook)
	http.Handle("/", router)
	router.HandleFunc("/send-notification", controller.SendNotification)
	router.HandleFunc("/post-notification", controller.PostNotification)
	router.HandleFunc("/user-control", controller.UserControl)

}

var router = mux.NewRouter()

func main() {

	Config := configs.LoadConfiguration("config.json")

	view.Init()

	dbcon.DbConnection(Config)

	HtmlHandlerMux()

	//creating server
	log.Println("Server runing! go to", Config.Server.Host+":"+Config.Server.Port)
	err := http.ListenAndServe(Config.Server.Host+":"+Config.Server.Port, nil)
	if err != nil {
		log.Println(err)
	}
}
