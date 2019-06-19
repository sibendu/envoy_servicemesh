package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"

	"net/http"
)

// Article - Our struct for all articles
type Author struct {
    Id      string    `json:"Id"`
    Name   string `json:"Name"`
}

// Article - Our struct for all articles
type Article struct {
    Id      string    `json:"Id"`
    Title   string `json:"Title"`
    Desc    string `json:"desc"`
    Content string `json:"content"`
}

// Article - Our struct for all articles
type Book struct {
    Id        string    `json:"Id"`
    Author    Author `json:"Author"`
    Article   Article `json:"Article"`
}


func handler(w http.ResponseWriter, r *http.Request) {

	//fmt.Fprintf(w, "Calling Service B <br>")

	req, err := http.NewRequest("GET", "http://localhost:8788/authors/1", nil)
	if err != nil {
		fmt.Printf("%s", err)
	}

	req.Header.Add("x-request-id", r.Header.Get("x-request-id"))
	req.Header.Add("x-b3-traceid", r.Header.Get("x-b3-traceid"))
	req.Header.Add("x-b3-spanid", r.Header.Get("x-b3-spanid"))
	req.Header.Add("x-b3-parentspanid", r.Header.Get("x-b3-parentspanid"))
	req.Header.Add("x-b3-sampled", r.Header.Get("x-b3-sampled"))
	req.Header.Add("x-b3-flags", r.Header.Get("x-b3-flags"))
	req.Header.Add("x-ot-span-context", r.Header.Get("x-ot-span-context"))

	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		fmt.Printf("%s", err)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("%s", err)
	}
	//fmt.Fprintf(w, string(body))

	var author Author 
        json.Unmarshal(body, &author)

	//fmt.Fprintf(w, "<br> Called service B <br>")

	//fmt.Fprintf(w, "Calling service C <br>")

	req, err = http.NewRequest("GET", "http://localhost:8791/articles/1", nil)
	if err != nil {
		fmt.Printf("%s", err)
	}

	req.Header.Add("x-request-id", r.Header.Get("x-request-id"))
	req.Header.Add("x-b3-traceid", r.Header.Get("x-b3-traceid"))
	req.Header.Add("x-b3-spanid", r.Header.Get("x-b3-spanid"))
	req.Header.Add("x-b3-parentspanid", r.Header.Get("x-b3-parentspanid"))
	req.Header.Add("x-b3-sampled", r.Header.Get("x-b3-sampled"))
	req.Header.Add("x-b3-flags", r.Header.Get("x-b3-flags"))
	req.Header.Add("x-ot-span-context", r.Header.Get("x-ot-span-context"))

	client = &http.Client{}
	resp, err = client.Do(req)

	if err != nil {
		fmt.Printf("%s", err)
	}

	defer resp.Body.Close()
	body, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("%s", err)
	}

	//fmt.Fprintf(w, string(body))

        var article Article 
        json.Unmarshal(body, &article)
	
	var book = Book{"1", author, article}
	
	//fmt.Fprintf(book)
	//fmt.Fprintf(w, "<br> Called service C <br>")

	json.NewEncoder(w).Encode(book)
	
}

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8081", nil))
}
