package components

type User struct {
	Id      	  int 		`json:"user_id"`
	Email         string    `json:"email" validate:"required,email"`
	Password      string    `json:"password,omitempty" validate:"required,min=6"`
	FirstName     string    `json:"first_name" validate:"required"`
	LastName      string    `json:"last_name" validate:"required"`
	Role          string    `json:"user_role"`
	Authenticated bool      `json:"-"` // omitting 'Authenticated' attribute from json
}

func (p User) GetFullName() string {
	return p.FirstName + " " + p.LastName
}


