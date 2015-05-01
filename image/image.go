package image

// #cgo CXXFLAGS: -std=c++11
// #cgo pkg-config: opencv
// #include "./image.h"
import "C"
import "errors"
import "unsafe"

import "github.com/vincent-petithory/dataurl"

type Image struct {
	Bytes  []byte
	Format string
}

var ErrInvalidContentType = errors.New("invalid content type")

func (i *Image) Process() *Image {
	cin := &C.image{
		length: C.uint(len(i.Bytes)),
		format: C.CString(i.Format),
		bytes:  (*C.uchar)(unsafe.Pointer(&i.Bytes[0])),
	}

	cout := C.process(cin)
	defer C.free(unsafe.Pointer(cout))

	return &Image{
		Bytes:  C.GoBytes(unsafe.Pointer(cout.bytes), C.int(cout.length)),
		Format: C.GoString(cout.format),
	}
}

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
