// Copyright (C) 2015 Scaleway. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE.md file.

package commands

import (
	"fmt"
	"os"

	log "github.com/scaleway/scaleway-cli/vendor/github.com/Sirupsen/logrus"

	types "github.com/scaleway/scaleway-cli/commands/types"
)

var cmdWorkflow = &types.Command{
	Exec:        runWorkflow,
	UsageLine:   "_workflow [OPTIONS] [WORKFLOW] [ARGS...]",
	Description: "This command is experimental and may change between versions",
	Hidden:      true,
	Help:        "Execute a workflow",
	Examples: `
    $ scw _workflow create-image-from-s3 http://test-images.fr-1.storage.online.net/scw-app-docker-latest.tar
`,
}

func init() {
	cmdWorkflow.Flag.BoolVar(&workflowHelp, []string{"h", "-help"}, false, "Print usage")
}

// Flags
var workflowHelp bool // -h, --help flag

type WorkflowFunc func(cmd *types.Command, args []string)

func WorkflowCreateImageFromS3(cmd *types.Command, args []string) {
	if len(args) != 1 {
		cmd.PrintUsage()
	}

	imageName := "new-image"
	log.Debugf("URL of the tarball: %s", args[0])
	log.Debugf("Target name: %s", imageName)
	/*server, err := sandbox.call("create", "--bootscript=rescue", "--volume=50G", "--name=image-writer", "1GB")
	if err != nil {
		log.Fatal(err)
	}*/

}

func runWorkflow(cmd *types.Command, args []string) {
	if workflowHelp {
		cmd.PrintUsage()
	}

	workflows := map[string]WorkflowFunc{
		"create-image-from-s3": WorkflowCreateImageFromS3,
	}

	if len(args) == 0 {
		fmt.Fprintf(os.Stderr, "Available workflows:\n")
		for name := range workflows {
			fmt.Fprintf(os.Stderr, "- %s\n", name)
		}
		cmd.PrintShortUsage()
	}

	if f, ok := workflows[args[0]]; ok {
		f(cmd, args[1:])
	} else {
		cmd.PrintUsage()
	}
}
