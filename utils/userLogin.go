package utils

import (
	"areTheyLeaving/helper/db"
	"areTheyLeaving/structure"
	"context"
	"fmt"
)

func Login(c context.Context, userStruct structure.User) error {
	err, foundUser := db.FetchUser(c, userStruct, "user")
	if err != nil {
		return err
	}
	if foundUser.Password != userStruct.Password {
		return fmt.Errorf("Invalid password")
	}
	return nil
}
