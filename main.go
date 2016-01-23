package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"time"

	"github.com/boltdb/bolt"
	"github.com/julienschmidt/httprouter"
)

// TODO add bcrypt

func init() {
	// Handles db/bucket creation
	db, err := bolt.Open("goblog.db", 0600, nil)

	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists([]byte("UsersBucket"))
		if err != nil {
			return fmt.Errorf("Error with UsersBucket: %s", err)
		}
		return nil
	})

}

func LoginPage(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	baseT := template.Must(template.New("base").Parse(base))
	baseT = template.Must(baseT.Parse(login))

	baseT.ExecuteTemplate(w, "base", map[string]string{
		"PageName": "login",
		"User":     getUser(w, req),
	})
}

func LoginHandler(w http.ResponseWriter, req *http.Request, p httprouter.Params) {
	username := p.ByName("email")
	password := p.ByName("password")

	if verifyUser(w, req, username, password) {
		http.Redirect(w, req, "/admin/", http.StatusFound)
	} else {
		http.Redirect(w, req, "/", http.StatusFound)
	}
}

func LogoutHandler(w http.ResponseWriter, req *http.Request, p httprouter.Params) {
	delete := http.Cookie{Name: "goblog", Value: "blah", Expires: time.Now(), HttpOnly: true, Path: "/"}
	http.SetCookie(w, &delete)
	http.Redirect(w, req, "/", http.StatusFound)
}

func MainPage(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	baseT := template.Must(template.New("base").Parse(base))
	baseT = template.Must(baseT.Parse(mainPage))

	baseT.ExecuteTemplate(w, "base", map[string]string{
		"PageName": "main",
		"User":     getUser(w, r),
	})
}

func SignupPage(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	baseT := template.Must(template.New("base").Parse(base))
	baseT = template.Must(baseT.Parse(signup))

	baseT.ExecuteTemplate(w, "base", map[string]string{
		"PageName": "signup",
		"User":     getUser(w, r),
	})
}

func SignupHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	username := ps.ByName("email")
	password := ps.ByName("password")

	if addUser(username, password) {
		http.Redirect(w, r, "/admin/", http.StatusFound)
	} else {
		http.Redirect(w, r, "/signup/", http.StatusFound)
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
		"User":     getUser(w, r),
	})
}

func verifyUser(w http.ResponseWriter, r *http.Request, username string, password string) bool {
	correctpass := []byte("")
	db, err := bolt.Open("goblog.db", 0600, nil)
	if err != nil {
		fmt.Println(err)
	}
	db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("UsersBucket"))
		correctpass = b.Get([]byte(username))
		return nil
	})
	if password == string(correctpass) {
		cookie := http.Cookie{Name: "goblog", Value: "blah", Expires: time.Now().Add(time.Hour * 24 * 7 * 52), HttpOnly: true, MaxAge: 50000, Path: "/"}
		http.SetCookie(w, &cookie)
		return true
	}
	return false
}

func addUser(username string, password string) bool {
	check := []byte("")
	db, err := bolt.Open("goblog.db", 0600, nil)
	db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("UsersBucket"))
		check = b.Get([]byte(username)) //username
		return nil
	})
	if err != nil {
		fmt.Println(err)
	}
	if len(check) > 2 {
		db.Update(func(tx *bolt.Tx) error {
			b := tx.Bucket([]byte("UsersBucket"))
			err := b.Put([]byte(username), []byte(password))
			return err
		})
		return true
	} else {
		return false
	}
}

func getUser(w http.ResponseWriter, r *http.Request) string {
	cookie, err := r.Cookie("goblog")

	if err == nil {
		if cookie.Value != "" {
			return cookieToUsername(cookie.Value)
		}
	}
	return ""
}

func cookieToUsername(cookie string) string {
	return "zain.hoda@gmail.com"
}

func main() {
	router := httprouter.New()
	router.GET("/", MainPage)
	//router.GET("/login/", LoginPage)
	router.POST("/login/", LoginHandler)
	router.GET("/signup/", SignupPage)
	router.POST("/signup/", SignupHandler)
	router.GET("/admin/", AdminPage)
	router.GET("/logout/", LogoutHandler)
	log.Fatal(http.ListenAndServe(":1337", router))
}
