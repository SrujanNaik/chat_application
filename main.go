package main

import (
	"fmt"
	"html/template"
	"net/http"
)

var templates = template.Must(template.ParseGlob("templates/*.html"))


func loginPage(w http.ResponseWriter, r *http.Request) {
	templates.ExecuteTemplate(w, "login.html", nil)
}

func loginHandler(w http.ResponseWriter, r * http.Request) {
	if r.Method != http.MethodPost {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	http.Redirect(w, r, "/chats.html", http.StatusSeeOther)
}

func dashboardPage(w http.ResponseWriter, r *http.Request){
	templates.ExecuteTemplate(w, "chats.html", nil)
}

func main() {
	http.Handle("/static/",
	    http.StripPrefix("/static/",
		http.FileServer(http.Dir("static")),
	    ),
	)


	http.HandleFunc("/", loginPage)
	// http.HandleFunc("/login", loginHandler)
	http.HandleFunc("/chats", dashboardPage)

	fmt.Println("Started...")
	fmt.Println("http://localhost:8080/")
	http.ListenAndServe(":8080", nil)
}
