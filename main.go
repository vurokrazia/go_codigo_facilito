package main
import (
	"net/http"
	"log"
	// "fmt"
	)
func main()  {
	redirect := http.RedirectHandler("https://www.google.com.mx", http.StatusMovedPermanently)
	notFound := http.NotFoundHandler()

	mux := http.NewServeMux()

	mux.Handle("/redirect", redirect)
	mux.Handle("/not", notFound)


	server := &http.Server{
		Addr: ":3000",
		Handler: mux,
	}
	// root_path()
	// find_type_routes()
	log.Fatal(server.ListenAndServe())
}