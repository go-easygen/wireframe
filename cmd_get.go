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

func getCLI(ctx *cli.Context) error {
	rootArgv = ctx.RootArgv().(*rootT)
	argv := ctx.Argv().(*getT)
	fmt.Printf("[get]:\n  %+v\n  %+v\n  %v\n", rootArgv, argv, ctx.Args())
	return DoGet()
}

func DoGet() error {
	fmt.Printf("%s v %s. Get from the service\n", progname, version)
	fmt.Println("Copyright (C) 2018, Myself <me@mine.org>")
	return nil
}
