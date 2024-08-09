package post

import (
	"areTheyLeaving/structure"
	responseStruct "areTheyLeaving/structure/response"
	"areTheyLeaving/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func UserLogin(c *gin.Context) {
	userStruct := structure.User{}
	if err := c.ShouldBind(&userStruct); err != nil {
		c.JSON(422, utils.SendErrorResponse(err))
		return
	}
	err := utils.Login(c, userStruct)
	if err != nil {
		resp := responseStruct.SuccessResponse{}
		resp.Message = err.Error()
		resp.Status = "false"
		c.JSON(http.StatusInternalServerError, resp)
		return
	}
	resp := responseStruct.SuccessResponse{}
	resp.Message = "User logged in successfully"
	resp.Status = "true"
	c.JSON(http.StatusOK, resp)
}
