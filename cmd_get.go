////////////////////////////////////////////////////////////////////////////
// Program: wireframe
// Purpose: wire framing
// Authors: Myself <me@mine.org> (c) 2018, All rights reserved
////////////////////////////////////////////////////////////////////////////

package main

import (
	"fmt"
	"io"
	"os"
	"strings"

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

	if !ctx.IsSet("--output") {
		fileo, err := os.Create(
			strings.Replace(argv.Filei.Name(), ".org", ".new", 1))
		clis.AbortOn("Output file create", err)
		argv.Fileo.SetWriter(fileo)
	}
	fileo := argv.Fileo
	clis.Verbose(2, "] %s\n", fileo.Name())
	defer fileo.Close()

	return DoGet(fileo)
}

//
func DoGet(w io.Writer) error {
	fmt.Printf("%s v %s. Get from the service\n", progname, version)
	fmt.Println("Copyright (C) 2018, Myself <me@mine.org>")
	fmt.Fprintf(w, "Someting\n")
	return nil
}
