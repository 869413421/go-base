package main

type IFile interface {
	Read(buf []byte) (n int, err error)
	Write(buf []byte) (n int, err error)
	Seek(off int64, whence int) (pos int64, err error)
	Close() error
}

type File struct {
}

func (f *File) Read(buf []byte) (n int, err error) {
	return
}
func (f *File) Write(buf []byte) (n int, err error) {
	return
}
func (f *File) Seek(off int64, whence int) (pos int64, err error) {
	return
}
func (f *File) Close() error {
	return nil
}
