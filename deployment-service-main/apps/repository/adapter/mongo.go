package adapter

import (
	"context"
	"deployment-service/constants"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// GetAll retrieves all documents from the specified collection
func (m *MongoDB) GetAll(collection string, filter bson.D) (*mongo.Cursor, error) {
	col := m.connection.Database(constants.MONGODB_NAME).Collection(collection)

	cursor, err := col.Find(context.TODO(), filter)
	if err != nil {
		return nil, err
	}
	return cursor, nil
}

// InsertOne inserts a single document into the specified collection
func (m *MongoDB) InsertOne(collection string, document interface{}) (*mongo.InsertOneResult, error) {
	col := m.connection.Database(constants.MONGODB_NAME).Collection(collection)
	return col.InsertOne(context.TODO(), document)
}

// InsertMany inserts multiple documents into the specified collection
func (m *MongoDB) InsertMany(collection string, documents []interface{}) (*mongo.InsertManyResult, error) {
	col := m.connection.Database(constants.MONGODB_NAME).Collection(collection)
	return col.InsertMany(context.TODO(), documents)
}

// FindOne finds a single document in the specified collection
func (m *MongoDB) FindOne(collection string, filter bson.M) *mongo.SingleResult {
	col := m.connection.Database(constants.MONGODB_NAME).Collection(collection)
	return col.FindOne(context.TODO(), filter)
}

// FindMany finds multiple documents in the specified collection
func (m *MongoDB) FindMany(collection string, filter bson.M) (*mongo.Cursor, error) {
	col := m.connection.Database(constants.MONGODB_NAME).Collection(collection)
	return col.Find(context.TODO(), filter)
}

// UpdateOne updates a single document in the specified collection
func (m *MongoDB) UpdateOne(collection string, filter bson.M, update bson.M) (*mongo.UpdateResult, error) {
	col := m.connection.Database(constants.MONGODB_NAME).Collection(collection)
	return col.UpdateOne(context.TODO(), filter, update)
}

// UpdateMany updates multiple documents in the specified collection
func (m *MongoDB) UpdateMany(collection string, filter bson.M, update bson.M) (*mongo.UpdateResult, error) {
	col := m.connection.Database(constants.MONGODB_NAME).Collection(collection)
	return col.UpdateMany(context.TODO(), filter, update)
}

// DeleteOne deletes a single document from the specified collection
func (m *MongoDB) DeleteOne(collection string, filter bson.M) (*mongo.DeleteResult, error) {
	col := m.connection.Database(constants.MONGODB_NAME).Collection(collection)
	return col.DeleteOne(context.TODO(), filter)
}

// DeleteMany deletes multiple documents from the specified collection
func (m *MongoDB) DeleteMany(collection string, filter bson.M) (*mongo.DeleteResult, error) {
	col := m.connection.Database(constants.MONGODB_NAME).Collection(collection)
	return col.DeleteMany(context.TODO(), filter)
}

// CountDocuments counts the documents in the specified collection that match the filter
func (m *MongoDB) CountDocuments(collection string, filter bson.M) (int64, error) {
	col := m.connection.Database(constants.MONGODB_NAME).Collection(collection)
	return col.CountDocuments(context.TODO(), filter)
}

// Aggregate runs an aggregation pipeline on the specified collection
func (m *MongoDB) Aggregate(collection string, pipeline mongo.Pipeline) (*mongo.Cursor, error) {
	col := m.connection.Database(constants.MONGODB_NAME).Collection(collection)
	return col.Aggregate(context.TODO(), pipeline)
}
