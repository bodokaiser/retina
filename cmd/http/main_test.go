package main

import "os"
import "testing"

import "gopkg.in/check.v1"

func TestCommand(t *testing.T) {
	check.Suite(&CommandSuite{})
	check.TestingT(t)
}

type CommandSuite struct{}

func (s *CommandSuite) TestFlags(c *check.C) {
	os.Args = append(os.Args, "--listen", ":4000")

	f := flags()

	c.Check(f.address, check.Equals, ":4000")
}
