package server

import (
	"os"
)

// Init нь шинэ router үүсгэх
func Init() {
	router := newRouter()
	port := os.Getenv("PORT")
	_ = router.Run(port)
}
