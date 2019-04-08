package boot

import (
	"github.com/joho/godotenv"
	"github.com/magrathealabs/butia/internal/helper"
)

// Env bootstrap
func Env() {
	helper.RaiseErr(godotenv.Load())
}
