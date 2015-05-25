package model

import "encoding/json"
import "encoding/base64"
import "testing"
import "io/ioutil"

import "gopkg.in/check.v1"

func TestReport(t *testing.T) {
	check.Suite(&ReportSuite{})
	check.TestingT(t)
}

type ReportSuite struct {
	image  []byte
	report *Report
}

func (s *ReportSuite) SetUpTest(c *check.C) {
	buf, err := ioutil.ReadFile("../image/fixture.jpg")
	c.Assert(err, check.IsNil)

	s.image = buf
	s.report = &Report{}
}

func (s *ReportSuite) TestUnmarshalJSON(c *check.C) {
	err := json.Unmarshal(s.json(), s.report)
	c.Assert(err, check.IsNil)

	c.Check(s.report.Input.Format, check.Equals, "jpeg")
	c.Check(s.report.Input.Bytes, check.DeepEquals, s.image)
}

func (s *ReportSuite) json() []byte {
	return []byte(`{"input":"data:image/jpeg;base64,` +
		base64.StdEncoding.EncodeToString(s.image) + `"}`)
}
