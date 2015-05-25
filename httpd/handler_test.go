package httpd

import "bytes"
import "net/http"
import "net/http/httptest"
import "testing"
import "io/ioutil"
import "encoding/json"

import "gopkg.in/check.v1"

import . "github.com/bodokaiser/retina/model"
import . "github.com/bodokaiser/retina/image"

func TestHandler(t *testing.T) {
	check.Suite(&HandlerSuite{})
	check.TestingT(t)
}

type HandlerSuite struct {
	report  *Report
	handler *Handler
}

func (s *HandlerSuite) SetUpTest(c *check.C) {
	buf, err := ioutil.ReadFile("../image/fixture.jpg")
	c.Assert(err, check.IsNil)

	s.report = &Report{
		Input: &Image{
			Bytes:  buf,
			Format: "jpeg",
		},
	}
	s.handler = NewHandler()
}

func (s *HandlerSuite) TestUpload(c *check.C) {
	res := httptest.NewRecorder()

	s.handler.ServeHTTP(res, s.request())
	s.report.Process()

	//json, _ := json.Marshal(s.report)

	c.Check(res.Code, check.Equals, http.StatusOK)
	//c.Check(res.Body.Bytes(), check.DeepEquals, json)
}

func (s *HandlerSuite) request() *http.Request {
	buf, _ := json.Marshal(s.report)
	req, _ := http.NewRequest("POST", "/", bytes.NewReader(buf))
	req.Header.Set("Content-Type", "application/json")

	return req
}
