package db

import (
	"areTheyLeaving/helper/mongo"
	"areTheyLeaving/structure"
	"context"
	"fmt"

	mongoose "go.mongodb.org/mongo-driver/mongo"

	"gopkg.in/mgo.v2/bson"
)

func FetchUser(ctx context.Context, user structure.User, collection string) (err error, foundUser structure.User) {
	filter := bson.M{"email": user.Email}
	err = mongo.Client().Database("areTheyLeaving").Collection(collection).FindOne(ctx, filter).Decode(&foundUser)
	if err != nil {
		if err == mongoose.ErrNoDocuments {
			return fmt.Errorf("user not found"), foundUser
		}
		return err, foundUser
	}
	return nil, foundUser
}
