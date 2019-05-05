package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os/exec"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/", Options).Methods("GET")
	router.HandleFunc("/close", CloseChrome).Methods("GET")
	fmt.Println("Starting the server...")
	http.ListenAndServe(":9090", router)
	fmt.Println("Server Started...")
}

func WarnMessage(w http.ResponseWriter, r *http.Request) {
	cmd := exec.Command("warn.vbs")
	err := cmd.Start()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Waiting for command to finish...")
	err = cmd.Wait()
	log.Printf("Command finished with error: %v", err)
	json.NewEncoder(w).Encode("Warn window displayed")
}

// Take the application name as request param
func CloseChrome(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Close request")
	cmd := exec.Command("taskkill", "/f", "/t", "/im", "chrome.exe")
	err := cmd.Start()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Waiting for command to finish...")
	err = cmd.Wait()
	log.Printf("Command finished with error: %v", err)
	json.NewEncoder(w).Encode("Closed")
}

func Options(w http.ResponseWriter, r *http.Request) {
	//	TODO: Move to html file
	var page = "<!DOCTYPE html>"
	page += "<html>"
	page += "<body>"
	page += "<h1> Hello </h1>"
	page += "<hr>"
	page += "Click <a href=\"close\"> here </a> to close Chrome in server."
	page += "</body>"
	page += "</html>"
	fmt.Fprintf(w, page)
}
