package boot

import (
	"os"

	"github.com/joho/godotenv"
	"github.com/magrathealabs/butia/internal/err"
)

// Env bootstrap
func Env() {
	err.RaiseErr(godotenv.Load())
}

// Getenv or otherValue by default
func Getenv(key string, otherValue string) (value string) {
	value = os.Getenv(key)

	if value != "" {
		return
	}

	return otherValue
}
