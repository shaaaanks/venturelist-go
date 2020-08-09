package main

type project struct {
	Username       string   `json:"username"`
	UserID         string   `json:"user_id"`
	ProfilePicture string   `json:"profile_picture"`
	Name           string   `json:"name"`
	Description    string   `json:"descrition"`
	Summary        string   `json:"summary"`
	WebsiteLink    string   `json:"website_link"`
	GithubLink     string   `json:"github_link"`
	TechStack      []string `json:"tech_stack"`
	Mockups        []string `json:"mockups"`
	Screenshots    []string `json:"screenshotss"`
}
