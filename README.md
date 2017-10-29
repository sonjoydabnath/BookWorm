# BookWorm
a platform where readers can read books and publishers can publish their books

![Linux Build Status](https://img.shields.io/badge/Linux%20Build-Pass-green.svg)
![Issues](https://img.shields.io/github/issues/sonjoydabnath/BookWorm.svg)



This is a web app made with GO

## App feature
* Readers can join as member to read books online by subscribing to specific books
* Readers can subscribe at most 3 books at a time
* Publishers can publish their books online but an admin have to approve to book to be published
* Admins can reject/unpublish any book if its necessary
* Admins can also block any specific users if its necessary too
* More feature yet to come
----
NB: UI Needs more development, we are working on that slowly. Actually we've just built the backend to see how golang works!
----

## Installation
* Install GO language on your Linux-based machine form [here](https://golang.org/) and set your GOPATH
* From terminal enter `go get github.com/sonjoydabnath/BookWorm`
* Now enter `cd $GOPATH/src/github.com/sonjoydabnath/BookWorm`
* Edit the `config.json` file with your `host`, `port`, `database.host`, `database.port` `database.schema` `database.username`, `database.password`
* Now run the `backup.sql` in you MySQL Workbench to create the full database
* Finally From terminal enter `go run main.go` and go to browser
-----

 ### Admin login credentials are already in use from given database, \\will be updated
 * Email: `sonjoy@gmail.com`
 * Password: `sonjoy`

 ### Other Libraries used

 * `github.com/gorilla/securecookie`
 * `github.com/gorilla/mux`
 * `github.com/go-sql-driver/mysql`
 -------
