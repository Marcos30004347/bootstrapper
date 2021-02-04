package database

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson/primitive"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"labix.org/v2/mgo/bson"
)

type User struct {
	Id       string `json:"id"     bson:"_id"`
	Username string `json:"username"  bson:"username"`
}

type Database struct {
	client *mongo.Client
}

var db *Database = nil

func InitDatabase(url string) {
	db = &Database{}

	clientOptions := options.Client().ApplyURI(url)

	// Connect to MongoDB
	client, err := mongo.Connect(context.Background(), clientOptions)

	if err != nil {
		log.Fatal(err)
		panic(1)
	}

	// Check the connection
	err = client.Ping(context.Background(), nil)

	if err != nil {
		log.Fatal(err)
		panic(1)
	}

	fmt.Println("Connected to MongoDB!")

	db.client = client
}

func ShutDownDatabase() {

	db.client.Disconnect(context.Background())
}

func CreateUser(user *User) (string, error) {

	result, err := db.client.Database("users").Collection("users").InsertOne(context.Background(), user)

	oid, ok := result.InsertedID.(primitive.ObjectID)
	if ok {
		return oid.Hex(), err
	}
	return "", DatabaseError{1, "error"}
}

func GetAll() ([]*User, error) {
	// passing bson.D{{}} matches all documents in the collection
	filter := bson.D{{}}
	return FilterUsers(filter)
}

func FilterUsers(filter interface{}) ([]*User, error) {
	var tasks []*User
	cur, err := db.client.Database("users").Collection("users").Find(context.Background(), filter)

	if err != nil {
		return tasks, err
	}

	for cur.Next(context.Background()) {

		var tmp User
		err := cur.Decode(&tmp)

		if err != nil {
			return tasks, err
		}

		tasks = append(tasks, &tmp)
	}

	if err := cur.Err(); err != nil {
		return tasks, err
	}

	cur.Close(context.Background())

	if len(tasks) == 0 {
		return tasks, mongo.ErrNoDocuments
	}

	return tasks, nil
}

func Update() {
	filter := bson.M{"title": "Teste"}
	update := bson.M{"$inc": bson.M{"age": 1}}

	updateResult, err := db.client.Database("users").Collection("users").UpdateOne(context.Background(), filter, update)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Matched %v documents and updated %v documents.\n", updateResult.MatchedCount, updateResult.ModifiedCount)

}

func FindUserById(id string) User {
	filter := bson.M{"_id": bson.ObjectIdHex(id)}

	findResult := db.client.Database("users").Collection("users").FindOne(context.Background(), filter)

	var tmp User
	findResult.Decode(&tmp)

	return tmp
}
