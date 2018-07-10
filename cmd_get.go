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
// get

func getCLI(ctx *cli.Context) error {
	rootArgv = ctx.RootArgv().(*rootT)
	argv := ctx.Argv().(*getT)
	clis.Setup(progname, rootArgv.Verbose.Value())
	clis.Verbose(2, "[get]:\n  %+v\n  %+v\n  %v\n", rootArgv, argv, ctx.Args())
	Opts.Self, Opts.Host, Opts.Port, Opts.Daemonize, Opts.Verbose =
		rootArgv.Self, rootArgv.Host, rootArgv.Port, rootArgv.Daemonize, rootArgv.Verbose.Value()
	//return nil
	return DoGet()
}

//
func DoGet() error {
	fmt.Printf("%s v %s. Get from the service\n", progname, version)
	fmt.Println("Copyright (C) 2018, Myself <me@mine.org>")
	return nil
}
