package types

type Preference struct {
	Channel string `json:"channel"`
	Enable bool `json:"enable"`
}

type SetPreferenceRequest struct {
	Preferences []Preference `json:"preferences"`
	ServiceID string `json:"service_id"`
	UserID string `json:"user_id"`
}

type SetUserDetailsRequest struct {
	Email string `json:"email"`
	UserID string `json:"user_id"`
	ServiceID string `json:"service_id"`
}
