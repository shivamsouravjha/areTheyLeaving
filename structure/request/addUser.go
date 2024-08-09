package requestStruct

type AddUser struct {
	UserId    string `json:"user_id" binding:"required"`
	CompanyId string `json:"company_id"`
	Name      string `json:"user_name" binding:"required"`
	Email     string `json:"user_email" binding:"required"`
	Password  string `json:"user_password" binding:"required"`
}
