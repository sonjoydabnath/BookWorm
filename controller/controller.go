package controller

import (
	"fmt"
	"html"
	"html/template"
	"io"
	"log"
	"github.com/sonjoydabnath/BookWorm/model"
	"net/http"
	"os"
	"strconv"
	"github.com/sonjoydabnath/BookWorm/view"
)

func Pr() {
	fmt.Println("Hello from Package")
}

func Home(res http.ResponseWriter, req *http.Request) {
	userId, userType := getUser(req)
	log.Println(userId, userType)
	if userType != "" {
		http.Redirect(res, req, "/user-home", 301)
		return
	}
	var data model.UData
	view.Home(res, req, data)
}

//var LoggedInUser model.User //set from session

func Login(res http.ResponseWriter, req *http.Request) {
	clearSession(res)
	log.Println("method login", req.URL.Path, "Method = ", req.Method)
	var data model.UData
	log.Println("Logedin user = " + data.User1.Name)
	//processing GET method
	if req.Method != "POST" {
		userId, userType := getUser(req)
		if userId != "" {
			log.Println("User Id from session = "+userId+"usertype = ", userType)
			uid, err := strconv.Atoi(userId)
			if err == nil {
				log.Println("Logedin user = " + userId)
				if uid > 0 {
					http.Redirect(res, req, "/user-home", 301)
					return
				}
			}
		}
		log.Println("Serving login Page!")
		view.Login(res, req, data)
		return
	}

	//processing POST method
	req.ParseForm()
	email := html.EscapeString(req.FormValue("email"))
	password := html.EscapeString(req.FormValue("password"))
	log.Println("User Login Attempt by: ", email, " ", password)
	var user model.User
	user = model.GetUser(email)

	if user.Email != email {
		log.Println("User not found")
		data.Message = "Invalid Email!"
		view.Login(res, req, data)
		return
	}
	if user.Password != password {
		log.Println("Password does not match")
		data.Message = "Incorrect Password!!"
		view.Login(res, req, data)
		return
	}

	//if user is blocked redirect him
	if user.IsActive == 0 {
		log.Println("User is blocked")
		data.Message = "User is Blocked!"
		view.Login(res, req, data)
		return
	}

	//Set Session for newly loggedIn user here****
	//**********************************************
	//	LoggedInUser = user
	uid := strconv.Itoa(user.UserId)
	setSession(uid, user.UserType, res)
	log.Println("Welcome success login id = " + uid + " Name = " + user.Name)
	//redirect according to user type
	data.Message = "Welcome " + user.Name + "! Login Succesful!!"
	data.User1 = user
	http.Redirect(res, req, "/user-home", 301) // redirect to user home(admin/pub/member)
}

func Logout(res http.ResponseWriter, req *http.Request) {
	clearSession(res)
	clearSession(res)
	http.Redirect(res, req, "/", 302)
}
func UserHome(res http.ResponseWriter, req *http.Request) {
	var data model.UData
	log.Println(req.URL.Path, "Method = ", req.Method)
	userId, userType := getUser(req)
	log.Println(userId, userType)
	if userId == "" {
		http.Redirect(res, req, "/login", 301)
		return
	}
	uid, _ := strconv.Atoi(userId)
	if uid == 0 {
		http.Redirect(res, req, "/login", 301)
		return
	}

	data.User1 = model.GetUserById(uid)
	if data.User1.UserType != "admin" {
		data.Books = model.SubscriptionList(data.User1.UserId)
	}
	data.Message = "Welcome " + data.User1.Name + "!!"
	view.UserHome(res, req, data)
	return
}

