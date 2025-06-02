package apiclient

type userResponse struct {
	User        User        `json:"data"`
	SupportInfo SupportInfo `json:"support"`
}

type User struct {
	ID        int64  `json:"id"`
	Email     string `json:"email"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	AvatarURL string `json:"avatar"`
}

type SupportInfo struct {
	URL  string `json:"url"`
	Text string `json:"text"`
}
