////////////////////////////////////////////////////////////////////////////
// Program: wireframe
// Purpose: wire framing
// Authors: Myself <me@mine.org> (c) 2018, All rights reserved
////////////////////////////////////////////////////////////////////////////

package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"

	"github.com/mkideal/cli"
	"github.com/mkideal/cli/clis"
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

	fileo, err := ioutil.TempFile(os.TempDir(), progname+".tmp.")
	clis.AbortOn("Create temp file", err)
	return DoPut(argv.Filei, fileo)
}

//
func DoPut(bi io.Reader, bw io.Writer) error {
	fmt.Printf("%s v %s. Upload into service\n", progname, version)
	fmt.Println("Copyright (C) 2018, Myself <me@mine.org>")
	return nil
}
