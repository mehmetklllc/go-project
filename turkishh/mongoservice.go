package turkishh

import (
	"context"
	"fmt"
	mkilic "github.com/mehmetklllc/go-project/go-custom-struct"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

)


func ConnectionAndGetCollection(url ,dbName ,collectionName string)  *mongo.Collection {

	clientOptions := options.Client().ApplyURI(url)
	client, err := mongo.Connect(context.TODO(), clientOptions)
	FailOnError(err,"not connetion to db!")
	collection := client.Database(dbName).Collection(collectionName)
	fmt.Println("Connected to MongoDB! "+collection.Name())
	//// Check the connection
	//err = client.Ping(context.TODO(), nil)
	//FailOnError(err,"not connetion to db!")

	return collection

}

func Save(jsonObject string,collection *mongo.Collection)  {
	object := mkilic.JsonToObject([]byte(jsonObject))
	go collection.InsertOne(context.TODO(), object)
}
