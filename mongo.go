package ApiwithGoChi

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const DefaultDatabase = "PersonDataBase"
const DefaultCollection = "person"

type MongoHandler struct {
	client   *mongo.Client
	database string
}

func NewHandler(address string) *MongoHandler {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	cl, _ := mongo.Connect(ctx, options.Client().ApplyURI(address))
	mh := &MongoHandler{
		client:   cl,
		database: DefaultDatabase,
	}
	return mh
}

func (mh *MongoHandler) GetOne(p *Person, filter interface{}) error {
	collection := mh.client.Database(DefaultDatabase).Collection(DefaultCollection)
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	err := collection.FindOne(ctx, filter).Decode(p)
	return err
}

func (mh *MongoHandler) AddOne(p *Person) (*mongo.InsertOneResult, error) {
	collection := mh.client.Database(mh.database).Collection(DefaultCollection)
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	result, err := collection.InsertOne(ctx, p)
	return result, err
}

func (mh *MongoHandler) Get(filter interface{}) []*Person {
	collection := mh.client.Database(mh.database).Collection(DefaultCollection)
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	cur, err := collection.Find(ctx, filter)
	if err != nil {
		log.Fatal(err)
	}
	defer cur.Close(ctx)
	var result []*Person
	for cur.Next(ctx) {
		person := &Person{}
		er := cur.Decode(person)
		if er != nil {
			log.Fatal(er)
		}
		result = append(result, person)
	}
	return result
}
