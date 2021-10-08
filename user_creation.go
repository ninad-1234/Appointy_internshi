package main
import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
) 
type userData struct {
	ID 	  string   `json: id`
	Name      string   `json: name`
	Email     string   `json: email`
	Password  string   `json: password`
}

func weatherHandler(w http.ResponseWriter, r *http.Request) {
	
	jsn, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatal("Error reading the body", err)
	}
	err = json.Unmarshal(jsn, &location)
	if err != nil {
		log.Fatal("Decoding error: ", err)
	}
	log.Printf("Received: %v\n", location)
	



}
	

func server() {
	http.HandleFunc("/", weatherHandler)
	http.ListenAndServe(":8088", nil)
}

func main() {
	go server()

}
