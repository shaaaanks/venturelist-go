package main

type project struct {
	Username       string   `json:"username"`
	UserID         string   `json:"user_id"`
	ProfilePicture string   `json:"profile_picture"`
	Name           string   `json:"name" validate:"required"`
	Description    string   `json:"descrition" validate:"required"`
	Summary        string   `json:"summary" validate:"required"`
	WebsiteLink    string   `json:"website_link" validate:"required"`
	GithubLink     string   `json:"github_link" validate:"required"`
	TechStack      []string `json:"tech_stack" validate:"required"`
	Mockups        []string `json:"mockups" validate:"required"`
	Screenshots    []string `json:"screenshots" validate:"required"`
}
