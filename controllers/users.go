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

func CreateUser(user model.User) interface{} {
	passwordHashed, _ := util.PasswordHasher(user.Password)
	user.Password = passwordHashed
	user.CreatedAt = time.Now().String()
	user.IsExists = true
	//if role is admin no course bought property is allowed
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

	userId := CreateUser(user)
	fmt.Println(userId)
	fmt.Println(reflect.TypeOf(r.Body))
	json.NewEncoder(w).Encode("User created Successfully")
}

//get all users control functions

func getAllUsers() []primitive.M {
	cursor, err := collection.Find(context.Background(), bson.D{{}}, nil)
	if err != nil {
		log.Fatal(err)
	}

	var users []primitive.M
	for cursor.Next(context.Background()) {
		var user primitive.M
		if err := cursor.Decode(&user); err != nil {
			log.Fatal(err)
		}

		users = append(users, user)
	}
	defer cursor.Close(context.Background())
	return users
}

//get all users function
func GetAllUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Allow-Control-Allow-Methods", "GET")
	fmt.Println(reflect.TypeOf(r.Body))
	usersData := getAllUsers()
	json.NewEncoder(w).Encode(usersData)
}

//Update User control Function
func UpdateSingleUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Allow-Control-Allow-Methods", "PUT")
	params := mux.Vars(r)
	id, err := primitive.ObjectIDFromHex(params["id"])
	if err != nil {
		log.Fatal(err)
	}
	filter := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"isLoggedin": false}}
	data, err1 := collection.UpdateOne(context.Background(), filter, update)
	if err1 != nil {
		log.Fatal(err1)
	}
	defer r.Body.Close()
	json.NewEncoder(w).Encode(data)
}

//Perminently delete user control function

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

//Perminently delete user function
func DeleteSingleUser(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	w.Header().Set("Content-type", "application/json")
	params := mux.Vars(r)
	deleteSingleUser(params["id"])
	json.NewEncoder(w).Encode("User deleted successfully")
}

//TODO:Delete If User is not exists, Means isExists=false(wrost case->User request to delete Accout perminently)
//TODO:Delete If User is Request to delete ->We need to set isExists=false (we need to take care in Frontend To display user when is isExists=false)
//TODO:We need to have two routes for this

//diable user control function
func disableUser(userId string) interface{} {
	id, err := primitive.ObjectIDFromHex(userId)

	if err != nil {
		log.Fatal(err)
	}
	filter := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"isexists": false}}
	data, err1 := collection.UpdateOne(context.Background(), filter, update)
	if err1 != nil {
		log.Fatal(err1)
	}
	fmt.Println("Data disabled successfully")
	return data
}

//Disable user fuction

func DisableUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	fmt.Println("cxfc", params["id"])
	count := disableUser(params["id"])

	json.NewEncoder(w).Encode(count)
}

//enable user control function
func enableUser(userId string) interface{} {
	id, err := primitive.ObjectIDFromHex(userId)
	if err != nil {
		log.Fatal(err)
	}
	filter := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"isexists": false}}
	data, err1 := collection.UpdateOne(context.Background(), filter, update)
	if err1 != nil {
		log.Fatal(err1)
	}
	fmt.Println("Data disabled successfully")
	return data
}

//enable user

func EnableUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Allow-Control-Allow-Methods", "PUT")
	params := mux.Vars(r)
	count := enableUser(params["id"])
	json.NewEncoder(w).Encode(count)
}

//Get the single user from the database
func getSingleUser(userId string) interface{} {
	id, err := primitive.ObjectIDFromHex(userId)

	if err != nil {
		log.Panic(err)
	}
	var user bson.M
	filter := bson.M{"_id": id}
	var singleUser []primitive.M
	err1 := collection.FindOne(context.Background(), filter).Decode(&user)
	if err1 != nil {
		log.Fatal(err)
	}
	singleUser = append(singleUser, user)

	return singleUser
}

func GetSingleUserDocument(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	data := getSingleUser(params["id"])
	json.NewEncoder(w).Encode(data)
}
