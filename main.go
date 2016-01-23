package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func LoginPage(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	baseT := template.Must(template.New("base").Parse(base))
	baseT = template.Must(baseT.Parse(login))

	baseT.ExecuteTemplate(w, "base", map[string]string{
		"PageName": "login",
	})
}

func MainPage(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	baseT := template.Must(template.New("base").Parse(base))
	baseT = template.Must(baseT.Parse(mainPage))

	baseT.ExecuteTemplate(w, "base", map[string]string{
		"PageName": "main",
	})
}

func SignupPage(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	username := ps.ByName("username")
	password := ps.ByName("password")

	baseT := template.Must(template.New("base").Parse(base))
	baseT = template.Must(baseT.Parse(signup))

	baseT.ExecuteTemplate(w, "base", map[string]string{
		"PageName": "signup",
	})

	// create new user
	if len(username)+len(password) < 11 {
		fmt.Fprintf(w, "New user"+username+" with pass:"+password+"has signed up")
	} else {
		fmt.Fprintf(w, "Invalid arguments!")
	}
}

func AdminPage(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	// authenticate username and password
	blogname := ps.ByName("blogname")

	if len(blogname) < 20 {
		// create new blog
	}

	baseT := template.Must(template.New("base").Parse(base))
	baseT = template.Must(baseT.Parse(admin))

	baseT.ExecuteTemplate(w, "base", map[string]string{
		"PageName": "admin",
	})
}

func main() {
	router := httprouter.New()
	router.GET("/", MainPage)
	router.GET("/login/", LoginPage)
	router.GET("/signup/", SignupPage)
	router.GET("/admin/", AdminPage)
	log.Fatal(http.ListenAndServe(":1337", router))
}
