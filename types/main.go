package types

type UserDetails struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Location string `json:"location"`
	UserType string `json:"userType"`
}

type UserDetailsDB struct {
	Id             string
	Username       string
	Email          string
	HashedPassword string
	Location       string
	UserType       string
}

type ProviderJSON struct {
	Id          string `json:"Id"`
	Service     string `json:"Service"`
	Description string `json:"Description"`
}

type Provider struct {
	Id          string
	Service     string
	Description string
}

type ProviderResponse struct {
	Service     string
	Description string
	Username    string
}

type ClientJSON struct {
	Id      string `json:"uuid"`
	Service string `json:"service"`
}

type Client struct {
	Id      string
	Service string
}

type PostJSON struct {
	Id      string `json:"Id"`
	Post    string `json:"Post"`
	Service string `json:"Service"`
}

type ProviderPostJSON struct {
	Id   string `json:"Id"`
	Post string `json:"post"`
}

type PostResponse struct {
	Id        string
	Username  string
	Post      string
	CreatedAt string
	Location  string
	Service   string
}

type ServiceJSON struct {
	Key_Service string `json:"key_service"`
}

type Service struct {
	Key_Service string
}

type CommentResponse struct {
	Username  string
	Comment   string
	CreatedAt string
}
