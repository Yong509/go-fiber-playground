package types

type (
	User struct {
		ID   int    `json:"ID"`
		Name string `json:"Name" validate:"required,min=3,max=32"`
		Age  int    `json:"Age" validate:"required,number"`
	}
)
