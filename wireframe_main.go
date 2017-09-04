////////////////////////////////////////////////////////////////////////////
// Program: wireframe
// Purpose: wire framing
// Authors: Tong Sun (c) 2017, All rights reserved
////////////////////////////////////////////////////////////////////////////

package main

//go:generate sh -v wireframe_cliGen.sh

import (
	"fmt"
	"os"

	"github.com/mkideal/cli"
)

////////////////////////////////////////////////////////////////////////////
// Global variables definitions

var (
	progname = "wireframe"
	version  = "0.1.0"
	date     = "2017-09-04"
)

var rootArgv *rootT

////////////////////////////////////////////////////////////////////////////
// Function definitions

// Function main
func main() {
	cli.SetUsageStyle(cli.ManualStyle) // up-down, for left-right, use NormalStyle
	//NOTE: You can set any writer implements io.Writer
	// default writer is os.Stdout
	if err := cli.Root(root,
		cli.Tree(putDef),
		cli.Tree(getDef)).Run(os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
	fmt.Println("")
}

//==========================================================================
// Main dispatcher

func wireframe(ctx *cli.Context) error {
	ctx.JSON(ctx.RootArgv())
	ctx.JSON(ctx.Argv())
	fmt.Println()

	return nil
}
