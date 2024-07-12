package main

// import gin

import (
	"github.com/sparsh2/notification-system/components/notification-server/web"
)

func main() {
	// Start the server
	r := web.GetRouter()
	r.Run()
}
