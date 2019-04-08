package butia

import "github.com/magrathealabs/butia/internal/boot"

func init() {
	boot.Env()
	boot.Gin()
}
