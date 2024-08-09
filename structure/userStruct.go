package structure

type User struct {
	Password string `json:"password" binding:"required"`
	Email    string `json:"mobile" binding:"required"`
}
