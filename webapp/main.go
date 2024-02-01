package main  

import (
	"fmt" 
	"net/http"
) 

// func main() {
// 	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request){
// 		fmt.Fprintf(w, "Welcome to my website!")
// 	})

// 	fs := http.FileServer(http.Dir("static/"))
// 	http.Handle("/static/", http.StripPrefix("/static/", fs))

// 	http.ListenAndServe(":8080", nil)
// }



func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request){
		fmt.Fprintf(w, "Welcome to my website!")
	})

	fs := http.FileServer(http.Dir("static/"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.ListenAndServe(":80", nil)
}




// func main() {
// 	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
// 		fmt.Fprintf(w, "Hello, you have requested: %s\n", r.URL.Path)

// 	})
// 	http.ListenAndServe(":80", nil)
	
// }


// Browse http.localhost:80


