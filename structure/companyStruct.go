package structure

type Company struct {
	CompanyId   string `json:"company_id" binding:"required"`
	CompanyName string `json:"company_name" binding:"required"`
}
