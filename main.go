package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
)

type Server struct {
}

type InvokeRequest struct {
	Data     map[string]interface{}
	Metadata map[string]interface{}
}

type InvokeResponse struct {
	Outputs     map[string]interface{}
	Logs        []string
	ReturnValue interface{}
}


func (s *Server) queueTriggerHandler(w http.ResponseWriter, r *http.Request) {
	var invokeReq InvokeRequest
	d := json.NewDecoder(r.Body)
	decodeErr := d.Decode(&invokeReq)
	if decodeErr != nil {
		http.Error(w, decodeErr.Error(), http.StatusBadRequest)
		return
	}
	fmt.Println("The JSON data is:invokeReq metadata......")
	fmt.Println(invokeReq.Metadata)
	fmt.Println("The JSON data is:invokeReq data......")
	fmt.Println(invokeReq.Data)

	returnValue := fmt.Sprintf("HelloWorld: %s", invokeReq.Data["myQueueItem"])
	invokeResponse := InvokeResponse{Logs: []string{"test log1", "test log2"}, ReturnValue: returnValue}

	js, err := json.Marshal(invokeResponse)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
	//w.WriteHeader(200)
}

func (s *Server) simpleHttp(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("got simple\n")
	fmt.Fprintf(w, "afken2 example...")
}

func (s *Server) simpleHttp2(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("got simple\n")
	fmt.Fprintf(w, "afken2 example 2...")
}

func (s *Server) routes() {
	http.HandleFunc("/testfunction", s.simpleHttp)
	http.HandleFunc("/testfunction2", s.simpleHttp2)
	http.HandleFunc("/queuetrigger", s.queueTriggerHandler)
}

func (s *Server) run() {
	port, exists := os.LookupEnv("FUNCTIONS_HTTPWORKER_PORT")
	if !exists {
		port = "8080"
	}
	fmt.Printf("port used is %s\n", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

func main() {

	fmt.Printf("hello world :)\n")
	s := Server{}
	s.routes()
	s.run()
}
