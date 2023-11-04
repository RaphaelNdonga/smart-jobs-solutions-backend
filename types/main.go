package types

type UserDetails struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Location string `json:"location"`
	UserType string `json:"userType"`
}

type UserDetailsDB struct {
	Username       string
	Email          string
	HashedPassword string
	Location       string
	UserType       string
}
