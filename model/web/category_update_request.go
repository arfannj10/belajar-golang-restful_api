package web

type CategoryUpdateRequest struct {
	Id   int    `va;idaet:"required"`
	Name string `validate:"required,max=200,min=1" json:"name"`
}