package fastHttpSwagger

import (
	"context"
	"golang.org/x/net/webdav"
	"io/fs"
	"os"
)

type myfile struct {
	file webdav.File
}

func (m *myfile) Stat() (fs.FileInfo, error) {
	return m.file.Stat()
}

func (m *myfile) Read(bytes []byte) (int, error) {
	return m.file.Read(bytes)
}

func (m *myfile) Seek(offset int64, whence int) (int64, error) {
	return m.file.Seek(offset, whence)
}

func (m *myfile) Close() error {
	return m.file.Close()
}

type myfs struct {
	fs webdav.FileSystem
}

func (f *myfs) Open(name string) (fs.File, error) {
	file, err := f.fs.OpenFile(context.Background(), name, os.O_WRONLY, fs.ModePerm)
	if err != nil {
		return nil, err
	}
	return &myfile{file: file}, err
}
