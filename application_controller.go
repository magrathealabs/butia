package butia

// ApplicationController implement a set of basic method to helpfull the controller management
type ApplicationController struct {
	*Server
}

// NewApplicationController instance ApplicationController as Controller interface
func NewApplicationController(server *Server) Controller {
	return &ApplicationController{server}
}

// Index a basic model over GET method
func (controller *ApplicationController) Index(c *Context) {
	c.String(404, "Not implemented")
}

// Create a basic instance over POST method
func (controller *ApplicationController) Create(c *Context) {
	c.String(404, "Not implemented")
}

// New a basic instance over GET method
func (controller *ApplicationController) New(c *Context) {
	c.String(404, "Not implemented")
}

// Edit a basic instance over GET method
func (controller *ApplicationController) Edit(c *Context) {
	c.String(404, "Not implemented")
}

// Show a basic instance over GET method
func (controller *ApplicationController) Show(c *Context) {
	c.String(404, "Not implemented")
}

// Update a basic instance over PATCH or PUT method
func (controller *ApplicationController) Update(c *Context) {
	c.String(404, "Not implemented")
}

// Destroy a basic instance over DELETE method
func (controller *ApplicationController) Destroy(c *Context) {
	c.String(404, "Not implemented")
}