func SignUp(res http.ResponseWriter, req *http.Request) {

	userId, userType := getUser(req)
	log.Println(userId, userType)
	if userType != "" {
		http.Redirect(res, req, "/user-home", 301)
		return
	}

	var data model.UData
	log.Println("Entered Method : SignUp")
	//before clicking submit option
	if req.Method != "POST" {
		view.SignUp(res, req, data)
		return
	}

	//getting signup information
	req.ParseForm()
	name := html.EscapeString(req.FormValue("name"))
	email := html.EscapeString(req.FormValue("email"))
	password1 := html.EscapeString(req.FormValue("password1"))
	password2 := html.EscapeString(req.FormValue("password2"))
	usertype := html.EscapeString(req.FormValue("UserType"))
	log.Println("Name ", name, "Email ", email, "password1 ", password1, "password2 ", password2, " Type ", usertype)

	//matching password for confirmation
	if password1 != password2 {
		log.Println("Password does not match")
		data.Message = "Password does not match"
		view.SignUp(res, req, data)
		//http.Redirect(res, req, "/signup", 301)
		return
	}
	//checking mail used or not
	var emailexist string
	var user model.User
	user = model.GetUser(email)
	emailexist = user.Email
	if emailexist == email {
		log.Println("Email already used")
		data.Message = "Email already used"
		view.SignUp(res, req, data)
		return
	}
	//generating unique user id
	var user_id int
	user_id = model.GenerateID(1)
	user.Set(user_id, email, password1, name, 1, usertype)
	model.SetUser(user)
	println("Sign Up successfull ", user_id)
	//storing new user in database user tab
	println("Stored in database")
	http.Redirect(res, req, "/login", 301)
}

func Contact(res http.ResponseWriter, req *http.Request) {
	t, _ := template.ParseFiles("HTMLS/contact.html")
	t.Execute(res, nil)
}
func About(res http.ResponseWriter, req *http.Request) {
	t, _ := template.ParseFiles("HTMLS/about.html")
	t.Execute(res, nil)
}

func PublishedBook(res http.ResponseWriter, req *http.Request) {
	log.Println(req.URL.Path, " Method = ", req.Method)
	userId, userType := getUser(req)
	log.Println(userId, userType)
	if userId == "" {
		http.Redirect(res, req, "/login", 301)
		return
	}
	var data model.UData
	pid := req.URL.Query().Get("pid")
	p, _ := strconv.Atoi(pid)

	if req.Method == http.MethodGet {
		//finding unpublished book id from 	database
		data.Books = model.GetBookList(1, p) // 1 - publishedbook, 0 - No specific user
		//fmt.Fprint(res, "hello")
		//log.Println(data.Books)
		view.PublishedBook(res, req, data)
		return
	}
	sortBy := req.FormValue("Sortby")
	keyword := req.FormValue("Keyword")
	log.Println("Now books will be filtered by pub id = ", pid, p)
	log.Println("Now Books will be sorted by " + sortBy + " Search Key = " + keyword)
	if keyword == "" {
		data.Books = model.GetBookListOrderBy(1, p, sortBy) // 1 - publishedbook, 0 - No specific user
	} else {
		//search database by keword
		data.Books = model.GetBookByKeyword(keyword)
		//log.Println(data.Books)
	}
	//fmt.Fprint(res, "hello")
	log.Println(data.Books)
	data.Message = sortBy
	view.PublishedBook(res, req, data)
	return
}

func MyPublishedBook(res http.ResponseWriter, req *http.Request) {

	log.Println(req.URL.Path)
	log.Println("MyPublishedBook() - method = ", req.Method)
	userId, userType := getUser(req)
	log.Println("Admin looking for unpublished book = ", userId, userType)
	if userType == "member" {
		http.Redirect(res, req, "/user-home", 301)
		return
	}
	if userType == "" {
		http.Redirect(res, req, "/login", 301)
		return
	}
	//
	//
	var data model.UData
	var BL model.BookList
	log.Println("Method:MyPublishedBook -> Publisher id = " + userId)
	//Take publisherid(LoggedInUser.UserId) from session
	//finding unpublished book id from 	database
	var uid int
	uid, _ = strconv.Atoi(userId)
	BL.Blist = model.GetBookList(1, uid) // 1 - publishedbook, 0 - No specific user
	data.Books = BL.Blist
	view.MyPublishedBook(res, req, data)
}

