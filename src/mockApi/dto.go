package mockApi

type MockApiRequestType struct {
	Title string `json:"title" binding:"required"`
}
