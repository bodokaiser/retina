package main

import "io/ioutil"

import "github.com/bodokaiser/retina/image"

func main() {
	buf, _ := ioutil.ReadFile("../medex/retina.jpg")

	src := &image.Image{buf, "jpeg"}
	src.Process()
}
