package types

type Preference struct {
	Channel string `json:"channel"`
	Enable  bool   `json:"enable"`
}

type SetPreferenceRequest struct {
	Preferences []Preference `json:"preferences"`
	ServiceID   string       `json:"service_id"`
	UserID      string       `json:"user_id"`
}

type SetUserDetailsRequest struct {
	Email     string `json:"email"`
	UserID    string `json:"user_id"`
	ServiceID string `json:"service_id"`
}

type NotificationRequest struct {
	Title     string            `json:"title"`
	Message   string            `json:"message"`
	UserID    string            `json:"user_id"`
	ServiceID string            `json:"service_id"`
	Metadata  map[string]string `json:"metadata,omitempty"`
}

type NotificationResponse struct {
	Success bool   `json:"success"`
	Error   string `json:"error,omitempty"`
}
