package main

import (
	"fmt"
	"net/http"
	"strconv"
)

func calculator(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		http.ServeFile(w, r, "calculator1.html")
	case "POST":
		fmt.Printf("GOT POST: %+v\n", r)
		// Call ParseForm() to parse the raw query and update r.PostForm and r.Form.
		if err := r.ParseMultipartForm(10000); err != nil {
			fmt.Fprintf(w, "ParseForm() err: %v", err)
			return
		}
		n, _ := strconv.ParseUint(r.FormValue("lastn"), 10, 32)
		m, _ := strconv.ParseUint(r.FormValue("firstn"), 10, 32)
		o := r.FormValue("operation")
		ans, _ := computePrep(uint(n), uint(m), o)
		fmt.Fprintf(w, ans)
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
