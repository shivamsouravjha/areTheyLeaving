package post

import (
	"areTheyLeaving/helper/db"
	requestStruct "areTheyLeaving/structure/request"
	responseStruct "areTheyLeaving/structure/response"
	"areTheyLeaving/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateEmploye(c *gin.Context) {
	userStruct := requestStruct.AddUser{}
	if err := c.ShouldBind(&userStruct); err != nil {
		c.JSON(422, utils.SendErrorResponse(err))
		return
	}
	err := db.CreateUser(c, userStruct, "user")
	if err != nil {
		resp := responseStruct.SuccessResponse{}
		resp.Message = err.Error()
		resp.Status = "false"
		c.JSON(http.StatusInternalServerError, resp)
		return
	}
	resp := responseStruct.SuccessResponse{}
	resp.Message = "User added successfully"
	resp.Status = "true"
	c.JSON(http.StatusOK, resp)
}
