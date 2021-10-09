
package main

import (
	"context" // manage multiple requests
	"fmt"     // Println() function
	"net/http"
	"os"
	"reflect" // get an object type
	"time"

	// import 'mongo-driver' package libraries
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoField struct {
	ID  		string 	`json: "id"`
	Caption  	string  `json: "caption"`
	IMG_URL 	string  `json: "imgurl"`
	Timestamp 	string  `json: "psw"`
}

func user_posts(id string, cnt int, post_arr []string){
	post_arr[cnt]= id
}
func server(post_arr[]string) {
	const url = "http://localhost:8000/post"
	var byte = []byte(post_arr)
	req, err := http.NewRequest("POST", url, byte)
	print(req,err)
}

func main() {
	// Declare host and port options to pass to the Connect() method
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	fmt.Println("clientOptions type:", reflect.TypeOf(clientOptions), "\n")

	// Connect to the MongoDB and return Client instance
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		fmt.Println("mongo.Connect() ERROR:", err)
		os.Exit(1)
	}

	// Declare Context type object for managing multiple API requests
	ctx, _ := context.WithTimeout(context.Background(), 15*time.Second)

	// Defining User id = ninad010#
	ip_id := "ninad010#"
	// Access a MongoDB collection through a database
	col := client.Database("Instagram_Database").Collection("Post COllection")
	fmt.Println("Collection type:", reflect.TypeOf(col), "\n")

	cursor, err := col.Find(context.TODO(), bson.D{})
	if err != nil {
        fmt.Println("Finding all documents ERROR:", err)
        defer cursor.Close(ctx)

    // If the API call was a success
    } else{
		var cnt int 
		cnt=0
		var post[20] string
		for cursor.Next(ctx) {
			var result bson.M
			err := cursor.Decode(&result)
			if err != nil {
                fmt.Println("cursor.Next() error:", err)
                os.Exit(1)
               
            // If there are no cursor.Decode errors
            }else {
				if(ip_id == result.ID){
					user_posts(result.ID,cnt,post)
					cnt++

				}
				
            }
		}
		server(post[:])
	}

}
