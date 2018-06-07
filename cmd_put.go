////////////////////////////////////////////////////////////////////////////
// Program: wireframe
// Purpose: wire framing
// Authors: Tong Sun (c) 2017, All rights reserved
////////////////////////////////////////////////////////////////////////////

package main

import (
	"fmt"

	"github.com/go-easygen/cli"
)

////////////////////////////////////////////////////////////////////////////
// put

func putCLI(ctx *cli.Context) error {
	rootArgv = ctx.RootArgv().(*rootT)
	argv := ctx.Argv().(*putT)
	fmt.Printf("[put]:\n  %+v\n  %+v\n  %v\n", rootArgv, argv, ctx.Args())
	return DoPut()
}

func DoPut() error {
	fmt.Printf("%s v %s. Upload into service\n", progname, version)
	fmt.Println("Copyright (C) 2018, Myself <me@mine.org>")
	return nil
}
