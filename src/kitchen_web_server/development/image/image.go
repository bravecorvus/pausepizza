package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

type request struct {
	Filename string `json:"filename"`
}

type response struct {
	Status  string
	Message string
}

var results []string

func pwd() string {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatal(err)
	}
	return dir
}

func main() {
	http.HandleFunc("/upload", uploadHandler)
	http.HandleFunc("/delete", deleteHandler)
	log.Println("Listening...")
	log.Fatal(http.ListenAndServe(":3000", nil))
}

func uploadHandler(w http.ResponseWriter, r *http.Request) {

	file, header, err := r.FormFile("file")
	// fmt.Println("file", file)
	// fmt.Println("header", header)

	if err != nil {
		json.NewEncoder(w).Encode(response{Status: "failure", Message: "Couldn't decode http request"})
		return
	}
	defer file.Close()
	// fmt.Println(reflect.TypeOf(file))
	out, err1 := os.Create(pwd() + "/files/" + header.Filename)
	if err1 != nil {
		json.NewEncoder(w).Encode(response{Status: "failure", Message: "Couldn't create file" + header.Filename})
		return
	}
	defer out.Close()
	_, err2 := io.Copy(out, file)
	if err2 != nil {
		json.NewEncoder(w).Encode(response{Status: "failure", Message: err2.Error()})
		return
	}
	fmt.Fprintf(w, "File uploaded successfully :")
	json.NewEncoder(w).Encode(response{Status: "success", Message: "File uploaded successfully"})
}

func deleteHandler(w http.ResponseWriter, r *http.Request) {

	decoder := json.NewDecoder(r.Body)
	var query request
	err := decoder.Decode(&query)
	if err != nil {
		json.NewEncoder(w).Encode(response{Status: "failure", Message: "HTTP header could not be decoded"})
		return
	}

	// fmt.Println(query.Filename)
	err1 := os.Remove(pwd() + "/files/" + query.Filename)

	if err1 != nil {
		json.NewEncoder(w).Encode(response{Status: "failure", Message: "Failed to delete file"})
		return
	}
	json.NewEncoder(w).Encode(response{Status: "success", Message: "Delete " + query.Filename + " successful"})
}
