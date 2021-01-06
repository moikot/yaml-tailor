package commands

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
	"github.com/moikot/djson"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"

	"github.com/moikot/yaml-tailor/yaml"
)

type yamlReaderWriter interface {
	ReadYaml(filename string) (map[string]interface{}, error)
	WriteYaml(filename string, yaml map[string]interface{}) error
}

type rootCmd struct {
	*cobra.Command

	values  []string
	strings []string

	yamlReaderWriter
}

// NewRootCmd creates a new instance of the root command.
func NewRootCmd() *cobra.Command {
	cmd := &rootCmd{
		yamlReaderWriter: yaml.NewReaderWriter(),
	}

	cmd.Command = &cobra.Command{
		Use:   "yaml-tailor [file]",
		Short: "Tailor YAML files to your needs",
		Long:  "YAML talor can adjust content of YAML files by overriding and adding elements.",
		Args:  cobra.ExactArgs(1),
		RunE: func(c *cobra.Command, args []string) error {
			return cmd.run(args)
		},
	}

	f := cmd.Flags()
	f.StringArrayVarP(&cmd.values, "value", "v", []string{}, "a value override")
	f.StringArrayVarP(&cmd.strings, "string", "s", []string{}, "a string override")

	return cmd.Command
}

func (c *rootCmd) run(args []string) error {
	m, err := c.ReadYaml(args[0])
	if err != nil {
		return errors.Wrapf(err, "failed to read YAML file %q", args[0])
	}

	for _, s := range c.values {
		err := djson.MergeValue(m, s)
		if err != nil {
			return errors.Wrapf(err, "failed to merge value %q", s)
		}
	}

	for _, s := range c.strings {
		err := djson.MergeString(m, s)
		if err != nil {
			return errors.Wrapf(err, "failed to merge string %q", s)
		}
	}

	return c.WriteYaml(args[0], m)
}
