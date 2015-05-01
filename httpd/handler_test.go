package httpd

import "net/http"
import "net/http/httptest"
import "strings"
import "testing"

import "gopkg.in/check.v1"

func TestHandler(t *testing.T) {
	check.Suite(&HandlerSuite{})
	check.TestingT(t)
}

type HandlerSuite struct {
	handler *Handler
}

func (s *HandlerSuite) SetUpTest(c *check.C) {
	s.handler = NewHandler()
}

func (s *HandlerSuite) TestUpload(c *check.C) {
	req, _ := http.NewRequest("POST", "/", strings.NewReader(`{"Handler":"SGVsbG8="}`))
	req.Header.Set("Content-Type", "application/json")
	res := httptest.NewRecorder()

	s.handler.ServeHTTP(res, req)

	c.Check(res.Code, check.Equals, http.StatusNoContent)
	c.Check(res.Body.String(), check.Equals, "")
}
