package controllers

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/PraveenKusuluri08/model"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

//👇 this function need to fire up only after the payment completion

//Update user when the user buys any course
//The course buyer model needs to fire up
func updateUserWhenCourseBought(userId string, courseBought model.CoursesBuyer) string {
	count, message := isUserExists(userId)

	if count != 0 && message != "" {
		return "No user user exists"
	}

	courseBought.RecentlyBoughtCourseDate = time.Now().String()
	id, err := primitive.ObjectIDFromHex(userId)
	if err != nil {
		log.Fatal(err)
	}

	filter := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"coursesbought": courseBought}}
	collection.FindOneAndUpdate(context.Background(), filter, update)
	return "Course Added Successfully"

}

func UpdateUserWhenCourseBought(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Allow-Control-Allow-Methods", "PUT")
	params := mux.Vars(r)
	var courseBought model.CoursesBuyer
	defer r.Body.Close()
	_ = json.NewDecoder(r.Body).Decode(&courseBought)
	fmt.Println(courseBought.CourseArticels)
	userUpdated := updateUserWhenCourseBought(params["id"], courseBought)
	json.NewEncoder(w).Encode(userUpdated)
}

//function to add to whishlist

//delete functio to delete the course from the cart
