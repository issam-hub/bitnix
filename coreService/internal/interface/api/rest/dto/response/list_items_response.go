package response

type ListItemsResponse struct {
	Items []ItemsResponse `json:"items"`
}

type ItemsResponse struct {
	ID        string  `json:"id"`
	Title     string  `json:"title"`
	Price     float64 `json:"price"`
	Thumbnail string  `json:"thumbnail"`
}
