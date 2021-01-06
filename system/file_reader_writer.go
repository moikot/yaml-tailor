package system

// -------------------------------------------------------------------------------
// MIT License
// Copyright (c) 2021 Sergey Anisimov
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.
// -------------------------------------------------------------------------------

import (
	"io/ioutil"
	"os"
)

// FileReaderWriter reades files from disks and writes them to disks.
type FileReaderWriter struct {
	readFile  func(file string) ([]byte, error)
	writeFile func(file string, data []byte, perm os.FileMode) error
}

// NewFileReaderWriter creates a new instance of a file reader/writer.
func NewFileReaderWriter() *FileReaderWriter {
	return &FileReaderWriter{
		readFile:  ioutil.ReadFile,
		writeFile: ioutil.WriteFile,
	}
}

// ReadFile reads the file named by filename and returns the contents.
func (u *FileReaderWriter) ReadFile(file string) ([]byte, error) {
	return u.readFile(file)
}

// WriteFile writes data to a file named by filename.
func (u *FileReaderWriter) WriteFile(file string, data []byte, perm os.FileMode) error {
	return u.writeFile(file, data, perm)
}
