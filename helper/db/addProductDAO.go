package db

import (
	"areTheyLeaving/helper/mongo"
	"areTheyLeaving/structure"
	requestStruct "areTheyLeaving/structure/request"
	"context"
	"fmt"

	mongoose "go.mongodb.org/mongo-driver/mongo"

	"go.mongodb.org/mongo-driver/bson"
)

func CreateItem(ctx context.Context, item requestStruct.AddCompany, collection string) error {
	_, err := mongo.Client().Database("areTheyLeaving").Collection(collection).InsertOne(ctx, item)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	return nil
}

func Find(filter bson.M, collection string) (*structure.Company, error) {
	var existingUser structure.Company
	dbCollection := mongo.Client().Database("areTheyLeaving").Collection(collection)
	err := dbCollection.FindOne(context.Background(), filter).Decode(&existingUser)
	if err != nil {
		if err == mongoose.ErrNoDocuments {
			// No document found for the given filter
			fmt.Printf("No document found for filter: %+v\n", filter)
		} else {
			// Other error occurred
			fmt.Printf("Error while querying MongoDB: %v\n", err)
		}
		fmt.Println(filter)
		fmt.Println(err.Error())
		return nil, err
	}
	return &existingUser, nil
}
