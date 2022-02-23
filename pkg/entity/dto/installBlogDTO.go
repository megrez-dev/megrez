package dto

type InstallBlogDTO struct {
	BlogTitle       string `json:"blogTitle"`
	BlogURL         string `json:"blogURL"`
	Username        string `json:"username"`
	Nickname        string `json:"nickname"`
	Email           string `json:"email"`
	Password        string `json:"password"`
	ConfirmPassword string `json:"confirmPassword"`
}
