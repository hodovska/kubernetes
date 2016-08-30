/*
Copyright 2015 The Kubernetes Authors.

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

package cmd

import (
	"fmt"
	"io"

	"github.com/renstrom/dedent"
	"github.com/spf13/cobra"

	"k8s.io/kubernetes/pkg/kubectl"
	cmdutil "k8s.io/kubernetes/pkg/kubectl/cmd/util"
)

var (
	limitLong = dedent.Dedent(`
		Create a limit with the specified name and properties.`)

	//TODO: add/modify examples/
	limitExample = dedent.Dedent(`
		# Create new limit named podmax2cpu for pods to have maximum of 2 CPU’s
		kubectl create limit new --set=pod.cpu.max=3, pod.cpu.default=2, container.memory.min=200M`)
)

func NewCmdCreateLimit(f *cmdutil.Factory, cmdOut io.Writer) *cobra.Command {
	cmd := &cobra.Command{
		//TODO: add
		Use:     "limit NAME []",
		Short:   "Create a limit with the specified name and properties",
		Long:    limitLong,
		Example: limitExample,
		Run: func(cmd *cobra.Command, args []string) {
			err := CreateLimit(f, cmdOut, cmd, args)
			cmdutil.CheckErr(err)
		},
	}
	cmdutil.AddApplyAnnotationFlags(cmd)
	cmdutil.AddValidateFlags(cmd)
	cmdutil.AddPrinterFlags(cmd)
	cmdutil.AddGeneratorFlags(cmd, cmdutil.LimitV1GeneratorName)
	cmd.Flags().String("set", "", "A comma-delimited set of type.resource.limit=value pairs that define a limit")
	return cmd
}

func CreateLimit(f *cmdutil.Factory, cmdOut io.Writer, cmd *cobra.Command, args []string) error {
	name, err := NameFromCommandArgs(cmd, args)
	if err != nil {
		return err
	}
	var generator kubectl.StructuredGenerator
	switch generatorName := cmdutil.GetFlagString(cmd, "generator"); generatorName {
	case cmdutil.LimitV1GeneratorName:
		//TODO: vlastnosti
		generator = &kubectl.LimitGeneratorV1{
			Name: name,
			Set: cmdutil.GetFlagString(cmd, "set"),
		}
	default:
		return cmdutil.UsageError(cmd, fmt.Sprintf("Generator %s not supported.", generatorName))
	}
	return RunCreateSubcommand(f, cmd, cmdOut, &CreateSubcommandOptions{
		Name:                name,
		StructuredGenerator: generator,
		DryRun:              cmdutil.GetFlagBool(cmd, "dry-run"),
		OutputFormat:        cmdutil.GetFlagString(cmd, "output"),
	})
}
