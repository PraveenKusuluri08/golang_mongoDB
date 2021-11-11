package controllers

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"github.com/PraveenKusuluri08/model"
	"github.com/PraveenKusuluri08/util"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const collection_name = "courses_list"

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
		collection = client.Database(dbName).Collection(collection_name)
	}
}

func createCourse(course model.NewCourse, userID string) interface{} {
	id, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		log.Fatal(err)
	}
	var user bson.M
	var singleUser []primitive.M
	filter := bson.M{"_id": id}
	err1 := collection.FindOne(context.TODO(), filter).Decode(&user)
	if err1 != nil {
		log.Fatal(err1)
	}
	singleUser = append(singleUser, user)
	//TODO:By now we get the single user document
	//TODO:Verify that document whether he is admin or not
	//TODO:if admin create new Course if not that case just set the courselist as null

	//just for sample we donot need to send the user document
	//insted we need to send the acknowledgment document created successfully or not
	return singleUser
}
func CreateCourse(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Allow-Control-Allow-Methods", "POST")
	params := mux.Vars(r)
	defer r.Body.Close()

	var newCourse model.NewCourse
	_ = json.NewDecoder(r.Body).Decode(&newCourse)
	data := createCourse(newCourse, params["id"])
	json.NewEncoder(w).Encode(data)

}
