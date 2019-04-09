package spec

import (
	"net/http"
	"net/http/httptest"
	"strings"
)

// ControllerSuite mock controller requests
type ControllerSuite struct {
	Suite
	Server http.Handler
}

// MockRequest on Server
func (suite *ControllerSuite) MockRequest(method, path, body string) *httptest.ResponseRecorder {
	req, _ := http.NewRequest(method, path, strings.NewReader(body))

	w := httptest.NewRecorder()
	suite.Server.ServeHTTP(w, req)
	return w
}

// GET on Server
func (suite *ControllerSuite) GET(path, body string) *httptest.ResponseRecorder {
	return suite.MockRequest("GET", path, body)
}

// POST on Server
func (suite *ControllerSuite) POST(path, body string) *httptest.ResponseRecorder {
	return suite.MockRequest("POST", path, body)
}

// PUT on Server
func (suite *ControllerSuite) PUT(path, body string) *httptest.ResponseRecorder {
	return suite.MockRequest("PUT", path, body)
}

// PATCH on Server
func (suite *ControllerSuite) PATCH(path, body string) *httptest.ResponseRecorder {
	return suite.MockRequest("PATCH", path, body)
}

// DELETE on Server
func (suite *ControllerSuite) DELETE(path, body string) *httptest.ResponseRecorder {
	return suite.MockRequest("DELETE", path, body)
}
