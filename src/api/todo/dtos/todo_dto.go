package dtos

// swagger:parameters CreateTodoRequest
type CreateTodoRequest struct {
	// required: true
	Name string `form:"name" json:"name" xml:"name"  binding:"required,min=1,max=300"`
	// required: true
	Completed bool `form:"completed" json:"completed" xml:"completed"`
}
