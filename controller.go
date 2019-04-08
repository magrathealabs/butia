package butia

// Controller can handle all CRUD operations for any model
type Controller interface {
	Index(c *Context)
	Create(c *Context)
	New(c *Context)
	Edit(c *Context)
	Show(c *Context)
	Update(c *Context)
	Destroy(c *Context)
}
