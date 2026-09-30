package post

type UpdatePostRequest struct {
	Title *string `josn:"title"`
	Body  *string `json:"body"`
}
