package types

type UserDetails struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
	UserType string `json:"userType"`
}

type UserDetailsDB struct {
	Username       string
	Email          string
	HashedPassword []byte
	UserType       string
}
