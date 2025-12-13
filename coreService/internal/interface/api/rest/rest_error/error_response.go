package resterror

type ErrGameNotFoundResponse struct {
	Error string `json:"error" example:"game not found"`
}

type ErrInternal struct {
	Error string `json:"error" example:"Internal Server Error"`
}

type ErrGetGameBadRequest struct {
	Error string `json:"error" example:"invalid game ID format"`
}
