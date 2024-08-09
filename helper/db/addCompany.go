package db

import (
	"areTheyLeaving/helper/mongo"
	requestStruct "areTheyLeaving/structure/request"
	"context"
	"fmt"
)

func CreateCompany(ctx context.Context, item requestStruct.AddCompany, collection string) error {
	_, err := mongo.Client().Database("areTheyLeaving").Collection(collection).InsertOne(ctx, item)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	return nil
}

// func FindCompan(filter bson.M, collection string) (*structure.Company, error) {
// 	var existingUser structure.Company
// 	dbCollection := mongo.Client().Database("areTheyLeaving").Collection(collection)
// 	err := dbCollection.FindOne(context.Background(), filter).Decode(&existingUser)
// 	if err != nil {
// 		if err == mongoose.ErrNoDocuments {
// 			// No document found for the given filter
// 			fmt.Printf("No document found for filter: %+v\n", filter)
// 		} else {
// 			// Other error occurred
// 			fmt.Printf("Error while querying MongoDB: %v\n", err)
// 		}
// 		fmt.Println(filter)
// 		fmt.Println(err.Error())
// 		return nil, err
// 	}
// 	return &existingUser, nil
// }
