package xk6_mongo

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/bson"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	k6modules "go.k6.io/k6/js/modules"
)

// Register the extension on module initialization, available to
// import from JS as "k6/x/mongo".
func init() {
	k6modules.Register("k6/x/mongo", new(Mongo))
}

// Mongo is the k6 extension for a Mongo client.
type Mongo struct{}

// Client is the Mongo client wrapper.
type Client struct {
	client   *mongo.Client
	database string
}

// NewClient represents the Client constructor (i.e. `new mongo.Client()`) and
// returns a new Mongo client object.
// connURI -> mongodb://username:password@address:port/db?connect=direct
func (*Mongo) NewClient(connURI string, database string) interface{} {
	clientOptions := options.Client().ApplyURI(connURI)
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		return err
	}
	return &Client{client: client, database: database}
}

func (c *Client) DropDatabase(database string) error {
	err := c.client.Database(database).Drop(context.TODO())
	if err != nil {
		return err
	}
	return nil
}

func (c *Client) InsertOne(collection string, doc interface{}) error {
	db := c.client.Database(c.database)
	col := db.Collection(collection)
	_, err := col.InsertOne(context.TODO(), doc)
	if err != nil {
		return err
	}
	return nil
}

func (c *Client) DeleteOne(collection string, filter interface{}) error {
	db := c.client.Database(c.database)
	col := db.Collection(collection)
	_, err := col.DeleteOne(context.TODO(), filter)
	if err != nil {
		return err
	}
	return nil
}

func (c *Client) Find(collection string, filter interface{}) []bson.M {
	db := c.client.Database(c.database)
	col := db.Collection(collection)

	log.Print("filter is ", filter)
	cur, err := col.Find(context.TODO(), filter)
	if err != nil {
		log.Fatal(err)
	}

	var results []bson.M
	if err = cur.All(context.TODO(), &results); err != nil {
		panic(err)
	}
	return results
}

func (c *Client) FindOne(collection string, filter interface{}) bson.M {
	log.Printf("Versao modificada")
	db := c.client.Database(c.database)
	col := db.Collection(collection)
	var result bson.M
	opts := options.FindOne().SetSort(bson.D{{"_id", 1}})
	log.Print("filter is ", filter)
	err := col.FindOne(context.TODO(), filter, opts).Decode(&result)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("found document %v", result)
	return result
}

func (c *Client) UpdateMany(collection string, filter interface{}, update interface{}) bson.M {
	log.Printf("Versao modificada")
	coll := c.client.Database(c.database).Collection(collection)

	log.Print("filter is ", filter)
	log.Print("update is ", update)
	result, err := coll.UpdateMany(context.TODO(), filter, update)
	if err != nil {
		log.Fatal(err)
		panic(err)
	}
	log.Print("update result", result)

	return nil
}
