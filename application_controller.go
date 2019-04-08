package butia

// ApplicationController implement a set of basic method to helpfull the controller management
type ApplicationController struct {
	Server *Server
}

// NewApplicationController instance ApplicationController as Controller interface
func NewApplicationController(server *Server) *ApplicationController {
	return &ApplicationController{server}
}

// Setup this controller into the server
func (controller *ApplicationController) Setup() {
	controller.Server.Controller("/", controller)
}

// Index a basic model over GET method
func (controller *ApplicationController) Index(c *Context) {
	c.String(200, "Index must be implemented")
}

// Create a basic instance over POST method
func (controller *ApplicationController) Create(c *Context) {
	c.String(200, "Create must be implemented")
}

// New a basic instance over GET method
func (controller *ApplicationController) New(c *Context) {
	c.String(200, "New must be implemented")
}

// Edit a basic instance over GET method
func (controller *ApplicationController) Edit(c *Context) {
	c.String(200, "Edit must be implemented for %s", c.Param("id"))
}

// Show a basic instance over GET method
func (controller *ApplicationController) Show(c *Context) {
	c.String(200, "Show must be implemented for %s", c.Param("id"))
}

// Update a basic instance over PATCH or PUT method
func (controller *ApplicationController) Update(c *Context) {
	c.String(200, "Update must be implemented for %s", c.Param("id"))
}

// Destroy a basic instance over DELETE method
func (controller *ApplicationController) Destroy(c *Context) {
	c.String(200, "Destroy must be implemented for %s", c.Param("id"))
}
