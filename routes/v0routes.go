package routes

import (
	delete "areTheyLeaving/controllers/DELETE"
	get "areTheyLeaving/controllers/GET"
	post "areTheyLeaving/controllers/POST"

	"github.com/gin-gonic/gin"
)

func v0Routes(route *gin.RouterGroup) {
	company := route.Group("/company")
	{
		company.POST("/create", post.CreateCompany)
		company.POST("/get", get.GetCompany)
	}
	user := route.Group("/user")
	{
		user.POST("/signup", post.CreateEmploye)
		user.POST("/login", post.UserLogin)
	}
	employee := route.Group("/employee")
	{
		employee.POST("/create", post.CreateEmploye)
		employee.POST("/get", post.GetEmployee)
		employee.POST("/delete", delete.RemoveEmploye)
		employee.POST("/getAll", post.GetAllEmployee)
	}

}
