package main

import "github.com/magrathealabs/butia/example/controllers"

func main() {
	controllers.NewApplicationController().BuildRoutes().Run("0.0.0.0:4000")
}
