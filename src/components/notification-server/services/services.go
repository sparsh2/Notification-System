package services

import "github.com/sparsh2/notification-system/src/components/notification-server/types"

type NotificationService interface {
	setPreference(*types.SetPreferenceRequest) error
}
