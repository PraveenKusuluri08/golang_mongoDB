package helpers

import (
	"context"
	"log"

	"github.com/PraveenKusuluri08/model"
	"go.mongodb.org/mongo-driver/mongo"
)

var collection *mongo.Collection

//TODO:when user is just creates account
//TODO:all the couses details are empty
//TODO:if user bought courses then update request need to be fired up
//TODO:When user create his/her accout password need to be hashed
func CreateUser(user model.User) interface{} {
	// passwordHashed, _ := util.PasswordHasher(user.Password)
	// user.Password = passwordHashed
	createUser, err := collection.InsertOne(context.Background(), user)

	if err != nil {
		log.Fatal(err)
	}
	return createUser.InsertedID

}
