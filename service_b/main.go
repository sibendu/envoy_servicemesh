// main.go
package main

import (
    "encoding/json"
    "fmt"
    "log"
    "io/ioutil"
    "net/http"

    "github.com/gorilla/mux"
)

// Article - Our struct for all articles
type Author struct {
    Id      string    `json:"Id"`
    Name   string `json:"Name"`
}

var Authors []Author

func homePage(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hello from Service B!")
    fmt.Println("Endpoint Hit: homePage")
}

func returnAllAuthors(w http.ResponseWriter, r *http.Request) {
    fmt.Println("Endpoint Hit: returnAllAuthors")
    json.NewEncoder(w).Encode(Authors)
}

func returnSingleAuthor(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    key := vars["id"]

    for _, author := range Authors {
        if author.Id == key {
            json.NewEncoder(w).Encode(author)
        }
    }
}


func createNewAuthor(w http.ResponseWriter, r *http.Request) {
    reqBody, _ := ioutil.ReadAll(r.Body)
    var author Author 
    json.Unmarshal(reqBody, &author)
    // update our global Articles array to include
    // our new Article
    Authors = append(Authors, author)

    json.NewEncoder(w).Encode(author)
}

func deleteAuthor(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    id := vars["id"]

    for index, author := range Authors {
        if author.Id == id {
            Authors = append(Authors[:index], Authors[index+1:]...)
        }
    }

}

func handleRequests() {
    myRouter := mux.NewRouter().StrictSlash(true)
    myRouter.HandleFunc("/", homePage)
    myRouter.HandleFunc("/authors", returnAllAuthors)
    myRouter.HandleFunc("/authors", createNewAuthor).Methods("POST")
    myRouter.HandleFunc("/authors/{id}", deleteAuthor).Methods("DELETE")
    myRouter.HandleFunc("/authors/{id}", returnSingleAuthor)
    log.Fatal(http.ListenAndServe(":8082", myRouter))
}

func main() {
    Authors = []Author{
        Author{Id: "1", Name: "Paulo"},
        Author{Id: "2", Name: "Gabriel"},
    }
    handleRequests()
}