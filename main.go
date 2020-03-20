package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

type Server struct {

}

func (s *Server) simpleHttp( w http.ResponseWriter, r *http.Request) {
	fmt.Printf("got simple\n")
  fmt.Fprintf(w,"testing testing 1,2,3")
}

func (s *Server) routes() {
  http.HandleFunc("/testfunction", s.simpleHttp)
}

func (s *Server) run() {
	port, exists := os.LookupEnv("FUNCTIONS_HTTPWORKER_PORT")
	if !exists {
		port = "8080"
	}
	fmt.Printf("port used is %s\n", port)
	log.Fatal(http.ListenAndServe(":"+port,nil))
}


func main() {

	fmt.Printf("hello world :)\n")
	s := Server{}
	s.routes()
	s.run()
}
