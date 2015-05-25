package model

import "github.com/bodokaiser/retina/image"

type Report struct {
	Input  *image.Image
	Output *image.Image
}

func (r *Report) Process() error {
	r.Output = r.Input.Process()

	return nil
}
