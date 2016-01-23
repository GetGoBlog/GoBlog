package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func mainpage(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprintf(w, "Welcome to GoBlog!")
}

func signup(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	username := ps.ByName("username")
	password := ps.ByName("password")
	// create new user
	if len(username)+len(password) < 11 {
		fmt.Fprintf(w, "New user"+username+" with pass:"+password+"has signed up")
	} else {
		fmt.Fprintf(w, "Invalid arguments!")
	}
}

func newblog(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	// authenticate username and password
	blogname := ps.ByName("blogname")

	if len(blogname) < 20 {
		// create new blog
	}
}

func main() {
	router := httprouter.New()
	router.GET("/", mainpage)
	router.GET("/new/:username/:password", signup)
	router.GET("/new/:username/:password/:blogname", newblog)
	log.Fatal(http.ListenAndServe(":1337", router))
}