func MyUnPublishedBook(res http.ResponseWriter, req *http.Request) {

	log.Println(req.URL.Path)
	userId, userType := getUser(req)
	log.Println("Admin looking for unpublished book = ", userId, userType)
	if userType == "member" {
		http.Redirect(res, req, "/user-home", 301)
		return
	}
	if userType == "" {
		http.Redirect(res, req, "/login", 301)
		return
	}

	log.Println("Package : controller, Method : MyUnPublishedBook ")
	var data model.UData
	//var BL model.BookList
	log.Println("Method:MyUnpublishedBook -> Publisher id = ", userId)
	//Take publisherid(LoggedInUser.UserId) from session
	//finding unpublished book id from 	database
	var uid int
	uid, _ = strconv.Atoi(userId)
	data.Books = model.GetBookList(0, uid) // 1 - publishedbook, 0 - No specific user
	//t, _ := template.ParseFiles("HTMLS/my-unpublished-book.html")
	//t.Execute(res, BL)
	view.MyUnPublishedBook(res, req, data)
}

//publishing a new Book
func PublishNewBook(res http.ResponseWriter, req *http.Request) {

	log.Println(req.URL.Path)
	userId, userType := getUser(req)
	log.Println(" ", userId, userType)
	if userType == "member" {
		http.Redirect(res, req, "/user-home", 301)
		return
	}
	if userType == "" {
		http.Redirect(res, req, "/login", 301)
		return
	}
	fmt.Println("Method:PublisNewBook", req.Method)

	var data model.UData
	if req.Method == "POST" {

		//finding unique book id
		var bid int
		bid = model.GenerateID(2)
		var book_id string
		book_id = strconv.Itoa(bid)
		var uid int
		uid, _ = strconv.Atoi(userId)
		//finding book publisher id
		publisher_id := uid //it is temporary finally session will generat publisher_id

		//finding book title,description and isbn no
		title := req.FormValue("title")
		description := req.FormValue("description")
		isbn := req.FormValue("isbn")
		if title == "" {
			data.Message = "Title can not be null"
			view.PublishNewBook(res, req, data)
			return
		}
		if description == "" {
			data.Message = "Description can not be null"
			view.PublishNewBook(res, req, data)
			return
		}
		if isbn == "" {
			data.Message = "Isbn can not be null"
			view.PublishNewBook(res, req, data)
			return
		}
		//finding book cover_photo and pdf version of book
		file, handler, err := req.FormFile("cover_photo")
		file2, handler2, err2 := req.FormFile("pdf")

		//error checking
		if err != nil {
			fmt.Println(err)
			//	http.Redirect(res, req, "/publish-new-book", 301)
			data.Message = "Problem Uploading Cover photo of book- Empty Cover"
			view.PublishNewBook(res, req, data)

			return
		}
		if err2 != nil {
			fmt.Println(err2)
			//			http.Redirect(res, req, "/publish-new-book", 301)
			data.Message = "Problem Uploading Pdf file of book- Empty Pdf"
			view.PublishNewBook(res, req, data)
			return
		}

		//closing
		defer file.Close()
		defer file2.Close()

		//changing file name
		handler.Filename = book_id + ".jpg"
		handler2.Filename = book_id + ".pdf"
		log.Println("File Name ", handler.Filename)
		log.Println("Pdf Name ", handler2.Filename)

		//saving file to their destination
		f, err := os.OpenFile("."+string(os.PathSeparator)+"uploads"+string(os.PathSeparator)+"CoverPhoto"+string(os.PathSeparator)+handler.Filename, os.O_WRONLY|os.O_CREATE, 0666)
		f2, err2 := os.OpenFile("."+string(os.PathSeparator)+"uploads"+string(os.PathSeparator)+"Pdf"+string(os.PathSeparator)+handler2.Filename, os.O_WRONLY|os.O_CREATE, 0666)
		if err != nil {
			fmt.Println(err)
			//			http.Redirect(res, req, "/publish-new-book", 301)
			data.Message = "Problem Uploading  Cover photo of book"
			view.PublishNewBook(res, req, data)
			return
		}
		if err2 != nil {
			fmt.Println(err2)
			//	http.Redirect(res, req, "/publish-new-book", 301)
			data.Message = "Problem Uploading Pdf file of book"
			view.PublishNewBook(res, req, data)
			return
		}
		defer f.Close()
		defer f2.Close()
		io.Copy(f, file)
		io.Copy(f2, file2)

		fmt.Println("Book id ", bid, "Publisher ", publisher_id, " title ", title, "Description ", description, "ISBN ", isbn)
		//	db.Exec("INSERT INTO Book (Book_id, publisher_id, Title, description, cover_photo, Isbn, pdf, is_published,Average_rating) VALUES (?,?,?,?,?,?,?,?,?)", cnt, publisher_id, title, description, handler.Filename, isbn, 0, 0.0)

		//if isbn number is not uniqu then
		//db.QueryRow("SELECT Isbn  FROM Book WHERE Isbn=?", isbn).Scan(&isbnexist)
		tmpBook := model.GetBookByIsbn(isbn)
		if tmpBook.Isbn == isbn {
			fmt.Println("Isbn already used")
			//http.Redirect(res, req, "/publish-new-book", 301)
			data.Message = "Isbn already used"
			view.PublishNewBook(res, req, data)
			return
		}

		//value updated to database
		//db.Exec("INSERT INTO  Book (book_id, publisher_id, Title, description, cover_photo, Isbn, pdf, is_published, Average_rating) VALUES (?, ?, ?,? , ?, ?, ?, 0, 0)", cnt, publisher_id, title, description, handler.Filename, isbn, handler2.Filename) //, cnt, publisher_id, title, description, handler.Filename, isbn, handler2.Filename, 0, 0)
		var book model.Book
		book.Set(bid, publisher_id, title, description, handler.Filename, isbn, handler2.Filename, 0, 0)
		model.SetBook(book)
		fmt.Println("New Book Store successfully")
		//	http.Redirect(res, req, "/publish-new-book", 301)
		data.Message = "New Book Stored successfully"
		view.PublishNewBook(res, req, data)
	} else {
		//data.Message = ""
		view.PublishNewBook(res, req, data)
	}
}

