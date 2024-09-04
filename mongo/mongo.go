package mongo

import (
	"context"
	"fmt"

	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"

	"go.mongodb.org/mongo-driver/bson"

	"github.com/HiWARPs/cactus-backend/models"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const DefaultDatabase = "test"

const ProjectsCollectionName = "projects"

const ElectronsCollectionName = "electrons"

const AdminsCollectionName = "admins"

const APITimeout = 1000 * time.Second

type MongoHandler struct {
	client   *mongo.Client
	database string

	Projects  *mongo.Collection
	Electrons *mongo.Collection
	Admins    *mongo.Collection
}

// MongoHandler Constructor
func NewHandler(address string) (*MongoHandler, error) {
	ctx, cancel := context.WithTimeout(context.Background(), APITimeout)
	defer cancel()
	cl, err := mongo.Connect(ctx, options.Client().ApplyURI(address))

	if err != nil {
		return nil, err
	}

	mh := &MongoHandler{
		client:   cl,
		database: DefaultDatabase,
	}

	mh.Projects = mh.client.Database(mh.database).Collection(ProjectsCollectionName)
	mh.Electrons = mh.client.Database(mh.database).Collection(ElectronsCollectionName)
	mh.Admins = mh.client.Database(mh.database).Collection(AdminsCollectionName)

	return mh, nil
}

func (mh *MongoHandler) GetOne(c *models.ElectronData, filter interface{}) error {
	ctx, cancel := context.WithTimeout(context.Background(), APITimeout)
	defer cancel()
	return mh.Electrons.FindOne(ctx, filter).Decode(c)
}

func (mh *MongoHandler) Get(filter interface{}) ([]*models.Project, error) {
	ctx, cancel := context.WithTimeout(context.Background(), APITimeout)
	defer cancel()

	cur, err := mh.Electrons.Find(ctx, filter)

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

func (mh *MongoHandler) Upload(doc interface{}) (*mongo.InsertOneResult, error) {
	ctx, cancel := context.WithTimeout(context.Background(), APITimeout)
	defer cancel()

	return mh.Electrons.InsertOne(ctx, doc)
}

func (mh *MongoHandler) QueryElectrons(filter interface{}) ([]models.ElectronData, error) {
	ctx, cancel := context.WithTimeout(context.Background(), APITimeout)
	defer cancel()

	//cur, err := mh.Electrons.Find(ctx, filter)
	cur, err := mh.Electrons.Find(ctx, bson.M{})
	if err != nil {
		return nil, fmt.Errorf("unable to find collection, %s", err)
	}
	defer cur.Close(ctx)
	fmt.Println(cur)

	var result []models.ElectronData
	for cur.Next(ctx) {
		var request models.ElectronData
		err := cur.Decode(request)
		if err != nil {
			return nil, err
		}
		result = append(result, request)
	}

	return result, nil
}

func FilterIncrements(increment, lowerBound, upperBound float64) []float64 {
	i := lowerBound
	var incrementedValues []float64

	for i <= upperBound+increment {
		rounded := float64(int(i*100)) / 100
		incrementedValues = append(incrementedValues, rounded)
		i = i + increment
	}

	return incrementedValues
}

func FilterData(result models.ElectronData, filter models.Request) []map[string]float64 {
	filteredData := []map[string]float64{}
	newFilteredData := []map[string]float64{}
	incrementList := FilterIncrements(filter.Range.Increment, filter.Range.LowerBound, filter.Range.UpperBound)

	for _, item := range result.Data {

		for _, toDelete := range filter.Exclude {
			delete(item, toDelete)
		}

		for _, value := range incrementList {

			// check if the x value is found in the incrementedValues
			if item[filter.Range.Name] == value && item[filter.Range.Name] < filter.Range.UpperBound+filter.Range.Increment {
				filteredData = append(filteredData, item)
			}
		}
	}

	if len(filter.X) != 0 {
		for _, item := range filteredData {
			for _, xQuery := range filter.X {
				if item[xQuery.Name] == xQuery.Value {
					newFilteredData = append(newFilteredData, item)
				}
			}
		}

		return newFilteredData
	}

	return filteredData
}

func AggregateFilter(result models.ElectronData, filter models.AggregateRequest) []map[string][]float64 {

	keyValuesMap := make(map[string][]float64)

	for _, m := range result.Data {

		for key, value := range m {

			keyValuesMap[key] = append(keyValuesMap[key], value)
		}
	}

	var newResultData []map[string][]float64
	for key, values := range keyValuesMap {
		newResultData = append(newResultData, map[string][]float64{key: values})
	}

	println("filtered data")
	return newResultData
}

func (mh *MongoHandler) AggregateElectron(query models.AggregateRequest, filter interface{}) (models.AggregateData, error) {
	ctx, cancel := context.WithTimeout(context.Background(), APITimeout)
	defer cancel()

	d := models.ElectronData{}
	mh.Electrons.FindOne(ctx, filter).Decode(&d)

	if d.ID.IsZero() {
		return models.AggregateData{}, fmt.Errorf("no data found")
	}
	newData := models.AggregateData{}
	newData.Data = AggregateFilter(d, query)
	newData.ID = d.ID

	return newData, nil
}

func GetAllColumnValues(result []models.ElectronData, dataName string, dataLB, dataUB, dataIncrement float64) []map[string]float64 {
	var dataRes []map[string]float64

	for _, doc := range result {
		for _, dataItem := range doc.Data {
			if dataItem[dataName] >= dataLB && dataItem[dataName] <= dataUB {
				dataRes = append(dataRes, dataItem)
			}
		}
	}

	return dataRes
}

func (mh *MongoHandler) QueryElectron(query models.Request) (models.ElectronData, error) {
	ctx, cancel := context.WithTimeout(context.Background(), APITimeout)
	defer cancel()

	objectID, err := primitive.ObjectIDFromHex(query.ID)
	if err != nil {
		return models.ElectronData{}, err
	}

	d := models.ElectronData{}
	mh.Electrons.FindOne(ctx, bson.M{"_id": objectID}).Decode(&d)

	if d.ID.IsZero() {
		return models.ElectronData{}, fmt.Errorf("no data found")
	}

	d.Data = FilterData(d, query)

	return d, nil
}

func (mh *MongoHandler) AddOne(c map[string]interface{}) (*mongo.InsertOneResult, error) {
	ctx, cancel := context.WithTimeout(context.Background(), APITimeout)
	defer cancel()

	return mh.Electrons.InsertOne(ctx, c)
}

func (mh *MongoHandler) Update(filter interface{}, update interface{}) (*mongo.UpdateResult, error) {
	ctx, cancel := context.WithTimeout(context.Background(), APITimeout)
	defer cancel()

	return mh.Electrons.UpdateOne(ctx, filter, update)
}

func (mh *MongoHandler) RemoveOne(filter interface{}) (*mongo.DeleteResult, error) {
	ctx, cancel := context.WithTimeout(context.Background(), APITimeout)
	defer cancel()

	return mh.Electrons.DeleteOne(ctx, filter)
}
