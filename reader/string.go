package reader

import (
	"strings"
)

func NewStringReader(src string) *strings.Reader {
	return strings.NewReader(src)
}
