package main

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var client *mongo.Client

//ConfigDB initialize MongoDB connection
func ConfigDB(hostname string, username string, password string) error {
	// Set client options
	clientOptions := options.Client().ApplyURI("mongodb://" + username + ":" + password + "@" + hostname + ":27017")
	// Connect to MongoDB
	var err error
	client, err = mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		return err
	}
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		return err
	}
	return nil
}

func insertRecord(record Record) error {
	fmt.Println("insertRecord")
	// Check the connection
	err := client.Ping(context.TODO(), nil)
	if err != nil {
		return err
	}

	fmt.Println("Connected to MongoDB!")

	// Get a handle for your collection
	collection := client.Database("testdb").Collection("people")
	log.Println("record", record)

	var result Record
	filter := bson.D{{"name", record.Name}}

	err = collection.FindOne(context.TODO(), filter).Decode(&result)
	log.Println("result", result)

	if result.Name != "" {
		log.Println("update")
		err = updateRecord(record)
		return err
	}

	// Insert a single document
	insertResult, err := collection.InsertOne(context.TODO(), record)
	if err != nil {
		return err
	}
	fmt.Println("Inserted a single document: ", insertResult.InsertedID)

	return nil
}

func updateRecord(record Record) error {

	// Check the connection
	err := client.Ping(context.TODO(), nil)
	if err != nil {
		return err
	}

	fmt.Println("Connected to MongoDB!")

	// Get a handle for your collection
	collection := client.Database("testdb").Collection("people")

	// Update a document
	filter := bson.D{{"name", record.Name}}

	update := bson.D{
		{"$set", bson.D{
			{"DateOfBirth", record.DateOfBirth},
		}},
	}

	updateResult, err := collection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		return err
	}
	fmt.Printf("Matched %v documents and updated %v documents.\n", updateResult.MatchedCount, updateResult.ModifiedCount)

	return nil
}

func findRecord(name string) (Record, error) {
	var result Record
	// Check the connection
	err := client.Ping(context.TODO(), nil)
	if err != nil {
		return result, err
	}

	fmt.Println("Connected to MongoDB!")

	// Get a handle for your collection
	collection := client.Database("testdb").Collection("people")

	// Find a single document

	filter := bson.D{{"name", name}}

	err = collection.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		return result, err
	}

	fmt.Printf("Found a single document: %+v\n", result)

	return result, nil
}
