
package main

import (
	"context"
	"fmt"
	"os"
	"reflect"
	"time"
	"io/ioutil"
	"log"
	"net/http"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoField struct {
	ID  		string 	`json: "id"`
	Caption  	string  `json: "caption"`
	IMG_URL 	string  `json: "imgurl"`
	Timestamp 	string  `json: "psw"`
}

func server(userDB MongoField){
	const url ="http://localhost:8000/post"
	var  byte  = []byte(userDB.ID)
	req, err := http.NewRequest("POST", url, byte)
	print(req,err)
}

func main() {
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	fmt.Println("ClientOptopm TYPE:", reflect.TypeOf(clientOptions), "\n")

	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		fmt.Println("Mongo.connect() ERROR: ", err)
		os.Exit(1)
	}
	ctx, _ := context.WithTimeout(context.Background(), 15*time.Second)

	col := client.Database("Instagram_Database").Collection("Post COllection")
	fmt.Println("Collection Type: ", reflect.TypeOf(col), "\n")

	userDB := MongoField{
		ID:  		"ITpost_1001",
		Caption:  	"Fact about microprocessor",
		IMG_URL: 	"microprocessor.jpeg",
		Timestamp: 	"nt2001$",
	}

	fmt.Println("oneDoc Type: ", reflect.TypeOf(userDB), "\n")

	result, insertErr := col.InsertOne(ctx, userDB)
	if insertErr != nil {
		fmt.Println("InsertONE Error:", insertErr)
		os.Exit(1)
	} else {
		server(MongoField{})
		fmt.Println("InsertOne() result type: ", reflect.TypeOf(result))

	}
}
