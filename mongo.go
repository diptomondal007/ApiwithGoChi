package ApiwithGoChi

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
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