//publisher update info of his book waiting for admin approval
func UpdateBook(res http.ResponseWriter, req *http.Request) {
	//
	log.Println(req.URL.Path)
	userId, userType := getUser(req)
	log.Println(userId, userType)
	uid, _ := strconv.Atoi(userId)
	var book_id = req.URL.Query().Get("book")
	bid, _ := strconv.Atoi(book_id)
	var TmpBook model.BookP
	TmpBook = model.GetBook(bid)

	if TmpBook.PubId != uid {

		http.Redirect(res, req, "/user-home", 301)
		return
	}
	//access secured

	var data model.UData
	data.Book1 = TmpBook

	if req.Method != http.MethodPost {
		fmt.Println("Method:UpdateBook GET Method, redirect from : /my-unpublished-book")
		view.UpdateBook(res, req, data)
		return
	}

	fmt.Println("Method: UpdateBook  POST Method,  redirect from : /update-book")
	//starting upload cover
	file, handler, err := req.FormFile("cover_photo")
	//error checking
	if err != nil {
		fmt.Println("No cover photo")
		fmt.Println(err)
	} else {
		fmt.Println("New cover photo found :", handler.Filename)
		//closing
		defer file.Close()
		//changing file name
		handler.Filename = book_id + ".jpg"
		fmt.Println("Cover Name ", handler.Filename)
		//saving file to their destination and at first deleting the already existing file
		os.Remove("." + string(os.PathSeparator) + "uploads" + string(os.PathSeparator) + "CoverPhoto" + string(os.PathSeparator) + handler.Filename)
		f, err := os.OpenFile("."+string(os.PathSeparator)+"uploads"+string(os.PathSeparator)+"CoverPhoto"+string(os.PathSeparator)+handler.Filename, os.O_WRONLY|os.O_CREATE, 0666)
		if err != nil {
			fmt.Println(err)
			return
		}
		defer f.Close()
		io.Copy(f, file)
	} //end of uploading cover photo

	//update pdf file
	file, handler, err = req.FormFile("pdf")
	//error checking
	if err != nil {
		fmt.Println("No cover photo")
		fmt.Println(err)

	} else {
		fmt.Println("New pdf found :", handler.Filename)
		defer file.Close()
		//changing file name
		handler.Filename = book_id + ".pdf"
		fmt.Println("Pdf Name ", handler.Filename)
		//saving file to their destination and at first deleting the already existing file
		os.Remove("." + string(os.PathSeparator) + "uploads" + string(os.PathSeparator) + "Pdf" + string(os.PathSeparator) + handler.Filename)
		f, err := os.OpenFile("."+string(os.PathSeparator)+"uploads"+string(os.PathSeparator)+"Pdf"+string(os.PathSeparator)+handler.Filename, os.O_WRONLY|os.O_CREATE, 0666)

		if err != nil {
			fmt.Println(err)
			return
		}
		defer f.Close()
		io.Copy(f, file)
	} //end of uploading pdf file

	title := req.FormValue("title")
	if title != "" {
		fmt.Println(" title update : ", title)
		model.UpdateBookTitle(bid, title)
	}

	description := req.FormValue("description")
	if description != "" {
		fmt.Println(" description update : ", description)
		model.UpdateBookDescription(bid, description)
	}
	view.UpdateBook(res, req, data)

}

