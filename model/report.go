package model

import "github.com/bodokaiser/retina/image"

type Report struct {
	Input  *image.Image `json:"input"`
	Output *image.Image `json:"output"`
}

func (r *Report) Process() error {
	r.Output = r.Input.Process()

	return nil
}
