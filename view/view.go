package view

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"github.com/sonjoydabnath/BookWorm/model"
	"net/http"
	"strings"
)

var templates *template.Template
var err error

func Init() {
	var allFiles []string
	files, err := ioutil.ReadDir("./templates")
	if err != nil {
		fmt.Println(err)
	}
	for _, file := range files {
		filename := file.Name()
		if strings.HasSuffix(filename, ".html") {
			allFiles = append(allFiles, "./templates/"+filename)
		}
	}
	templates, err = template.ParseFiles(allFiles...)
	if err != nil {
		log.Println(err)
	}
}

func Home(res http.ResponseWriter, req *http.Request, data model.UData) {
	t := templates.Lookup("home.html")
	t.ExecuteTemplate(res, "home", data)
}

func SignUp(res http.ResponseWriter, req *http.Request, data model.UData) {
	log.Println("SignUp Method In View controller")
	t := templates.Lookup("signup.html")
	t.ExecuteTemplate(res, "signup", data)
}
func Login(res http.ResponseWriter, req *http.Request, data model.UData) {
	log.Println("serve the login page please!")
	t3 := templates.Lookup("login.html")
	t3.ExecuteTemplate(res, "login", data)
}

func SignOut(res http.ResponseWriter, req *http.Request, data interface{}) {
	log.Println("Log out view! You have been logged out")
	fmt.Fprint(res, "Logged out!")
}

func UserHome(res http.ResponseWriter, req *http.Request, data model.UData) {
	t := templates.Lookup("user-home.html")
	log.Println("Now I will serve user home = ", data.User1.Name)

	if data.User1.UserType == "admin" {
		t.ExecuteTemplate(res, "admin-home", data)
	} else if data.User1.UserType == "publisher" {
		t.ExecuteTemplate(res, "publisher-home", data)
	} else if data.User1.UserType == "member" {
		t.ExecuteTemplate(res, "member-home", data)
	}
}
func PublishedBook(res http.ResponseWriter, req *http.Request, data model.UData) {
	log.Println("Published Booklist View serving")
	t := templates.Lookup("booklist.html")
	err := t.ExecuteTemplate(res, "book-list", data)
	if err != nil {
		log.Println(err)
	}
}

func UnPublishedBook(res http.ResponseWriter, req *http.Request, data model.UData) {

	log.Println("Package : view, method : UnpublishedBook")
	//log.Println(data.Books.BooKId)

	t := templates.Lookup("un-published-book.html")
	t.ExecuteTemplate(res, "un-published-book", data)
}

func MyPublishedBook(res http.ResponseWriter, req *http.Request, data model.UData) {
	log.Println("Controller :  view  ,   Method: MyPublishedBook  ")
	t := templates.Lookup("my-published-book.html")
	t.ExecuteTemplate(res, "my-published-book", data)

}
func MyUnPublishedBook(res http.ResponseWriter, req *http.Request, data model.UData) {
	log.Println("Package : view, Method : MyUnPublishedBook ")
	t := templates.Lookup("my-un-published-book.html")
	t.ExecuteTemplate(res, "my-un-published-book", data)

}
func PublishNewBook(res http.ResponseWriter, req *http.Request, data model.UData) {
	t := templates.Lookup("publish-new-book.html")
	t.ExecuteTemplate(res, "publish-new-book", data)
}
func AdminReviewBook(res http.ResponseWriter, req *http.Request, data model.UData) {
	log.Println("Package : view , Method : Admin review book, BookId ")
	t := templates.Lookup("adminreviewbook.html")
	t.ExecuteTemplate(res, "adminreviewbook", data)
}
func UpdateBook(res http.ResponseWriter, req *http.Request, data model.UData) {
	t := templates.Lookup("update-book.html")
	t.ExecuteTemplate(res, "update-book", data)
}

func SendNoti(res http.ResponseWriter, req *http.Request, data model.UData) {
	t := templates.Lookup("adminreviewbook.html")
	t.ExecuteTemplate(res, "send-noti", data)
}
func UserList(res http.ResponseWriter, req *http.Request, data model.UData) {
	log.Println("Package : View ,Method : UserList  Admin  Entered to view user list")
	t := templates.Lookup("user-list.html")
	t.ExecuteTemplate(res, "user-list", data)

}
func ViewBook(res http.ResponseWriter, req *http.Request, data model.ViewBookData) {
	log.Println("Package : View ,Method : Viewbook", data)
	t := templates.Lookup("view-book.html")
	t.ExecuteTemplate(res, "view-book", data)

}

func ReadBook(res http.ResponseWriter, req *http.Request, data model.ViewBookData) {
	t := templates.Lookup("read-book.html")
	t.ExecuteTemplate(res, "read-book", data)
}