func ViewBook(res http.ResponseWriter, req *http.Request) {
	userId, userType := getUser(req)
	log.Println(userId, userType)
	if userType == "" {
		http.Redirect(res, req, "/login", 301)
		return
	}
	uid, _ := strconv.Atoi(userId)
	var book_id = req.URL.Query().Get("book")
	fmt.Println("Requested book ID : ", book_id)
	bid, _ := strconv.Atoi(book_id)
	var data model.ViewBookData
	data.Book1 = model.GetBook(bid)

	//---------button control------------
	if userType == "admin" {
		data.Unpub = 1
		data.Read = 1
	} else if userType == "publisher" {
		if (model.CheckSub(uid, bid) == 1) || (data.Book1.PubId == uid) {
			data.Read = 1
		}
		if (data.Book1.PubId != uid) && (model.CheckSub(uid, bid) == 1) {
			data.Unsub = 1
		}
		if (data.Book1.PubId != uid) && (model.CheckSub(uid, bid) == 0) {
			data.Sub = 1
		}
	} else if userType == "member" {
		if model.CheckSub(uid, bid) == 1 {
			data.Read = 1
			data.Unsub = 1
		}
		if model.CheckSub(uid, bid) == 0 {
			data.Sub = 1
		}
	}
	log.Println("Button permission form userId ", userId, " Read,Sub,Unsub,Unpub", data.Read, data.Sub, data.Unsub, data.Unpub)
	//-------------button control done

	//fmt.Println("Single book view ViewBook.go")
	//GET method handle
	if req.Method == http.MethodGet {
		data.RatRev = model.GetRatingReview(bid)
		log.Println("GET -> Rating Review = ", data.RatRev)
		//log.Println(data.RatRev)
		view.ViewBook(res, req, data)
		return
	}
	//End of GET method

	//--------POST method Handle
	unp := req.FormValue("unpub")
	read := req.FormValue("read")
	sub := req.FormValue("sub")
	unsub := req.FormValue("unsub")

	if userType == "admin" {
		if unp == "unpub" {
			log.Println("unpublishing bookid = ", bid, data)
			model.PublishBook(bid, 0)
			data.Unpub = 0
			http.Redirect(res, req, "/publishedbook", 301)
			return
		} else if read == "read" {
			//redirect to reading page
			http.Redirect(res, req, "/uploads/Pdf/"+book_id+".pdf", 301)
			return
		}
	} else if userType == "publisher" {
		if sub == "sub" {
			model.SubScripeBook(bid, uid)
			data.Sub = 0
			data.Unsub = 1
			data.Read = 1
		} else if unsub == "unsub" {
			model.UnsubscribeBook(bid, uid)
			data.Unsub = 0
			data.Sub = 1
			data.Read = 0
		} else if read == "read" {
			//redirec to reading  page
			http.Redirect(res, req, "/uploads/Pdf/"+book_id+".pdf", 301)
			return
		}
	} else if userType == "member" {
		if sub == "sub" {
			model.SubScripeBook(bid, uid)
			data.Sub = 0
			data.Unsub = 1
			data.Read = 1
		} else if unsub == "unsub" {
			model.UnsubscribeBook(bid, uid)
			data.Unsub = 0
			data.Sub = 1
			data.Read = 0
		}
		if read == "read" {
			//redirect to reading page
			http.Redirect(res, req, "/uploads/Pdf/"+book_id+".pdf", 301)
			return
		}
	}

	reviewButton := req.FormValue("review-button")
	if reviewButton == "review-button" {
		rev := req.FormValue("review")
		rat := req.FormValue("rating")
		log.Println("Rating and reviw posted by UserId ", uid, " | rating = ", rat, " and review = ", rev)
		var ratrevdata model.RatingReview
		ratrevdata.BookId = bid
		ratrevdata.UserId = uid
		ratrevdata.UserName = model.GetUserById(uid).Name
		rt, _ := strconv.Atoi(rat)
		ratrevdata.Rating = float32(rt)
		ratrevdata.Review = rev
		model.SetRatingReview(ratrevdata)
	}

	data.RatRev = model.GetRatingReview(bid)
	log.Println("POST -> Rating Review = ", data.RatRev)
	log.Println("Button permission form userId ", userId, " Read,Sub,Unsub,Unpub", data.Read, data.Sub, data.Unsub, data.Unpub)
	view.ViewBook(res, req, data)
}

