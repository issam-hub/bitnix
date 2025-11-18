package resterror

type ErrInternal struct {
	Error string `json:"error" example:"Internal Server Error"`
}
