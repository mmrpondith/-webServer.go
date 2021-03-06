package main

import (
	"fmt"
	"net/http"
)

/*type Handler interface {
	ServeHTTP(ResponseWriter, *Request)
}*/
func main() {
	//var name Date type
	//var x string
	//var handler func (ResponseWriter, *Request)
	http.HandleFunc("/", Home)
	http.HandleFunc("/About", About)
	http.HandleFunc("/Contect", Contect)
	http.ListenAndServe(":8888", nil)
}
func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, `This is my new document page `)
}
func About(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, `welcome to our About page `)
}
func Contect(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, `welcome to our Contect page `)
}
