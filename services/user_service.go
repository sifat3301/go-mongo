package services

import (
	"context"
	"go-mongo/configs"
	"go-mongo/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
)

var userCollection *mongo.Collection = configs.GetCollection(configs.DB, "users")

func CreateUserService(user *models.User, ctx context.Context) (models.User, error) {
	newUser := models.User{
		ID:       primitive.NewObjectID(),
		Name:     user.Name,
		Location: user.Location,
		Title:    user.Title,
	}
	userCollection.InsertOne(ctx, newUser)
	return newUser, nil
}

func GetUserService(userId *string, ctx *context.Context) (models.User, error) {
	var user models.User
	objId, _ := primitive.ObjectIDFromHex(*userId)
	err := userCollection.FindOne(*ctx, bson.M{"id": objId}).Decode(&user)
	if err != nil {
		return user, nil
	}
	return user, nil
}

func GetAllUserService(ctx *context.Context) ([]models.User, error) {
	cursor, err := userCollection.Find(*ctx, bson.M{})
	if err != nil {
		log.Fatal(err)
	}
	var userList []models.User
	//var users []bson.M
	//This is good for small dataset
	//if err = cursor.All(*ctx, &users); err != nil {
	//	log.Fatal(err)
	//}
	defer cursor.Close(*ctx)

	//In big datagset iterate over cursor good practice
	for cursor.Next(*ctx) {
		var user models.User
		if err = cursor.Decode(&user); err != nil {
			log.Fatal(err)
		}
		userList = append(userList, user)
	}
	return userList, nil
}

func UpdateUserService(userId *string, user *models.User, ctx *context.Context) (models.User, error) {

	objId, _ := primitive.ObjectIDFromHex(*userId)
	update := bson.M{"name": user.Name, "location": user.Location, "title": user.Title}
	result, _ := userCollection.UpdateOne(*ctx, bson.M{"id": objId}, bson.M{"$set": update})

	var updateUser models.User
	if result.MatchedCount == 1 {
		err := userCollection.FindOne(*ctx, bson.M{"id": objId}).Decode(&updateUser)
		if err != nil {
			return updateUser, err
		}
	}
	return updateUser, nil

}

func DeleteUserService(userId *string, ctx *context.Context) (bool, error) {

	objId, _ := primitive.ObjectIDFromHex(*userId)

	result, err := userCollection.DeleteOne(*ctx, bson.M{"id": objId})
	if err != nil {
		return false, err
	}
	if result.DeletedCount < 1 {
		return false, err
	}
	return true, nil

}

