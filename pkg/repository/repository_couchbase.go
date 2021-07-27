package repository

import (
	"context"
	"log"
	"time"

	"github.com/bino1490/crypto-svc/pkg/config"
	"github.com/bino1490/crypto-svc/pkg/entity"
	"github.com/bino1490/crypto-svc/pkg/logger"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	cbConnStr = config.SrvConfig.GetString(
		"database.mongodb.connectionstring")
	cbTable = config.SrvConfig.GetString(
		"database.mongodb.collection")
	cbDatabase = config.SrvConfig.GetString(
		"database.mongodb.db")
)

type CbRepository struct {
	Collection *mongo.Collection
}

//NewCbRepository to initialize the MonoDB connection
//Conntects with respective Table
func NewCbRepository() *CbRepository {
	logger.BootstrapLogger.Debug("Entering Repository.NewCbRepository() ...")

	clientOptions := options.Client().ApplyURI(cbConnStr)
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		logger.BootstrapLogger.Error(err)
		panic(err)
	}
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		logger.BootstrapLogger.Error(err)
		panic(err)
	}
	logger.BootstrapLogger.Debug("Connected to MongoDB!")
	collection := client.Database(cbDatabase).Collection(cbTable)
	return &CbRepository{
		Collection: collection,
	}
}

//GetDBRecords to fetch the datas
func (r *CbRepository) GetDBRecords(request entity.DBRequest) ([]entity.DBRecord, error) {
	DateMonthFormatConst := "2006-01-02"
	findOptions := options.Find()
	findOptions.SetLimit(25)
	var records []entity.DBRecord
	std, _ := time.Parse(DateMonthFormatConst, request.StartDate)
	edt, _ := time.Parse(DateMonthFormatConst, request.EndDate)
	cur, err := r.Collection.Find(context.TODO(), bson.M{
		"createdAt": bson.M{
			"$gte": std,
			"$lte": edt,
		},
	}, findOptions)
	if err != nil {
		return nil, err
	}
	for cur.Next(context.TODO()) {
		var rec entity.DBRecord
		err := cur.Decode(&rec)
		if err != nil {
			log.Fatal(err)
		}
		tc := 0
		for _, value := range rec.Counts {
			tc = tc + value
		}
		rec.TotalCount = tc
		rec.Counts = make([]int, 0)
		if (rec.TotalCount >= request.MinCount) && (rec.TotalCount <= request.MaxCount) {
			records = append(records, rec)
		}

	}
	logger.Logger.Debug("Found multiple documents (array of pointers): %+v\n", records)
	// Close the cursor once finished
	defer cur.Close(context.TODO())
	return records, nil
}
