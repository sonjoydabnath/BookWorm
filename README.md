# BookWorm
a platform where readers can read books and publishers can publish their books

![Linux Build Status](https://img.shields.io/travis/jekyll/jekyll/master.svg?label=Linux%20build)



This is a web app made with GO!


## Installation 

* Test in Linux based Operating System
* Install GO form [here](https://golang.org/)
* Now go to your `GOPATH` or GO workspace 
* Now clone this project `git clone https://github.com/sonjoydabnath/BookWorm.git'
* Build packeges `dbcon`,`model`,`view`,`controller` using `go build` & `go install command`
* Create database schema named `BookWorm` in MysQL and create database from given [sql dump]()
* Set your database `address`, `username` and `password` from `model/dbcon/dbcon.go`
* Run `go run main.go`
* Go to 'http://localhost:8080/' from your browser to check

 ### Admin login credentials are already in use from given database, //will be updated
 * Email: `sonjoy@gmail.com`
 * Password: `sonjoy`
 
 ### Other Libraries used
 
 * `github.com/gorilla/securecookie` 
 * `github.com/gorilla/mux`
 * `github.com/go-sql-driver/mysql`












