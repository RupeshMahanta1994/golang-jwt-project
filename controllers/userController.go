package controllers

import (
	"context"
	"net/http"
	"time"

	"github.com/RupeshMahanta1994/go-jwt-project/database"
	"github.com/RupeshMahanta1994/go-jwt-project/helpers"
	"github.com/RupeshMahanta1994/go-jwt-project/models"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var userCollection *mongo.Collection = database.OpenCollection(database.Client, "user")
var validate = validator.New()

// func HashPassword() {}

// func VerifyPassword() {}

func SignUp(c *gin.Context) {
	var _, cancel = context.WithTimeout(context.Background(), 100*time.Second)
	var user models.User

	//Bind JSON to user struct
	err := c.ShouldBindJSON(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error in binding json": err.Error()})
	}
	//validate user struct
	err = validate.Struct(user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error in validating the user": err.Error()})
	}

	//set User id and bind time
	user.ID = primitive.NewObjectID()
	user.Created_at = time.Now()
	user.Updated_at = time.Now()
	user.User_id = user.ID.Hex()

	//insert user into mongo db
	_, err = userCollection.InsertOne(context.Background(), user)
	defer cancel()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error in inserting the user": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, user)
}

// func Login() {}

// func GetUsers() {}

func Getuser() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		userId := ctx.Param("user_id")
		err := helpers.MatchUserTypeToUid(ctx, userId)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"err": err.Error()})
		}

		var _, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		var user models.User
		err = userCollection.FindOne(ctx, bson.M{"user_id": userId}).Decode(&user)
		defer cancel()
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(http.StatusOK, user)
	}
}
