package main

import (
	"fmt"
	"net/http"
	"strconv"
)

func html(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "WebApp.html")
}
func factornum(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		if err := r.ParseMultipartForm(10000); err != nil {
			fmt.Fprintf(w, "ParseForm() err: %v", err)
		}
		a, _ := strconv.ParseUint(r.FormValue("numToFactor"), 10, 32)
		n := uint(a)
		fcts, ans := mainFactor(n)
		fmt.Fprintf(w, "%v\n %v", fcts, ans)
		fmt.Println(a)
		// address := r.FormValue("address")
		// fmt.Fprintf(w, "Name = %s\n", name)
		// fmt.Fprintf(w, "Address = %s\n", address)
	default:
		fmt.Fprintf(w, "Sorry, only POST method is supported.")
	}
}
func calculator(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		if err := r.ParseMultipartForm(10000); err != nil {
			fmt.Fprintf(w, "ParseForm() err: %v", err)
		}
		n, _ := strconv.ParseUint(r.FormValue("firstn"), 10, 32)
		m, _ := strconv.ParseUint(r.FormValue("lastn"), 10, 32)
		o := r.FormValue("op")
		ans, _ := computePrep(uint(n), uint(m), o)
		fmt.Fprintf(w, ans)
		// address := r.FormValue("address")
		// fmt.Fprintf(w, "Name = %s\n", name)
		// fmt.Fprintf(w, "Address = %s\n", address)
	default:
		fmt.Fprintf(w, "Sorry, only POST method is supported.")
	}
}

func mainServer() {
	http.HandleFunc("/", html)
	http.HandleFunc("/factor", factornum)
	http.HandleFunc("/calculator", calculator)
	http.ListenAndServe(":8090", nil)
}
