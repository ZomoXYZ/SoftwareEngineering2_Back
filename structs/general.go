package structs

type ErrorJson struct {
	Error string `json:"error" binding:"required"`
}