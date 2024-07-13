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

