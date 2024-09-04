package mongo

import (
	"context"

	"github.com/HiWARPs/cactus-backend/models"
	"go.mongodb.org/mongo-driver/mongo"
)

func (mh *MongoHandler) AddOneAdmin(a *models.Admin) (*mongo.InsertOneResult, error) {
	ctx, cancel := context.WithTimeout(context.Background(), APITimeout)
	defer cancel()

	return mh.Admins.InsertOne(ctx, a)
}

func (mh *MongoHandler) GetOneAdmin(a *models.Admin, filter interface{}) error {
	ctx, cancel := context.WithTimeout(context.Background(), APITimeout)
	defer cancel()

	return mh.Admins.FindOne(ctx, filter).Decode(a)
}
