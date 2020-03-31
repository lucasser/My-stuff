package main

import (
	"fmt"
	"net/http"
)

func calculator(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		http.ServeFile(w, r, "calculator1.html")
	case "POST":
		// Call ParseForm() to parse the raw query and update r.PostForm and r.Form.
		if err := r.ParseForm(); err != nil {
			fmt.Fprintf(w, "ParseForm() err: %v", err)
			return
		}
		fmt.Printf("GOT POST: %+v\n", r)
		n := r.FormValue("firstn")
		m := r.FormValue("lastn")
		o := r.FormValue("operation")
		fmt.Fprintf(w, "Got %v, %v, %v", n, m, o)
		// address := r.FormValue("address")
		// fmt.Fprintf(w, "Name = %s\n", name)
		// fmt.Fprintf(w, "Address = %s\n", address)
	default:
		fmt.Fprintf(w, "Sorry, only GET and POST methods are supported.")
	}
}

func main_server() {
	http.HandleFunc("/calculator", calculator)
	http.ListenAndServe(":8090", nil)
}
