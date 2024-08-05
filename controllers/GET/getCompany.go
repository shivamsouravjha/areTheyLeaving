package get

import (
	"areTheyLeaving/helper/db"
	responseStruct "areTheyLeaving/structure/response"
	"areTheyLeaving/utils"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetCompany(c *gin.Context) {
	companyId := c.Query("companyId")
	if companyId == "" {
		c.JSON(422, utils.SendErrorResponse(errors.New("Can't find company")))
	}

	resp := responseStruct.CompanyResponse{}
	objID, _ := primitive.ObjectIDFromHex(companyId)

	filter := bson.M{"_id": objID}

	existingUser, err := db.Find(filter, "company")
	if err != nil {
		c.JSON(422, utils.SendErrorResponse(err))
		return

	}
	resp.Company = *existingUser
	resp.Status = "true"
	c.JSON(http.StatusOK, resp)
}
