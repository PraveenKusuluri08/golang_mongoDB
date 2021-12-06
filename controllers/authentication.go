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
	"github.com/dgrijalva/jwt-go"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func signinUser(signin model.AuthSignIn) interface{} {
	isExists, message := IsEmailExists(signin.Email)
	var actualUser model.User
	filter := bson.M{"email": signin.Email}
	if isExists && message == "" {
		if err := collection.FindOne(context.TODO(), filter).Decode(&actualUser); err != nil {
			log.Fatal(err.Error())
		}
		mathches := util.CompareHashAndPassword(signin.Password, actualUser.Password)
		if mathches {
			var token model.Token
			tokenString := GenerateJwt(signin.Email)
			token.TokenString = tokenString
			token.Email = signin.Email
			return token
		}
	}
	return message
}

func SingIn(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Header().Set("Allow-Control-Allow-Methods", "POST")
	var signin model.AuthSignIn
	_ = json.NewDecoder(r.Body).Decode(&signin)
	data := signinUser(signin)
	json.NewEncoder(w).Encode(data)
	// fmt.Println(data)

}
func signup(signup model.User) interface{} {
	isExists, message := IsEmailExists(signup.Email)
	var token model.Token
	fmt.Println("isExists", isExists)
	if !isExists {
		userId := CreateUser(signup)
		fmt.Println(userId)
		if userId != nil {
			tokenString := GenerateJwt(signup.Email)
			token.TokenString = tokenString
			token.Email = signup.Email
			return token
		}
	}
	return message
}

func SignUp(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Allow-Control-Allow-Methods", "POST")
	var signUp model.User
	_ = json.NewDecoder(r.Body).Decode(&signUp)
	token := signup(signUp)
	json.NewEncoder(w).Encode(token)

}

//helper functions
func IsEmailExists(email string) (bool, string) {
	filter := bson.M{"email": email}
	var user bson.M
	var singleUser []primitive.M
	err := collection.FindOne(context.Background(), filter).Decode(&user)

	singleUser = append(singleUser, user)

	if err != nil {
		return false, "No user exists with the requested email address"
	}

	if len(singleUser[len(singleUser)-1]) != 0 {
		return true, ""
	}
	return false, "User Not exists"
}

func GenerateJwt(email string) string {
	secretKey, err := util.LoadConfig(".")

	if err != nil {
		log.Fatal(err)
	}
	var mySigninKey = []byte(secretKey.SecretKey)
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["authorized"] = true

	claims["email"] = email
	claims["empTime"] = time.Now().Add(time.Minute * 5).Unix()

	tokenString, err1 := token.SignedString(mySigninKey)

	if err1 != nil {
		panic(err1)
	}
	return tokenString
}

func IsAdmin(userId string, role int) (bool, string) {
	id, err := primitive.ObjectIDFromHex(userId)

	if err != nil {
		log.Fatal(err)
	}
	var user model.User
	filter := bson.M{"_id": id}
	ctx, _ := context.WithTimeout(context.Background(), 100*time.Second)
	err1 := collection.FindOne(ctx, filter).Decode(&user)

	if err1 != nil {
		return false, "No user exists with the requested id"
	}
	if role == user.Role {
		return true, "User is admin"
	}
	return false, "User is not admin"
}

//implement isAuthorised method for the privilage the end points

//implement forgot password method

//add mailing service to the application
