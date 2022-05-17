package markdown

import "github.com/livebud/bud/runtime/transform"

type Transform struct{}

func (t *Transform) MdToSvelte(file *transform.File) error {
	file.Code = []byte("<h1>hello</h1>")
	return nil
}
