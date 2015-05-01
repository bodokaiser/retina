package image

import "testing"

import "gopkg.in/check.v1"

func TestImage(t *testing.T) {
	check.Suite(&ImageSuite{
		base64: []byte(`"data:image/jpeg;base64,SGVsbG8="`),
	})
	check.TestingT(t)
}

type ImageSuite struct {
	image  *Image
	base64 []byte
}

func (s *ImageSuite) SetUpTest(c *check.C) {
	s.image = new(Image)
}

func (s *ImageSuite) TestProcess(c *check.C) {
	Process("../../medex/retina.jpg", "/tmp/foo.jpg")
}

func (s *ImageSuite) TestUnmarshalJSON(c *check.C) {
	err := s.image.UnmarshalJSON(s.base64)

	c.Assert(err, check.IsNil)
	c.Check(s.image.Format, check.Equals, "jpeg")
	c.Check(s.image.Bytes, check.DeepEquals, []byte("Hello"))
}
