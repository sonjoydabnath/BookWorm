package controller

import (
	//"html/template"
	"log"
	"github.com/sonjoydabnath/BookWorm/model"
	"net/http"
	"strconv"
	"github.com/sonjoydabnath/BookWorm/view"
)

//Admin Home page
/*
func Admin(res http.ResponseWriter, req *http.Request) {
	t, _ := template.ParseFiles("HTMLS/admin/admin.html")
	t.Execute(res, nil)
}
*/

//list of  All unpublished book for admin
func UnPublishedBook(res http.ResponseWriter, req *http.Request) {
	log.Println(req.URL.Path)
	userId, userType := getUser(req)
	log.Println("Admin looking for unpublished book = ", userId, userType)
	if userType != "admin" {
		http.Redirect(res, req, "/user-home", 301)
		return
	}

	log.Println("Method :UnpublishedBook in Controller, List of All unpublished book and only admin can View")
	var data model.UData
	data.Books = model.GetBookList(0, 0)
	view.UnPublishedBook(res, req, data)
}

//admin reviewing single book for publishing
func AdminReviewBook(res http.ResponseWriter, req *http.Request) {

	userId, userType := getUser(req)
	log.Println(userId, userType)
	if userType != "admin" {
		http.Redirect(res, req, "/user-home", 301)
		return
	}
	var data model.UData
	var book_id = req.URL.Query().Get("book")
	log.Println("Package : controller , Method : Admin review book, BookId ", book_id)
	bid, _ := strconv.Atoi(book_id)
	book := model.GetBook(bid)
	data.Book1 = book
	view.AdminReviewBook(res, req, data)
}

func ApproveBook(res http.ResponseWriter, req *http.Request) {
	userId, userType := getUser(req)
	log.Println(userId, userType)
	if userType != "admin" {
		http.Redirect(res, req, "/user-home", 301)
		return
	}

	book_id := req.URL.Query().Get("book")

	//	bid, _ = strconv.Atoi(book_id)
	bid, _ := strconv.Atoi(book_id)
	var data model.UData
	data.Book1 = model.GetBook(bid)
	if data.Book1.PubId == bid {
		http.Redirect(res, req, "/un-published-book", 301)
		return
	}
	//
	log.Println("Book to be approved is = " + book_id)

	model.PublishBook(bid, 1)
	http.Redirect(res, req, "/un-published-book", 301)
}

func RejectBook(res http.ResponseWriter, req *http.Request) {
	userId, userType := getUser(req)
	log.Println(userId, userType)
	if userType != "admin" {
		http.Redirect(res, req, "/user-home", 301)
		return
	}
	book_id := req.URL.Query().Get("book")
	log.Println("Book to be rejected is = " + book_id)
	bid, _ := strconv.Atoi(book_id)
	model.PublishBook(bid, 2)
	http.Redirect(res, req, "/un-published-book", 301)
}
func UnpublishBook(res http.ResponseWriter, req *http.Request) {
	userId, userType := getUser(req)
	log.Println(userId, userType)
	if userType != "admin" {
		http.Redirect(res, req, "/user-home", 301)
		return
	}
	book_id := req.URL.Query().Get("book")
	log.Println("Book to be unpublished is = " + book_id)
	bid, _ := strconv.Atoi(book_id)
	model.PublishBook(bid, 0)
	model.UnSubForAll(bid)
	http.Redirect(res, req, "/publishedbook", 301)
}

func SendNotification(res http.ResponseWriter, req *http.Request) {
	userId, userType := getUser(req)
	log.Println(userId, userType)
	if userType != "admin" {
		http.Redirect(res, req, "/user-home", 301)
		return
	}
	bookId := req.URL.Query().Get("book")
	var data model.UData
	bid, _ := strconv.Atoi(bookId)
	data.Book1 = model.GetBook(bid)
	view.SendNoti(res, req, data)
}

func PostNotification(res http.ResponseWriter, req *http.Request) {
	userId, userType := getUser(req)
	log.Println(userId, userType)
	if userType != "admin" {
		http.Redirect(res, req, "/user-home", 301)
		return
	}
	bookId := req.URL.Query().Get("book")
	bid, _ := strconv.Atoi(bookId)
	var nd model.Notification
	nd.BookId = bid
	nd.AdminNotification = req.FormValue("noti")
	model.SendNotification(nd)
	http.Redirect(res, req, "/un-published-book", 301)
}

func UserList(res http.ResponseWriter, req *http.Request) {

	userId, userType := getUser(req)
	log.Println(userId, userType)
	if userType != "admin" {
		http.Redirect(res, req, "/user-home", 301)
		return
	}
	log.Println("Package : Controller ,Method : UserList  Admin ", userId, " Entered to view user list")
	var data model.UData
	//var UL UserList

	data.Users = model.GetUserList()
	view.UserList(res, req, data)
	//t, _ := template.ParseFiles("HTMLS/admin/userlist.html")
	//t.Execute(res, UL)

}

func UserControl(res http.ResponseWriter, req *http.Request) {

	userId, userType := getUser(req)
	log.Println(userId, userType)
	if userType != "admin" {
		http.Redirect(res, req, "/user-home", 301)
		return
	}
	userid := req.URL.Query().Get("userid")
	isBlock := req.URL.Query().Get("doblock")
	var uid, is int
	uid, _ = strconv.Atoi(userid)
	is, _ = strconv.Atoi(isBlock)
	var isb int
	if is == 0 {
		isb = 1
	} else {

		isb = 0
	}

	model.SetActiveUser(uid, isb)
	http.Redirect(res, req, "/user-list", 301)
}
