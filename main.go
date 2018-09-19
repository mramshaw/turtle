package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"os"
	"text/template"
	"time"

	"github.com/gorilla/mux"
)

// Page represents the web page
type Page struct {
	Speed int32
}

func main() {
	port := os.Getenv("PORT")

	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", index)
	router.HandleFunc("/todos", todoIndex)
	router.HandleFunc("/todos/{todoID}", todoShow)
	router.HandleFunc("/login", login)

	log.Printf("Now listening on http://localhost:%v ...\n", port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}

func login(w http.ResponseWriter, r *http.Request) {
	speed := turtleMe()
	page := &Page{Speed: speed}

	fmt.Println("method:", r.Method) //get request method
	if r.Method == "GET" {
		t, _ := template.ParseFiles("./resources/sample.html")
		t.Execute(w, page)
	} else {
		r.ParseForm()
		// logic part of log in
		fmt.Println("username:", r.Form["username"])
		fmt.Println("password:", r.Form["password"])
	}
}

func turtleMe() int32 {
	yourTime := rand.Int31n(10)
	time.Sleep(time.Duration(yourTime) * time.Second)
	return yourTime
}

func index(w http.ResponseWriter, r *http.Request) {
	yourTime := turtleMe()
	fmt.Fprintf(w, "It took %v second(s) to get the result.\n\n", yourTime)
	fmt.Fprintln(w, "Welcome!")
}

func todoIndex(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Todo Index!")
}

func todoShow(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	todoID := vars["todoID"]
	fmt.Fprintln(w, "Todo show:", todoID)
}
