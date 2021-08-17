package repository

import (
	"context"
	"fmt"
	"github.com/MahmudulTushar/graphql/graph/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"os"
	"time"
)

const DatasBase = "GraphQl"
const CourseCollection = "Course"

type CourseRepository interface {
	Save(course *model.Course)
	FindAll() []*model.Course
	UpdateById(id string, course *model.Course) string
	Delete(id string) string
}

type database struct {
	client *mongo.Client
}

func (db *database) Save(course *model.Course) {
	collection := db.client.Database(DatasBase).Collection(CourseCollection)
	_, err := collection.InsertOne(context.TODO(), course)
	if err != nil {
		log.Fatal(err)
	}

}

func (db *database) Delete(id string) string {
	collection := db.client.Database(DatasBase).Collection(CourseCollection)
	res, err := collection.DeleteOne(context.Background(), bson.M{"_id": id})
	if err != nil {
		log.Fatal(err)
		return "Not Deleted"
	}

	if res.DeletedCount > 0 {
		return "Deleted"
	}

	return "No record has been deleted"
}

func (db *database) UpdateById(id string, course *model.Course) string {
	// Declare an _id filter to get a specific MongoDB document
	filter := bson.M{"_id": bson.M{"$eq": id}}
	// Declare a filter that will change a field's integer value to `42`
	update := bson.M{"$set": bson.M{"name": course.Name, "description": course.Description}}
	collection := db.client.Database(DatasBase).Collection(CourseCollection)
	// Call the driver's UpdateOne() method and pass filter and update to it
	res, err := collection.UpdateOne(
		context.Background(),
		filter,
		update,
	)
	if err != nil {
		log.Fatal(err)
		return "Error while updating the record"
	}
	if res.ModifiedCount > 0 {
		return "Record Update"
	}
	return "No record found"
}

func (db *database) FindAll() []*model.Course {
	collection := db.client.Database(DatasBase).Collection(CourseCollection)
	cursor, err := collection.Find(context.TODO(), bson.D{})
	if err != nil {
		log.Fatal(err)
	}
	defer cursor.Close(context.TODO())
	var result []*model.Course
	for cursor.Next(context.TODO()) {
		var v *model.Course
		err := cursor.Decode(&v)
		if err != nil {
			log.Fatal(err)
		}
		result = append(result, v)
	}

	return result
}

func NewDatabaseInstance() CourseRepository {
	mongoDb := os.Getenv("MONGODB_URI")
	clientOptions := options.Client().ApplyURI(mongoDb)
	clientOptions = clientOptions.SetMaxPoolSize(50)
	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	dbClient, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connect to Db!")
	return &database{
		client: dbClient,
	}
}
