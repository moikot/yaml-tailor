package main

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

	"github.com/moikot/yaml-tailor/commands"
	"github.com/spf13/cobra"
)

type command interface {
	Execute() error
	Find(args []string) (*cobra.Command, []string, error)
}

var (
	newRootCmdFunc = newRootCmd
	exitFunc       = os.Exit
)

func newRootCmd(args []string) command {
	cmd := commands.NewRootCmd()
	flags := cmd.PersistentFlags()
	_ = flags.Parse(args)
	return cmd
}

func main() {
	cmd := newRootCmdFunc(os.Args[1:])
	if err := cmd.Execute(); err != nil {
		exitFunc(1)
	}
}
