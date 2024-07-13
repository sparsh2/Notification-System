package types

type Preference struct {
	Channel string `json:"channel"`
	Enable bool `json:"enable"`
}

type SetPreferenceRequest struct {
	Preferences []Preference `json:"preferences"`
	UserID string `json:"user_id"`
}

