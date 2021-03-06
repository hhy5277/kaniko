/*
Copyright 2018 Google LLC

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package dockerfile

import (
	"bytes"
	"strings"

	"github.com/docker/docker/builder/dockerfile/instructions"
	"github.com/docker/docker/builder/dockerfile/parser"
)

// Parse parses the contents of a Dockerfile and returns a list of commands
func Parse(b []byte) ([]instructions.Stage, error) {
	p, err := parser.Parse(bytes.NewReader(b))
	if err != nil {
		return nil, err
	}
	stages, _, err := instructions.Parse(p.AST)
	if err != nil {
		return nil, err
	}
	return stages, err
}

// ParseCommands parses an array of commands into an array of instructions.Command; used for onbuild
func ParseCommands(cmdArray []string) ([]instructions.Command, error) {
	var cmds []instructions.Command
	cmdString := strings.Join(cmdArray, "\n")
	ast, err := parser.Parse(strings.NewReader(cmdString))
	if err != nil {
		return nil, err
	}
	for _, child := range ast.AST.Children {
		cmd, err := instructions.ParseCommand(child)
		if err != nil {
			return nil, err
		}
		cmds = append(cmds, cmd)
	}
	return cmds, nil
}
