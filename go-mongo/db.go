package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)


func db() *mongo.Client {

	clientOptions := options.Client().ApplyURI("mongodb://root:example@localhost:27017")// Connect to //MongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}
// Check the connection
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to MongoDB!")
	return client
}

type expense struct{
	Name string `json:"name"`
	City string `json:"city"`
	Age int `json:"age"`
	}
var userCollection = db().Database("goTest").Collection("users")


// POST 
//Create
func createExpense(w http.ResponseWriter, r *http.Request) {w.Header().Set("Content-Type", "application/json") 
	var exp expense
	
	err := json.NewDecoder(r.Body).Decode(&exp) // storing in exp   //variable of type expense
	if err != nil {
		fmt.Print(err)
	}
	insertResult, err := userCollection.InsertOne(context.TODO(),exp)
	if err != nil {
		log.Fatal(err)
	}

	json.NewEncoder(w).Encode(insertResult.InsertedID) 
}
// GET
//Read
func getExpense(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var body expense
	e := json.NewDecoder(r.Body).Decode(&body)
	if e != nil {
	fmt.Print(e)
	}
	var result primitive.M //  an unordered representation of a BSON //document which is a Map
	err := userCollection.FindOne(context.TODO(), bson.D{{"name", body.Name}}).Decode(&result)
	if err != nil {
	fmt.Println(err)
	}
	json.NewEncoder(w).Encode(result) // returns a Map containing //mongodb document
}
// PUT
// Update
func updateExpense(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	type updateBody struct {
	Name string `json:"name"` //value that has to be matched
	City string `json:"city"` // value that has to be modified
	}
	var body updateBody
	e := json.NewDecoder(r.Body).Decode(&body)
	if e != nil {
	fmt.Print(e)
	}
	filter := bson.D{{"name", body.Name}} // converting value to BSON 
	after := options.After         // for returning updated document    
	returnOpt := options.FindOneAndUpdateOptions{
	ReturnDocument: &after,
	}
	update := bson.D{{"$set", bson.D{{"city", body.City}}}}
	updateResult := userCollection.FindOneAndUpdate(context.TODO(), filter, update, &returnOpt)
	var result primitive.M
	_ = updateResult.Decode(&result)
	json.NewEncoder(w).Encode(result)
	}
// Delete
		func deleteExpense(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/json")
			params := mux.Vars(r)["id"] //get Parameter value as string
			_id, err := primitive.ObjectIDFromHex(params) // convert params to //mongodb Hex ID
			if err != nil {
			fmt.Printf(err.Error())
			}
			opts := options.Delete().SetCollation(&options.Collation{}) // to specify language-specific rules for string comparison, such as //rules for lettercase
			res, err := userCollection.DeleteOne(context.TODO(), bson.D{{"_id", _id}}, opts)
			if err != nil {
			log.Fatal(err)
			}
			json.NewEncoder(w).Encode(res.DeletedCount) // return number of //documents deleted
			}