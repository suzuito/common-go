package cmarkdown

import (
	"context"
	"io"
)

type Converter interface {
	Convert(
		ctx context.Context,
		source []byte,
		w io.Writer,
		tocs *[]CMTOC,
		images *[]CMImage,
	) error
}
