package models

type User struct {
	CommonFields
	Username  string `json:"username,omitempty"`
	Password  string `json:"password,omitempty"`
	Fullname  string `json:"fullname,omitempty"`
	Nickname  string `json:"nickname,omitempty"`
	Avatar    string `json:"avatar,omitempty"`
	Email     string `json:"email,omitempty"`
	Education string `json:"education,omitempty"`
	Github    string `json:"github,omitempty"`
}
