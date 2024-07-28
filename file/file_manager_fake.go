package file

// FileManagerFake provides a fake implementation of the FileManagerInterface.
type FileManagerFake struct {
	content string
	err     error
}

// Read returns the predefined content speficied in `f.content`. If `f.err` is
// defined, then it returns the error instead.
func (f *FileManagerFake) Read(path string) (string, error) {
	if f.err != nil {
		return "", f.err
	}
	return f.content, nil
}
