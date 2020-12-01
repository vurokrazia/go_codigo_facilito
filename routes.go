package main
import (
	"net/http"
	"fmt"
	)
func root_path(){
	http.HandleFunc("/", func (w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w,"Hello World")
	})	
}
func find_type_routes(){
	http.HandleFunc("/dos", func (w http.ResponseWriter, r *http.Request) {
		fmt.Println( r.Method)
		fmt.Fprintf(w, "Hello World " + r.Method )
	})
}