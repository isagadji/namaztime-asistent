package alisa

type AlisaRequest struct {
	Meta struct {
		Locale     string `json:"locale"`
		Timezone   string `json:"timezone"`
		ClientId   string `json:"client_id"`
		Interfaces struct {
			Screen struct {
			} `json:"screen"`
			Payments struct {
			} `json:"payments"`
			AccountLinking struct {
			} `json:"account_linking"`
			AudioPlayer struct {
			} `json:"audio_player"`
		} `json:"interfaces"`
	} `json:"meta"`
	Request struct {
		Type string `json:"type"`
	} `json:"request"`
	Session struct {
		MessageId int    `json:"message_id"`
		SessionId string `json:"session_id"`
		SkillId   string `json:"skill_id"`
		UserId    string `json:"user_id"`
		User      struct {
			UserId      string `json:"user_id"`
			AccessToken string `json:"access_token"`
		} `json:"user"`
		Application struct {
			ApplicationId string `json:"application_id"`
		} `json:"application"`
		New bool `json:"new"`
	} `json:"session"`
	State struct {
		Session struct {
			Value int `json:"value"`
		} `json:"session"`
		User struct {
			Value int `json:"value"`
		} `json:"user"`
		Application struct {
			Value int `json:"value"`
		} `json:"application"`
	} `json:"state"`
	Version string `json:"version"`
}
