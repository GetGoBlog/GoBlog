package main

import (
	"html/template"
	"log"
	"net/http"
)

func LoginPage(w http.ResponseWriter, req *http.Request) {
	baseT := template.Must(template.New("base").Parse(base))
	baseT = template.Must(baseT.Parse(login))

	baseT.ExecuteTemplate(w, "base", map[string]string{
		"PageName": "login",
	})
}

func SignupPage(w http.ResponseWriter, req *http.Request) {
	baseT := template.Must(template.New("base").Parse(base))
	baseT = template.Must(baseT.Parse(signup))

	baseT.ExecuteTemplate(w, "base", map[string]string{
		"PageName": "signup",
	})
}

func AdminPage(w http.ResponseWriter, req *http.Request) {
	baseT := template.Must(template.New("base").Parse(base))
	baseT = template.Must(baseT.Parse(admin))

	baseT.ExecuteTemplate(w, "base", map[string]string{
		"PageName": "admin",
	})
}

func MainPage(w http.ResponseWriter, req *http.Request) {
	baseT := template.Must(template.New("base").Parse(base))
	baseT = template.Must(baseT.Parse(mainPage))

	baseT.ExecuteTemplate(w, "base", map[string]string{
		"PageName": "main",
	})
}

func main() {
	http.HandleFunc("/login", LoginPage)
	http.HandleFunc("/admin", AdminPage)
	http.HandleFunc("/signup", SignupPage)
	http.HandleFunc("/", MainPage)
	err := http.ListenAndServe(":1337", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
