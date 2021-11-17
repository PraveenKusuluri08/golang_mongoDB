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
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const collection_name = "courses_list"
const DBName = "COURSES"

var Collection *mongo.Collection

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
		Collection = client.Database(DBName).Collection(collection_name)
	}
}

func createCourse(course model.NewCourse, userId string) string {
	count, message := isUserExists(userId)
	if count != 0 && message != "" {
		return "User Not Exists!!"
	}
	course.UserId = userId
	course.CourseExists = true
	course.CreatedAt = time.Now().String()
	_, err := Collection.InsertOne(context.Background(), course)
	if err != nil {
		return "Failed to create Course"
	}
	return "Course Created Successfully"
}

func CreateCourse(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Allow-Control-Allow-Methods", "POST")
	defer r.Body.Close()
	var course model.NewCourse
	params := mux.Vars(r)
	_ = json.NewDecoder(r.Body).Decode(&course)
	message := createCourse(course, params["id"])
	json.NewEncoder(w).Encode(message)
}

//helper function for checking user is Exists or in the db
func isUserExists(userId string) (int, string) {
	id, err := primitive.ObjectIDFromHex(userId)
	if err != nil {
		fmt.Println("Not valid id")
	}
	var user bson.M
	filter := bson.M{"_id": id, "isexists": true}
	var singleUser []primitive.M
	//check user exists in the Users Collection
	err1 := collection.FindOne(context.Background(), filter).Decode(&user)
	singleUser = append(singleUser, user)
	if len(singleUser[len(singleUser)-1]) != 0 {
		return 0, ""
	}
	return 1, err1.Error()
}
