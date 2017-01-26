package main

import (
    "fmt"
    "log"
    "math/rand"
    "net/http"
    "text/template"
    "time"

    "github.com/gorilla/mux"
)

type Page struct {
    Speed int32
}

func main() {
    router := mux.NewRouter().StrictSlash(true)
    router.HandleFunc("/", Index)
    router.HandleFunc("/todos", TodoIndex)
    router.HandleFunc("/todos/{todoId}", TodoShow)
    router.HandleFunc("/login", login)

    log.Fatal(http.ListenAndServe(":8080", router))
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

func Index(w http.ResponseWriter, r *http.Request) {
    yourTime := turtleMe()
    fmt.Fprintf(w, "Took %v to get to the result", yourTime)
    fmt.Fprintln(w, "Welcome!")
}

func TodoIndex(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintln(w, "Todo Index!")
}

func TodoShow(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    todoId := vars["todoId"]
    fmt.Fprintln(w, "Todo show:", todoId)
}
