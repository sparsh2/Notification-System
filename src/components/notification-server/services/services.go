package services

import (
	"github.com/sparsh2/notification-system/src/components/notification-server/storage"
	"github.com/sparsh2/notification-system/src/components/notification-server/types"
)

type INotificationService interface {
	SetPreference(*types.SetPreferenceRequest) error
	SetUserDetails(*types.SetUserDetailsRequest) error
}

var Service INotificationService

func init() {
	Service = &NotificationService{}
}

type NotificationService struct {
}

func (ns *NotificationService) SetPreference(req *types.SetPreferenceRequest) error {
	err := storage.Storage.SetPreference(req)
	if err != nil {
		return err
	}
	return nil
}

func (ns *NotificationService) SetUserDetails(req *types.SetUserDetailsRequest) error {
	err := storage.Storage.SetUserDetails(req)
	if err != nil {
		return err
	}
	return nil
}
