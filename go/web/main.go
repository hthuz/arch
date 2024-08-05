package main

import (
	"fmt"
	"html/template"
	"net/http"
	"time"
)

var (
	templates     = template.Must(template.ParseFiles("login.html", "register.html", "home.html"))
	users         = map[string]string{} // mock user storage, username -> password
	sessionStore  = map[string]string{} // mock session storage, sessionID -> username
	cookieName    = "session_id"
	sessionLength = time.Hour * 1
)

func main() {
	http.HandleFunc("/register", registerHandler)
	http.HandleFunc("/login", loginHandler)
	http.HandleFunc("/home", homeHandler)
	http.HandleFunc("/logout", logoutHandler)

	fmt.Println("Starting server at :8080")
	http.ListenAndServe(":8080", nil)
}

func registerHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		username := r.FormValue("username")
		password := r.FormValue("password")
		if _, exists := users[username]; exists {
			templates.ExecuteTemplate(w, "register.html", "User already exists")
			return
		}
		users[username] = password
		templates.ExecuteTemplate(w, "register.html", "Registration successful. Please log in.")
		return
	}
	templates.ExecuteTemplate(w, "register.html", nil)
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		username := r.FormValue("username")
		password := r.FormValue("password")
		if storedPassword, exists := users[username]; exists && storedPassword == password {
			sessionID := fmt.Sprintf("%d", time.Now().UnixNano())
			sessionStore[sessionID] = username
			http.SetCookie(w, &http.Cookie{
				Name:    cookieName,
				Value:   sessionID,
				Expires: time.Now().Add(sessionLength),
			})
			http.Redirect(w, r, "/home", http.StatusSeeOther)
			return
		}
		templates.ExecuteTemplate(w, "login.html", "Invalid username or password")
		return
	}
	cookie, err := r.Cookie(cookieName)
	if err == nil && sessionStore[cookie.Value] != "" {
		http.Redirect(w, r, "/home", http.StatusSeeOther)
		return
	}
	templates.ExecuteTemplate(w, "login.html", nil)
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie(cookieName)
	if err != nil || sessionStore[cookie.Value] == "" {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}
	username := sessionStore[cookie.Value]
	templates.ExecuteTemplate(w, "home.html", username)
}

func logoutHandler(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie(cookieName)
	if err == nil {
		delete(sessionStore, cookie.Value)
		http.SetCookie(w, &http.Cookie{
			Name:   cookieName,
			Value:  "",
			MaxAge: -1,
		})
	}
	http.Redirect(w, r, "/login", http.StatusSeeOther)
}
