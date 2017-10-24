package model

import (
	"fmt"
	"log"
	"github.com/sonjoydabnath/BookWorm/model/dbcon"
	"strconv"
)

func SetUser(user User) {
	log.Println("User ", user.Name)
	_, err := dbcon.Db.Exec("INSERT INTO user_info (user_id, email, password, name, is_active, user_type) VALUES (?, ?, ?, ?, ?, ?)", user.UserId, user.Email, user.Password, user.Name, user.IsActive, user.UserType)
	if err != nil {
		log.Print(err)
	}

}
func GetUserById(uid int) User {
	var user User
	err := dbcon.Db.QueryRow("SELECT * FROM user_info WHERE user_id=?", uid).Scan(&user.UserId, &user.Email, &user.Password, &user.Name, &user.IsActive, &user.UserType)
	if err != nil {
		log.Println(err)
	}
	log.Println(" UserInfo : ", user.Email, " ", user.Name)
	return user
}

func GetUser(Email string) User {
	var user User
	err := dbcon.Db.QueryRow("SELECT * FROM user_info WHERE email=?", Email).Scan(&user.UserId, &user.Email, &user.Password, &user.Name, &user.IsActive, &user.UserType)
	log.Println(" UserInfo : ", user.Email, " ", user.Name)
	if err != nil {
		log.Println(err)
	}
	return user
}
func GetUserList() []User {
	var u User
	var ul []User
	rows, err := dbcon.Db.Query("SELECT * FROM user_info WHERE user_type != 'admin'")
	if err != nil {
		log.Println(err)
	}
	for rows.Next() {
		err = rows.Scan(&u.UserId, &u.Email, &u.Password, &u.Name, &u.IsActive, &u.UserType)
		if err != nil {
			log.Println(err)
		} else {
			ul = append(ul, u)
		}
	}
	return ul
}

func GenerateID(TableType int) int {
	var tablename string
	if TableType == 1 {
		tablename = "user_info"
	}
	if TableType == 2 {
		tablename = "Book"
	}
	var xid int
	var sql = "select count(*) from " + tablename
	dbcon.Db.QueryRow(sql).Scan(&xid)
	log.Print("Generate user id: ", xid, tablename)
	xid += 1
	return xid
}

/*Get a List of book in a Array of BookP
bookType - (1/0) ->> (published/unpublished)
pubId - 0 ->> for no specific publisher_id but any publisher_id
pubId - greater than 0 and also matching publisher_id ->> for specific publishers book
*/

func GetBookList(bookType int, pubId int) []BookP {

	sql := "select * from Book where is_published = " + strconv.Itoa(bookType)
	if pubId != 0 {
		sql += " AND publisher_id = " + strconv.Itoa(pubId)
	}
	var bookArray []BookP
	rows, err := dbcon.Db.Query(sql)
	if err != nil {
		log.Println(err)
	}
	defer rows.Close()
	for rows.Next() {
		var TmpBook BookP
		err := rows.Scan(&TmpBook.BookId, &TmpBook.PubId, &TmpBook.Title, &TmpBook.Description, &TmpBook.Cover, &TmpBook.Isbn, &TmpBook.Pdf, &TmpBook.IsPubed, &TmpBook.AvrgRating)
		if err != nil {
			log.Println(err)
		}
		//to get Publisher name from user_info Table using PubId from Book Table
		sql = "Select name from user_info Where user_info.user_id=" + strconv.Itoa(TmpBook.PubId)
		dbcon.Db.QueryRow(sql).Scan(&TmpBook.PubName)

		bookArray = append(bookArray, TmpBook)
		log.Println("GetBookList in Model controller : ID ", TmpBook.BookId, " book name ", TmpBook.Title)
	}
	err = rows.Err()
	if err != nil {
		log.Println(err)
	}
	return bookArray
}

