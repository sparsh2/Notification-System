package main

// import gin

import (
	"github.com/sparsh2/notification-system/src/components/notification-server/config"
	"github.com/sparsh2/notification-system/src/components/notification-server/web"
)

func main() {
	// Load the configs
	config.LoadConfig()

	// Start the server
	r := web.GetRouter()
	r.Run()
}
