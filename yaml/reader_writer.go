package yaml

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
	"os"

	"github.com/moikot/yaml-tailor/system"
	"github.com/pkg/errors"
	yaml "gopkg.in/yaml.v2"
)

type fileReaderWriter interface {
	ReadFile(file string) ([]byte, error)
	WriteFile(file string, data []byte, perm os.FileMode) error
}

// ReaderWriter reads and writes yaml files.
type ReaderWriter struct {
	fileReaderWriter
}

// NewReaderWriter creates a new instance of the YAML reader/writer.
func NewReaderWriter() *ReaderWriter {
	return &ReaderWriter{
		fileReaderWriter: system.NewFileReaderWriter(),
	}
}

// ReadYaml reads a YAML file and deserializes it.
func (c *ReaderWriter) ReadYaml(filename string) (map[string]interface{}, error) {
	bytes, err := c.ReadFile(filename)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to read from file %q", filename)
	}

	var yamlObj map[string]interface{}
	if err := yaml.Unmarshal(bytes, &yamlObj); err != nil {
		return nil, errors.Wrap(err, "failed to unmarshal YAML")
	}
	return yamlObj, nil
}

// WriteYaml serializes a YAML object and writes it to a file.
func (c *ReaderWriter) WriteYaml(filename string, yamlObj map[string]interface{}) error {
	bytes, err := yaml.Marshal(yamlObj)
	if err != nil {
		return errors.Wrap(err, "failed to marshal YAML")
	}

	err = c.WriteFile(filename, bytes, 0644)
	if err != nil {
		return errors.Wrapf(err, "failed to write YAML to file %q", filename)
	}
	return nil
}
