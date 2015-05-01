package model

import (
	"encoding/json"
	"testing"
)

import "gopkg.in/check.v1"

func TestReport(t *testing.T) {
	check.Suite(&ReportSuite{
		json: []byte(`{"image":"data:image/jpeg;base64,SGVsbG8="}`),
	})
	check.TestingT(t)
}

type ReportSuite struct {
	report *Report
	json   []byte
}

func (s *ReportSuite) SetUpTest(c *check.C) {
	s.report = &Report{Image: new(Image)}
}

func (s *ReportSuite) TestUnmarshalJSON(c *check.C) {
	err := json.Unmarshal(s.json, s.report)

	c.Assert(err, check.IsNil)
	c.Check(s.report.Image.Format, check.Equals, "jpeg")
	c.Check(s.report.Image.Bytes, check.DeepEquals, []byte("Hello"))
}
