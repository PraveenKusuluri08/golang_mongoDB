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

func createCourse(course model.NewCourse, userId string, authorName string) string {
	count, message := isUserExists(userId)
	if count != 0 && message != "" {
		return "User Not Exists!!"
	}
	course.UserId = userId
	course.IsCourseOwner = true
	course.CourseExists = true
	course.Author = authorName
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
	var user model.User
	params := mux.Vars(r)
	_ = json.NewDecoder(r.Body).Decode(&course)
	message := createCourse(course, params["id"], user.UserName)
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

func checkCourseTitleExists(title string) (int, string) {
	var course bson.M
	filter := bson.M{"title": title}
	var singleCourse []primitive.M
	err := collection.FindOne(context.Background(), filter).Decode(&course)
	singleCourse = append(singleCourse, course)
	if err != nil {
		log.Fatal(err)
	}
	if len(singleCourse[len(singleCourse)-1]) != 0 {
		return 0, ""
	}
	return 1, err.Error()
}

//creating the categories for the course

//category-1 => Web development
//category-2 => Mobile application development
//category-3 => Cloud technologies
//category-4 => Game development
//category-5 => Personality developement
//category-6 => Marketing
//category-7 => School Mathematics

//Course category are comes under the part of the course creation
//course creator must define which category the course is belongs to

//if the course are not exists with the requested category

//another Function needs to fire up,
//example
// Web Development requested course category
//there is no web development course
// slice the string and cut the development part from the Category
// provide all the course which are in development

func getAllCourseByCategories(category string) []primitive.M {
	cursor, err := Collection.Find(context.Background(), bson.M{"category": category})
	if err != nil {
		log.Fatal(err)
	}
	var allCourses []primitive.M

	for cursor.Next(context.TODO()) {
		var courses primitive.M
		if err := cursor.Decode(&courses); err != nil {
			log.Fatal(err)
		}
		allCourses = append(allCourses, courses)
	}
	defer cursor.Close(context.Background())
	return allCourses
}

func GetAllCoursesWithCategories(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Allow-Control-Allow-Methods", "GET")
	params := mux.Vars(r)
	courses := getAllCourseByCategories(params["category"])
	json.NewEncoder(w).Encode(courses)
	defer r.Body.Close()
}
