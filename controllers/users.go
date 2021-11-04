package controllers

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/PraveenKusuluri08/model"
	"github.com/PraveenKusuluri08/util"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const dbName = "USERS"
const collectionName = "user_list"

var collection *mongo.Collection

//Function to connect to the db
func init() {
	mongoUri, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal(err)
	} else {
		clientOptions := options.Client().ApplyURI(mongoUri.MongoURI)

		client, err := mongo.Connect(context.TODO(), clientOptions)
		if err != nil {
			log.Fatal(err)
		}
		collection = client.Database(dbName).Collection(collectionName)
	}
}

func createUser(user model.User) interface{} {
	passwordHashed, _ := util.PasswordHasher(user.Password)
	user.Password = passwordHashed
	user.CreatedAt = time.Now().String()

	createUser, err := collection.InsertOne(context.Background(), user)

	if err != nil {
		log.Fatal(err)
	}
	return createUser.InsertedID
}

func CreateUserAccout(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Allow-Control-Allow-Methods", "POST")

	defer r.Body.Close()
	var user model.User
	_ = json.NewDecoder(r.Body).Decode(&user)

	userId := createUser(user)
	fmt.Println(userId)

	json.NewEncoder(w).Encode("User created Successfully")
}
