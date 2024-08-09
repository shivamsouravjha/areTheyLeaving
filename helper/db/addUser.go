package db

import (
	"areTheyLeaving/helper/mongo"
	requestStruct "areTheyLeaving/structure/request"
	"context"
	"fmt"
)

func CreateUser(ctx context.Context, item requestStruct.AddUser, collection string) error {
	_, err := mongo.Client().Database("areTheyLeaving").Collection(collection).InsertOne(ctx, item)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	return nil
}