func SubscribeBook(res http.ResponseWriter, req *http.Request) {

	userId, userType := getUser(req)
	log.Println(userId, userType)

	if userType == "" {
		http.Redirect(res, req, "/login", 301)
		return
	}
	if userType == "admin" {
		http.Redirect(res, req, "/user-home", 301)
		return
	}

	//
	var book_id = req.URL.Query().Get("book")
	var user_id, _ = strconv.Atoi(userId)
	log.Println("Method: SubscribeBook,  Entered with book id : ", book_id, " and  user id: ", user_id)
	var bid int
	bid, _ = strconv.Atoi(book_id)
	model.SubScripeBook(bid, user_id)
	http.Redirect(res, req, "/view-book?book="+book_id, 301)

}

func UnsubscribeBook(res http.ResponseWriter, req *http.Request) {

	userId, userType := getUser(req)
	log.Println(userId, userType)

	if userType == "" {
		http.Redirect(res, req, "/login", 301)
		return
	}
	if userType == "admin" {
		http.Redirect(res, req, "/user-home", 301)
		return
	}
	var book_id = req.URL.Query().Get("book")
	var user_id, _ = strconv.Atoi(userId)
	log.Println("Method: UnubscribeBook,  Entered with book id : ", book_id, " and  user id: ", user_id)
	var bid int
	bid, _ = strconv.Atoi(book_id)
	model.UnsubscribeBook(bid, user_id)
	http.Redirect(res, req, "/view-book?book="+book_id, 301)
}