func GetBookListOrderBy(bookType int, pubId int, orderBy string) []BookP {

	sql := "select * from Book where is_published = " + strconv.Itoa(bookType)
	if pubId != 0 {
		sql += " AND publisher_id = " + strconv.Itoa(pubId)
		log.Println("Publisher Id = ", pubId)
	}

	if orderBy == "Rating" {
		sql += " ORDER BY Average_rating DESC, Title ASC"
	} else if orderBy == "Title" {
		sql += " ORDER BY Title ASC, Average_rating DESC"
	}

	var bookArray []BookP
	rows, err := dbcon.Db.Query(sql)
	if err != nil {
		log.Println(err)
	}
	defer rows.Close()
	for rows.Next() {
		var TmpBook BookP
		err := rows.Scan(&TmpBook.BookId, &TmpBook.PubId, &TmpBook.Title, &TmpBook.Description, &TmpBook.Cover, &TmpBook.Isbn, &TmpBook.Pdf, &TmpBook.IsPubed, &TmpBook.AvrgRating)
		if err != nil {
			log.Println(err)
		}
		//to get Publisher name from user_info Table using PubId from Book Table
		sql = "Select name from user_info Where user_info.user_id=" + strconv.Itoa(TmpBook.PubId)
		dbcon.Db.QueryRow(sql).Scan(&TmpBook.PubName)

		bookArray = append(bookArray, TmpBook)
		///log.Println("GetBookList in Model controller : ID ", TmpBook.BookId, " book name ", TmpBook.Title)
	}
	err = rows.Err()
	if err != nil {
		log.Println(err)
	}
	return bookArray
}

func GetBookByKeyword(key string) []BookP {
	sql := "SELECT book_id, publisher_id, Title, cover_photo, Average_rating, name from Book, user_info WHERE ( Book.is_published = 1 AND Book.publisher_id = user_info.user_id ) AND (Book.Title LIKE '%" + key + "%' OR user_info.name LIKE '%" + key + "%' OR Book.description Like '%" + key + "%')"
	rows, _ := dbcon.Db.Query(sql)
	var tbook BookP
	var books []BookP
	for rows.Next() {
		rows.Scan(&tbook.BookId, &tbook.PubId, &tbook.Title, &tbook.Cover, &tbook.AvrgRating, &tbook.PubName)
		books = append(books, tbook)
	}
	return books
}

func GetBook(bookId int) BookP {
	var b BookP
	sql := "select * from Book where book_id = " + strconv.Itoa(bookId)
	dbcon.Db.QueryRow(sql).Scan(&b.BookId, &b.PubId, &b.Title, &b.Description, &b.Cover, &b.Isbn, &b.Pdf, &b.IsPubed, &b.AvrgRating)
	sql = "Select name from user_info Where user_info.user_id=" + strconv.Itoa(b.PubId)
	dbcon.Db.QueryRow(sql).Scan(&b.PubName)
	return b
}
func SetRating(bookId int) {
	var avgrating string
	dbcon.Db.QueryRow("SELECT avg(rating) FROM rating_review where book_id=?", bookId).Scan(&avgrating)
	dbcon.Db.Exec("UPDATE Book SET Average_rating=? WHERE book_id=?", avgrating, bookId)
}
func GetBookByIsbn(isbn string) BookP {
	var b BookP
	sql := "select * from Book where Isbn = " + isbn
	dbcon.Db.QueryRow(sql).Scan(&b.BookId, &b.PubId, &b.Title, &b.Description, &b.Cover, &b.Isbn, &b.Pdf, &b.IsPubed, &b.AvrgRating)
	//sql = "Select name from user_info Where user_info.user_id=" + strconv.Itoa(b.PubId)
	//dbcon.Db.QueryRow(sql).Scan(&b.PubName)
	return b
}

func SetBook(book Book) {
	_, err := dbcon.Db.Exec("INSERT INTO  Book (book_id, publisher_id, Title, description, cover_photo, Isbn, pdf, is_published, Average_rating) VALUES (?, ?, ?,? , ?, ?, ?, ?, ?)", book.BookId, book.PubId, book.Title, book.Description, book.Cover, book.Isbn, book.Pdf, book.IsPubed, book.AvrgRating)
	if err != nil {
		log.Println(err)
	}
}

func UpdateBookTitle(bookId int, bookTitle string) {
	dbcon.Db.Exec("UPDATE Book SET Title=? WHERE book_id=?", bookTitle, bookId)
}

