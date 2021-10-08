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
	
	url := "http://localhost:8000/post"
	var username = []byte(`{"Id":"ninad0104"}`)
    	req, err := http.NewRequest("POST", url, bytes.NewBuffer(username))
	
	User_data := userData{
		Id: 	   "ninad0104",
		Name:      "Ninad Topale",
		Email:     "topaleninad01@gmail.com",
		Password:  "nt1004$",
		},
	}
	server(req)
// Above userData is input from the UI

}
	

func server(string req) {
	client := &http.Client{}
    	resp, err := client.Do(req)
    	if err != nil {
            panic(err)
        }
    defer resp.Body.Close()
	//Redirect to Home page 
}

func main() {
	go userDataHandler()

}
