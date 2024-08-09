package get

// func GetCompany(c *gin.Context) {
// 	companyId := c.Query("companyId")
// 	if companyId == "" {
// 		c.JSON(422, utils.SendErrorResponse(errors.New("Can't find company")))
// 	}

// 	resp := responseStruct.CompanyResponse{}
// 	objID, _ := primitive.ObjectIDFromHex(companyId)

// 	filter := bson.M{"_id": objID}

// 	existingCompany, err := db.Find(filter, "company")
// 	if err != nil {
// 		c.JSON(422, utils.SendErrorResponse(err))
// 		return

// 	}
// 	resp.Company = *existingCompany
// 	resp.Status = "true"
// 	c.JSON(http.StatusOK, resp)
// }
