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
	"strings"

	"github.com/mkideal/cli"
	"github.com/labstack/gommon/color"
)

////////////////////////////////////////////////////////////////////////////
// Constant and data type/structure definitions

// The OptsT type defines all the configurable options from cli.
type OptsT struct {
	Self      *rootT
	Host      string
	Port      int
	Daemonize bool
	Verbose   int
}

////////////////////////////////////////////////////////////////////////////
// Global variables definitions

var (
	progname = "wireframe"
	version  = "0.1.0"
	date     = "2017-09-04"

	rootArgv *rootT
	// Opts store all the configurable options
	Opts OptsT
)

////////////////////////////////////////////////////////////////////////////
// Function definitions

// Function main
func main() {
	// cli.SetUsageStyle(cli.ManualStyle) // up-down, for left-right, use NormalStyle
	cli.SetUsageStyle(cli.DenseNormalStyle) // for left-right
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

// SUPPORT-FUNCTIONS
//==========================================================================
// support functions

// Abs returns the absolute value of x.
func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

// Basename returns the file name without extension.
func Basename(s string) string {
	n := strings.LastIndexByte(s, '.')
	if n > 0 {
		return s[:n]
	}
	return s
}

// IsExist checks if the given file exist
func IsExist(fileName string) bool {
	_, err := os.Stat(fileName)
	return err == nil || os.IsExist(err)
}

func warning(m string) {
	fmt.Fprintf(os.Stderr, "[%s] %s: %s\n", progname, color.Yellow("Warning"), m)
}

func warnOn(errCase string, e error) {
	if e != nil {
		fmt.Fprintf(os.Stderr, "[%s] %s, %s: %v\n",
			color.White(progname), color.Yellow("Error"), errCase, e)
	}
}

// abortOn will quit on anticipated errors gracefully without stack trace
func abortOn(errCase string, e error) {
	if e != nil {
		fmt.Fprintf(os.Stderr, "[%s] %s, %s: %v\n",
			color.White(progname), color.Red("Error"), errCase, e)
		os.Exit(1)
	}
}

// verbose will print info to stderr according to the verbose level setting
func verbose(levelSet int, format string, args ...interface{}) {
	if Opts.Verbose >= levelSet {
		fmt.Fprintf(os.Stderr, "[%s] ", color.White(progname))
		fmt.Fprintf(os.Stderr, format+"\n", args...)
	}
}
