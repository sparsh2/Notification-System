package storage

import (
	"database/sql"
	"encoding/json"
	"fmt"

	"github.com/go-sql-driver/mysql"
	"github.com/sparsh2/notification-system/src/components/notification-server/config"
	"github.com/sparsh2/notification-system/src/components/notification-server/types"
)

type IStorage interface {
	SetPreference(*types.SetPreferenceRequest) error
	SetUserDetails(*types.SetUserDetailsRequest) error
	Close()
}

type NotificationStorage struct {
	dbConn *sql.DB
}

var Storage IStorage

func InitStorage() {
	cfg := mysql.Config{
		User:   config.Configs.DBConfig.User,
		Passwd: config.Configs.DBConfig.Password,
		Net:    "tcp",
		Addr:   config.Configs.DBConfig.Host,
		DBName: config.Configs.DBConfig.DBName,
	}
	db, err := sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		panic(err)
	}
	Storage = &NotificationStorage{
		dbConn: db,
	}
}

func (ns *NotificationStorage) SetPreference(req *types.SetPreferenceRequest) error {
	preferences, err := json.Marshal(req.Preferences)
	if err != nil {
		return err
	}

	// error out if there is no entry in the table for the given user_id and service_id
	var count int
	err = ns.dbConn.QueryRow(
		"SELECT COUNT(*) FROM notifications WHERE user_id = ? AND service_id = ?",
		req.UserID, req.ServiceID,
	).Scan(&count)
	if err != nil {
		return err
	}
	if count == 0 {
		return fmt.Errorf("no entry found for user_id: %s", req.UserID)
	}

	_, err = ns.dbConn.Exec(
		"UPDATE notifications SET preferences = ? WHERE user_id = ? AND service_id = ?",
		preferences, req.UserID, req.ServiceID,
	)
	if err != nil {
		return err
	}
	return nil
}

func (ns *NotificationStorage) SetUserDetails(req *types.SetUserDetailsRequest) error {
	// extract email into json string
	email, err := json.Marshal(req.Email)
	if err != nil {
		return fmt.Errorf("error marshalling email: %w", err)
	}
	_, err = ns.dbConn.Exec("INSERT INTO notifications (user_id, service_id, preferences, user_data) VALUES (?, ?, \"\", ?) ON DUPLICATE KEY UPDATE user_data = ?", req.UserID, req.ServiceID, email, email)
	if err != nil {
		return err
	}
	return nil
}

func (ns *NotificationStorage) Close() {
	ns.dbConn.Close()
}
