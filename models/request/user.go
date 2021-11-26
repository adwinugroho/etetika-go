package request

type (
	User struct {
		Email    string   `form:"email"`
		FullName string   `json:"fullName"`
		Phone    string   `json:"phone"`
		Address  string   `json:"address"`
		Role     []string `json:"roles"`
	}
)
