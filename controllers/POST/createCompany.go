package post

import (
	"areTheyLeaving/helper/db"
	requestStruct "areTheyLeaving/structure/request"
	responseStruct "areTheyLeaving/structure/response"
	"net/http"

	"areTheyLeaving/utils"

	"github.com/gin-gonic/gin"
)

func CreateCompany(c *gin.Context) {
	companyStruct := requestStruct.AddCompany{}
	if err := c.ShouldBind(&companyStruct); err != nil {
		c.JSON(422, utils.SendErrorResponse(err))
		return
	}
	err := db.CreateItem(c, companyStruct, "company")
	if err != nil {
		resp := responseStruct.SuccessResponse{}
		resp.Message = err.Error()
		resp.Status = "false"
		c.JSON(http.StatusInternalServerError, resp)
		return
	}
	resp := responseStruct.SuccessResponse{}
	resp.Message = "Company added successfully"
	resp.Status = "true"
	c.JSON(http.StatusOK, resp)
}
