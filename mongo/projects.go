package mongo

import (
	"context"

	"github.com/HiWARPs/cactus-backend/models"
	"go.mongodb.org/mongo-driver/mongo"
)

func (mh *MongoHandler) GetOneProject(c *models.Project, filter interface{}) error {
	ctx, cancel := context.WithTimeout(context.Background(), APITimeout)
	defer cancel()
	return mh.Projects.FindOne(ctx, filter).Decode(c)
}

func (mh *MongoHandler) GetProjects(filter interface{}) ([]*models.Project, error) {
	ctx, cancel := context.WithTimeout(context.Background(), APITimeout)
	defer cancel()

	cur, err := mh.Projects.Find(ctx, filter)

	if err != nil {
		return nil, err
	}
	defer cur.Close(ctx)

	var result []*models.Project
	for cur.Next(ctx) {
		project := &models.Project{}
		er := cur.Decode(project)
		if er != nil {
			return nil, er
		}
		result = append(result, project)
	}
	return result, nil
}

func (mh *MongoHandler) AddOneProject(c *models.ProjectCreate) (*mongo.InsertOneResult, error) {
	ctx, cancel := context.WithTimeout(context.Background(), APITimeout)
	defer cancel()

	return mh.Projects.InsertOne(ctx, c)
}

func (mh *MongoHandler) UpdateProject(filter interface{}, update interface{}) (*mongo.UpdateResult, error) {
	ctx, cancel := context.WithTimeout(context.Background(), APITimeout)
	defer cancel()

	return mh.Projects.UpdateOne(ctx, filter, update)
}

func (mh *MongoHandler) RemoveOneProject(filter interface{}) (*mongo.DeleteResult, error) {
	ctx, cancel := context.WithTimeout(context.Background(), APITimeout)
	defer cancel()

	return mh.Projects.DeleteOne(ctx, filter)
}
