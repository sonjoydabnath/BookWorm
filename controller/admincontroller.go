package controller

import (
	//"html/template"
	"log"
	"net/http"
	"strconv"

	"github.com/sonjoydabnath/BookWorm/model"
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
		http.Redirect(res, req, "/user-home", http.StatusFound)
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
		http.Redirect(res, req, "/user-home", http.StatusFound)
		return
	}
	var data model.UData
	var book_id = req.URL.Query().Get("book")
	log.Println("Package : controller , Method : Admin review book, BookId ", book_id)
	bid, _ := strconv.Atoi(book_id)
	book := model.GetBook(bid)
	data.Book1 = book
	if req.Method == http.MethodGet {
		view.AdminReviewBook(res, req, data)
		return
	}

	//POST method
	read := req.FormValue("read")
	if read == "read" {
		//redirect to reading page
		http.Redirect(res, req, "/uploads/Pdf/"+book_id+".pdf", http.StatusFound)
		return
	}
}

func ApproveBook(res http.ResponseWriter, req *http.Request) {
	userId, userType := getUser(req)
	log.Println(userId, userType)
	if userType != "admin" {
		http.Redirect(res, req, "/user-home", http.StatusFound)
		return
	}

	book_id := req.URL.Query().Get("book")

	uid, _ := strconv.Atoi(userId)
	bid, _ := strconv.Atoi(book_id)
	var data model.UData
	data.Book1 = model.GetBook(bid)
	if data.Book1.PubId == uid {
		log.Println("Admin can't do this operation on admin's own Book", uid, bid)
		http.Redirect(res, req, "/un-published-book", http.StatusFound)
		return
	}
	//
	log.Println("Book to be approved is = " + book_id)

	model.PublishBook(bid, 1)
	http.Redirect(res, req, "/un-published-book", http.StatusFound)
}

func RejectBook(res http.ResponseWriter, req *http.Request) {
	userId, userType := getUser(req)
	log.Println(userId, userType)
	if userType != "admin" {
		http.Redirect(res, req, "/user-home", http.StatusFound)
		return
	}
	book_id := req.URL.Query().Get("book")
	log.Println("Book to be rejected is = " + book_id)

	uid, _ := strconv.Atoi(userId)
	bid, _ := strconv.Atoi(book_id)
	var data model.UData
	data.Book1 = model.GetBook(bid)
	if data.Book1.PubId == uid {
		log.Println("Admin can't do this operation on admin's own Book", uid, bid)
		http.Redirect(res, req, "/un-published-book", http.StatusFound)
		return
	}
	model.PublishBook(bid, 2)
	http.Redirect(res, req, "/un-published-book", http.StatusFound)
}
func UnpublishBook(res http.ResponseWriter, req *http.Request) {
	userId, userType := getUser(req)
	log.Println(userId, userType)
	if userType != "admin" {
		http.Redirect(res, req, "/user-home", http.StatusFound)
		return
	}
	book_id := req.URL.Query().Get("book")
	log.Println("Book to be unpublished is = " + book_id)
	bid, _ := strconv.Atoi(book_id)
	uid, _ := strconv.Atoi(userId)
	var data model.UData
	data.Book1 = model.GetBook(bid)
	if data.Book1.PubId == uid {
		log.Println("Admin can't do this operation on admin's own Book", uid, bid)
		http.Redirect(res, req, "/un-published-book", http.StatusFound)
		return
	}
	model.PublishBook(bid, 0)
	model.UnSubForAll(bid)
	http.Redirect(res, req, "/publishedbook", http.StatusFound)
}

func SendNotification(res http.ResponseWriter, req *http.Request) {
	userId, userType := getUser(req)
	log.Println(userId, userType)
	if userType != "admin" {
		http.Redirect(res, req, "/user-home", http.StatusFound)
		return
	}
	bookId := req.URL.Query().Get("book")

	uid, _ := strconv.Atoi(userId)
	bid, _ := strconv.Atoi(bookId)
	var data model.UData
	data.Book1 = model.GetBook(bid)
	if data.Book1.PubId == uid {
		log.Println("Admin can't do this operation on admin's own Book", uid, bid)
		http.Redirect(res, req, "/un-published-book", http.StatusFound)
		return
	}
	view.SendNoti(res, req, data)
}

func PostNotification(res http.ResponseWriter, req *http.Request) {
	userId, userType := getUser(req)
	log.Println(userId, userType)
	if userType != "admin" {
		http.Redirect(res, req, "/user-home", http.StatusFound)
		return
	}
	bookId := req.URL.Query().Get("book")
	uid, _ := strconv.Atoi(userId)
	bid, _ := strconv.Atoi(bookId)
	var data model.UData
	data.Book1 = model.GetBook(bid)
	if data.Book1.PubId == uid {
		log.Println("Admin can't do this operation on admin's own Book", uid, bid)
		http.Redirect(res, req, "/un-published-book", http.StatusFound)
		return
	}

	var nd model.Notification
	nd.BookId = bid
	nd.AdminNotification = req.FormValue("noti")
	model.SendNotification(nd)
	http.Redirect(res, req, "/un-published-book", http.StatusFound)
}

func UserList(res http.ResponseWriter, req *http.Request) {

	userId, userType := getUser(req)
	log.Println(userId, userType)
	if userType != "admin" {
		http.Redirect(res, req, "/user-home", http.StatusFound)
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
		http.Redirect(res, req, "/user-home", http.StatusFound)
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
	http.Redirect(res, req, "/user-list", http.StatusFound)
}
