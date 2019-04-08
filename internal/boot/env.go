package boot

import (
	"github.com/joho/godotenv"
	"github.com/magrathealabs/butia/internal/err"
)

// Env bootstrap
func Env() {
	err.RaiseErr(godotenv.Load())
}