func UpdateBookDescription(bookId int, bookDescrptn string) {
	dbcon.Db.Exec("UPDATE Book SET description=? WHERE book_id=?", bookDescrptn, bookId)
}

/*publish, unpublish, rejec a book with id of bookId
isPub = 0 >> unpublish
isPub = 1 >> publish
isPub = 2 >> reject
*/
func PublishBook(bookId int, isPub int) {
	dbcon.Db.Exec("UPDATE Book SET is_published=? WHERE book_id=?", isPub, bookId)
}

func SubScripeBook(bookid int, userid int) {

	var cnt int
	dbcon.Db.QueryRow("SELECT COUNT(*) FROM subscription WHERE book_id=? AND user_id=?", bookid, userid).Scan(&cnt)
	if cnt != 0 {
		log.Println("Method : SubScripeBook, Already exist connection between book_id : ", bookid, " and user_id : ", userid)

		return
	}
	log.Println("Method : SubScripeBook, ", userid, " want to subscribe book, ", bookid)

	dbcon.Db.QueryRow("SELECT COUNT(*) FROM subscription WHERE  user_id=?", userid).Scan(&cnt)
	if cnt >= 3 {
		fmt.Println("Already subscribed for 3 books")
		return
	}
	dbcon.Db.Exec("INSERT INTO subscription (book_id, user_id) VALUES (?,?)", bookid, userid)

}

func CheckSub(userid int, bookid int) int {
	var cnt int
	dbcon.Db.QueryRow("SELECT COUNT(*) FROM subscription WHERE book_id=? AND user_id=?", bookid, userid).Scan(&cnt)
	return cnt
}

func UnsubscribeBook(bookid int, userid int) {
	log.Println("Method : UnsubsCribe Book  user id : ", userid, " and book id: ", bookid)
	dbcon.Db.Exec("DELETE FROM subscription WHERE book_id=? AND user_id =?", bookid, userid)

}

func UnSubForAll(bookId int) {

	_, err := dbcon.Db.Exec("DELETE from subscription where book_id = ?", bookId)
	if err != nil {
		log.Println(err)
	}
}

func SubscriptionList(userid int) []BookP {

	log.Println("Method : Subscription List, User Id : ", userid)
	rows, err := dbcon.Db.Query("SELECT book_id FROM subscription WHERE  user_id=?", userid)

	if err != nil {

		log.Println(err)
	}
	var BL []BookP
	for rows.Next() {
		var bookid int
		err := rows.Scan(&bookid)
		if err != nil {

			log.Println(err)
		}
		BL = append(BL, GetBook(bookid))
	}
	return BL
}

// admin make active and unactive to a user using userid
// active->0 admin will block a user
// active->1 admin will unblock a user
func SetActiveUser(userid int, active int) {
	log.Println("Method SetActiveUser; User id : ", userid, " active: ", active)
	var sql = "UPDATE user_info SET is_active=" + strconv.Itoa(active) + " WHERE user_id=?"
	dbcon.Db.Exec(sql, userid)
}

func SetRatingReview(RatingReviewData RatingReview) {
	dbcon.Db.Exec("INSERT INTO rating_review(user_id,book_id,rating,review) VALUES(?,?,?,?)", RatingReviewData.UserId, RatingReviewData.BookId, RatingReviewData.Rating, RatingReviewData.Review)
}

func GetRatingReview(bookId int) []RatingReview {
	var rdata []RatingReview
	var tmpRat RatingReview
	rows, _ := dbcon.Db.Query("Select rating_review.user_id, book_id, rating, review, name from rating_review, user_info where book_id = ? AND rating_review.user_id = user_info.user_id", bookId)
	for rows.Next() {
		rows.Scan(&tmpRat.UserId, &tmpRat.BookId, &tmpRat.Rating, &tmpRat.Review, &tmpRat.UserName)
		rdata = append(rdata, tmpRat)
	}
	return rdata
}

func SendNotification(NotificationData Notification) {
	dbcon.Db.Exec("INSERT INTO notification_table(book_id,notification) VALUES(?,?)", NotificationData.BookId, NotificationData.AdminNotification)

}
