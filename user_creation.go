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

func userDataHandler(w http.ResponseWriter, r *http.Request) {
	
	jsn, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatal("Error reading the body", err)
	}
	log.Printf("Received: %v\n", location)
	
	User_data := userData{
		Id: 	   "ninad0104",
		Name:      "Ninad Topale",
		Email:     "topaleninad01@gmail.com",
		Password:  "nt1004$",
		},
	}
// Above userData is input from the UI

}
	

func server() {
	http.HandleFunc("/", userDataHandler)
	http.ListenAndServe(":8088", nil)
	//Redirect to Home page 
}

func main() {
	go server()

}
