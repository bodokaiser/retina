package image

// #cgo CXXFLAGS: -std=c++11
// #cgo pkg-config: opencv
// #include "./image.h"
import "C"
import "errors"

import "github.com/vincent-petithory/dataurl"

func Process(src, dst string) {
	C.process(C.CString(src), C.CString(dst))
}

type Image struct {
	Bytes  []byte
	Format string
}

var ErrInvalidContentType = errors.New("invalid content type")

func (i *Image) UnmarshalJSON(b []byte) error {
	s := string(b[1 : len(b)-1])

	d, err := dataurl.DecodeString(s)
	if err != nil {
		return err
	}

	if d.MediaType.Type != "image" {
		return ErrInvalidContentType
	}

	i.Bytes = d.Data
	i.Format = d.MediaType.Subtype

	return nil
}
