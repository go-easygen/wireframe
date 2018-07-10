////////////////////////////////////////////////////////////////////////////
// Program: wireframe
// Purpose: wire framing
// Authors: Myself <me@mine.org> (c) 2018, All rights reserved
////////////////////////////////////////////////////////////////////////////

package main

import (
	"fmt"

	"github.com/go-easygen/cli"
	"github.com/go-easygen/cli/clis"
)

////////////////////////////////////////////////////////////////////////////
// put

func putCLI(ctx *cli.Context) error {
	rootArgv = ctx.RootArgv().(*rootT)
	argv := ctx.Argv().(*putT)
	clis.Setup(progname, rootArgv.Verbose.Value())
	clis.Verbose(2, "[put]:\n  %+v\n  %+v\n  %v\n", rootArgv, argv, ctx.Args())
	Opts.Self, Opts.Host, Opts.Port, Opts.Daemonize, Opts.Verbose =
		rootArgv.Self, rootArgv.Host, rootArgv.Port, rootArgv.Daemonize, rootArgv.Verbose.Value()
	//return nil
	return DoPut()
}

//
func DoPut() error {
	fmt.Printf("%s v %s. Upload into service\n", progname, version)
	fmt.Println("Copyright (C) 2018, Myself <me@mine.org>")
	return nil
}
