package controllers

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"reflect"
	"time"

	"github.com/PraveenKusuluri08/model"
	"github.com/PraveenKusuluri08/util"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
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
	fmt.Println(reflect.TypeOf(r.Body))
	json.NewEncoder(w).Encode("User created Successfully")
}

func getAllUsers() []primitive.M {
	cursor, err := collection.Find(context.Background(), bson.D{{}}, nil)
	if err != nil {
		log.Fatal(err)
	}
	var users []primitive.M
	for cursor.Next(context.Background()) {
		var user primitive.M
		err := cursor.Decode(&user)
		if err != nil {
			log.Fatal(err)
		}
		users = append(users, user)
	}
	defer cursor.Close(context.Background())
	return users
}

func GetAllUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Allow-Control-Allow-Methods", "GET")
	fmt.Println(reflect.TypeOf(r.Body))
	usersData := getAllUsers()
	json.NewEncoder(w).Encode(usersData)
}

func UpdateSingleUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Allow-Control-Allow-Methods", "PUT")
	params := mux.Vars(r)
	id, err := primitive.ObjectIDFromHex(params["id"])
	if err != nil {
		log.Fatal(err)
	}
	filter := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"isExists": false, "isLoggedin": false}}
	data, err1 := collection.UpdateOne(context.Background(), filter, update)
	if err1 != nil {
		log.Fatal(err1)
	}
	defer r.Body.Close()
	json.NewEncoder(w).Encode(data)
}

func deleteSingleUser(userID string) {
	id, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		panic(err)
	}
	filter := bson.M{"_id": id}
	count, err := collection.DeleteOne(context.TODO(), filter)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(count)
}

func DeleteSingleUser(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	w.Header().Set("Content-type", "application/json")
	w.Header().Set("Allow-Control-Allow-Methods", "PUT")
	params := mux.Vars(r)
	deleteSingleUser(params["id"])
	json.NewEncoder(w).Encode("User deleted successfully")
}
