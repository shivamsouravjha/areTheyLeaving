package routes

import (
	delete "areTheyLeaving/controllers/DELETE"
	post "areTheyLeaving/controllers/POST"

	"github.com/gin-gonic/gin"
)

func v0Routes(route *gin.RouterGroup) {
	company := route.Group("/company")
	{
		company.POST("/create", post.CreateCompany)
	}
	user := route.Group("/user")
	{
		user.POST("/signup", post.CreateEmploye)
		user.POST("/login", post.UserLogin)
	}
	employee := route.Group("/employee")
	{
		// employee.POST("/add", post.CreateEmploye)
		employee.POST("/get", post.GetEmployee)
		employee.POST("/delete", delete.RemoveEmploye)
		employee.POST("/getAll", post.GetAllEmployee)
	}
	employeeData := route.Group("/employeeData")
	{
		// employee.POST("/add", post.CreateEmploye)
		employeeData.POST("/get", post.GetEmployee)
		employeeData.POST("/delete", delete.RemoveEmploye)
		employeeData.POST("/getAll", post.GetAllEmployee)
	}
}
