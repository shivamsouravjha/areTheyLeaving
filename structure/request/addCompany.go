package requestStruct

type AddCompany struct {
	CompanyId   string `json:"company_id" binding:"required"`
	CompanyName string `json:"company_name" binding:"required"`
	CreatorId   string `json:"company_id" binding:"required"`
}
