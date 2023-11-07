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

type ServiceProviderJSON struct {
	Id          string `json:"uuid"`
	Service     string `json:"service"`
	Description string `json:"description"`
}

type ServiceProvider struct {
	Id          string
	Service     string
	Description string
}

type ClientJSON struct {
	Id      string `json:"uuid"`
	Service string `json:"service"`
}

type Client struct {
	Id      string
	Service string
}
