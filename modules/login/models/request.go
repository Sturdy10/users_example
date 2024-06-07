package models

type LoginRequest struct {
	OrgmbEmail  string `json:"email"`
	OrgmbcrPassword string `json:"password"`
}
