package models

import (
	"context"
	"log"
	"time"

	"fmt"

	"github.com/webhook-repo/util"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func init() {
	MongoDbConnection.ConnectMongodb()
}

type mongodbConnectionType struct {
	MongoClient *mongo.Client
	isconnected bool
}

type Doc struct {
	RequestId  string    `json:"request_id"`
	Author     string    `json:"author"`
	Action     string    `json:"action"`
	FromBranch string    `json:"from_branch"`
	ToBranch   string    `json:"to_branch"`
	Time       time.Time `json:"time"`
}

var MongoDbConnection = &mongodbConnectionType{}

func (db *mongodbConnectionType) ConnectMongodb() error {

	if !db.isconnected {
		mongodb := util.GetConfig()
		host := mongodb["url"].(string)
		fmt.Println(host)

		client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(host))
		if err != nil {
			return err

		}

		db.MongoClient = client
		db.isconnected = true
		log.Println("DB Connected!")
	} else {
		// Do Nothing
	}
	return nil
}

func (db *mongodbConnectionType) CloseConnection() {
	log.Fatal("cleaning up db connection")
	db.MongoClient.Disconnect(context.TODO())
	db.isconnected = false
}

func (db *mongodbConnectionType) InserOneDoc(doc interface{}) error {
	if !db.isconnected {
		db.ConnectMongodb()
	}
	mongodb := util.GetConfig()
	database := mongodb["database"].(string)
	collection := mongodb["collecton"].(string)
	client := db.MongoClient
	coll := client.Database(database).Collection(collection)
	log.Println("Doc is being Added in mongoDB .....")
	_, err := coll.InsertOne(context.TODO(), doc)
	if err != nil {
		log.Fatal("Doc didn't add ")
		return err
	}
	log.Println("successfully added !")
	return err

}
func (db *mongodbConnectionType) InserManyDoc(docs []interface{}) error {
	if !db.isconnected {
		db.ConnectMongodb()
	}
	mongodb := util.GetConfig()
	database := mongodb["database"].(string)
	collection := mongodb["collecton"].(string)
	client := db.MongoClient
	coll := client.Database(database).Collection(collection)
	log.Println("Data is being Added in mongoDB .....")
	_, err := coll.InsertMany(context.TODO(), docs)
	if err != nil {
		//log.Fatal(err)
		return err

	}

	log.Println("successfully added !")
	return err

}

func (db *mongodbConnectionType) FindAllDoc() ([]Doc, error) {
	if !db.isconnected {
		db.ConnectMongodb()
	}
	mongodb := util.GetConfig()
	database := mongodb["database"].(string)
	collection := mongodb["collecton"].(string)
	client := db.MongoClient
	coll := client.Database(database).Collection(collection)

	cur, err := coll.Find(context.TODO(), bson.D{{}})
	if err != nil {
		log.Fatal("Doc didn't add ")
		return nil, err
	}
	var results []Doc
	for cur.Next(context.TODO()) {
		//Create a value into which the single document can be decoded
		var elem Doc
		err := cur.Decode(&elem)
		if err != nil {
			log.Fatal(err.Error())
		}

		results = append(results, elem)

	}
	//log.Println("successfully added !")
	return results, nil

}
