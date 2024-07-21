package main

// import gin

import (
	"fmt"

	"github.com/sparsh2/notification-system/src/components/notification-server/config"
	"github.com/sparsh2/notification-system/src/components/notification-server/storage"
	"github.com/sparsh2/notification-system/src/components/notification-server/web"
)

func main() {
	// Load the configs
	config.LoadConfig()
	// Initialize the storage layer
	storage.InitStorage()

	// Start the server
	r := web.GetRouter()
	err := r.Run()
	defer storage.Storage.Close()
	if err != nil {
		fmt.Println("Error: ", err)
	}
}
