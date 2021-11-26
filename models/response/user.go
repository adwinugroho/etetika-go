package response

type (
	User struct {
		Email    string   `json:"email"`
		FullName string   `json:"fullName"`
		Phone    string   `json:"phone"`
		Address  string   `json:"address"`
		Role     []string `json:"roles"`
	}
)
