package structs

type TemplateBuffer struct {
	Buffer *string
}

func (t *TemplateBuffer) Write(p []byte) (n int, err error) {
	*t.Buffer += string(p)
	return len(p), nil
}
