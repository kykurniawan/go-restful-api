package web

type CategoryUpdateRequest struct {
	Id   int    `validate:"required" json:"id"`
	Name string `validate:"required,max=128,min=1" json:"name"`
}
