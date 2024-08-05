package responseStruct

import "areTheyLeaving/structure"

type CompanyResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
	structure.Company
}
