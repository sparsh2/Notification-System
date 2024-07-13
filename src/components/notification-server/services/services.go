package services

import (
	"database/sql"
	"encoding/json"

	"github.com/go-sql-driver/mysql"
	"github.com/sparsh2/notification-system/src/components/notification-server/config"
	"github.com/sparsh2/notification-system/src/components/notification-server/types"
)

type INotificationService interface {
	SetPreference(*types.SetPreferenceRequest) error
}

var Service INotificationService

func init() {
	Service = &NotificationService{}
}


type NotificationService struct {
}

func (ns *NotificationService) SetPreference(req *types.SetPreferenceRequest) error {
	cfg := mysql.Config{
		User:   config.Configs.DBConfig.User,
		Passwd: config.Configs.DBConfig.Password,
		Net:    "tcp",
		Addr:   config.Configs.DBConfig.Host,
		DBName: config.Configs.DBConfig.DBName,
	}
	db, err := sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		return err
	}
	defer db.Close()
	preferences, err := json.Marshal(req.Preferences)
	if err != nil {
		return err
	}
	_, err = db.Exec("DELETE FROM user_preferences WHERE user_id = ? AND service_id = ?", req.UserID, req.ServiceID)
	if err != nil {
		return err
	}
	_, err = db.Exec("INSERT INTO user_preferences (user_id, service_id, preferences) VALUES (?, ?, ?)",
		req.UserID, req.ServiceID, string(preferences))
	if err != nil {
		return err
	}
	return nil
}
